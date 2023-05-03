package handler

import (
	"api-blog/api/config"
	"api-blog/api/util"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"api-blog/templates"
	"bytes"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase usecase.UserUsecase
	config  config.Config
}

func NewUserHandler(usecase usecase.UserUsecase, config config.Config) *UserHandler {

	return &UserHandler{
		usecase: usecase,
		config:  config,
	}
}

// @CreateUser godoc
// @Summary Create User
// @Description Create New UserUsecase
// @Tags Auth
// @Param todo body entities.SignUpBody true "New User"
// @Accept json
// @Success 200
// @Router /auth/register [post]
func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := new(entities.SignUpBody)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "something bad happened")
	}
	req.Password = string(hashPassword)
	errors := handler.usecase.CreateUser(*req)
	if errors != nil && strings.Contains(errors.Error(), "duplicate key value violates unique") {
		if strings.Contains(errors.Error(), "username") {
			return fiber.NewError(fiber.StatusConflict, "Username already exist")
		}
		return fiber.NewError(fiber.StatusConflict, "Email already exist")
	} else if errors != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something bad happened")
	}
	return c.Status(http.StatusCreated).SendString("Create success")
}

// Login
// @Login godoc
// @Summary User Login
// @Description Use for login response the refresh_token and accessToken
// @Tags Auth
// @Accept json
// @Param todo body entities.UserLogin true "Login"
// @Success 200 {object} entities.Auth{}
// @Router /auth/login [post]
func (handler *UserHandler) Login(c *fiber.Ctx) error {

	req := new(entities.UserLogin)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	user, err := handler.usecase.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Identifier does not exist")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "incorrect password")
	}

	accessToken, refreshToken := util.GenerateToken(user.Id, []byte(handler.config.AuthConfig.JWTSecret), []byte(handler.config.AuthConfig.JWTRefreshToken))

	if err != nil {
		return fiber.ErrInternalServerError
	}
	tokenResponse := entities.Auth{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(tokenResponse)
}

// GetAuthUserByToken
// @GetUserByAuthId godoc
// @Summary Get user infor by token
// @Description Get UserInfo by accessToken
// @Tags Auth
// @Accept json
// @Success 200 {object} entities.UserResponse{}
// @Security ApiKeyAuth
// @Router /auth/me [get]
func (handler *UserHandler) GetAuthUserById(c *fiber.Ctx) error {
	authId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not parse Id from token")
	}
	user, err := handler.usecase.GetUserById(authId)
	if err != nil {
		return fiber.ErrNotFound
	}

	newUser := entities.UserResponse{
		Id:          user.Id,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Username:    user.Username,
		Fullname:    user.Fullname,
		PhoneNumber: user.PhoneNumber,
	}
	return c.Status(fiber.StatusOK).JSON(newUser)
}

// UpdateUserByToken
// @Update godoc
// @Summary Update user infor by token
// @Description Update UserInfo by Id from accessToken
// @Tags User
// @Accept json
// @Param todo body entities.UserReq true "Updated User"
// @Success 200 {object} entities.User
// @Security ApiKeyAuth
// @Router /user/update [put]
func (handler *UserHandler) UpdateUserById(c *fiber.Ctx) error {
	authId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not parse Id from token")
	}
	user, err := handler.usecase.GetUserById(authId)
	if err != nil {
		return fiber.ErrNotFound
	}
	req := new(entities.UserReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	updateErr := handler.usecase.UpdateUserInfo(req.Fullname, req.Username, req.PhoneNumber, req.Email, req.Avatar, user.Id)
	if updateErr != nil && strings.Contains(updateErr.Error(), "duplicate key value violates unique") {
		if strings.Contains(updateErr.Error(), "username") {
			return fiber.NewError(fiber.StatusConflict, "Username already exist")
		}
		return fiber.NewError(fiber.StatusConflict, "Email already exist")
	} else if updateErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something bad happened")
	}
	newUser, err := handler.usecase.GetUserById(user.Id)

	return c.JSON(newUser)
}

type UserEmailReq struct {
	Email string
}

// ForgotPassword
// @ForgotPassword godoc
// @Summary option when user forgot password
// @Description send email to user for reset password
// @Tags Auth
// @Accept json
// @Param todo body handler.UserEmailReq true "user email"
// @Success 200
// @Router /auth/forgot-password [post]
func (handler *UserHandler) ForgotPassword(c *fiber.Ctx) error {
	reqEmail := new(UserEmailReq)
	if err := c.BodyParser(reqEmail); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	user, err := handler.usecase.GetUserByEmail(reqEmail.Email)
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, err.Error())
	}
	go handler.sendMail(reqEmail, user)
	return c.SendString("Please check your email")
}

func (handler *UserHandler) sendMail(reqEmail *UserEmailReq, user *entities.User) {
	accessToken, _ := util.GenerateToken(user.Id, []byte(handler.config.AuthConfig.JWTSecret), []byte(handler.config.AuthConfig.JWTRefreshToken))

	auth := smtp.PlainAuth(
		"",
		handler.config.AuthEmail.Email,
		handler.config.AuthEmail.Password,
		"smtp.gmail.com",
	)
	var emailReponse entities.ResponseEmail
	emailReponse.Link = handler.config.AppAPI.Link + "/forgot-password?token=" + accessToken
	emailReponse.Username = user.Username
	emailReponse.Sender = "Blog team"
	tmpl := template.Must(template.New("").Parse(templates.TemplateEmail))
	buff := new(bytes.Buffer)
	tmpl.Execute(buff, emailReponse)

	subject := "Subject: Reset Password \n"
	mine := "MINE-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mine + buff.String())

	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		handler.config.AuthEmail.Email,
		[]string{reqEmail.Email},
		msg,
	); err != nil {
		log.Println("Failed to send email! Err: ", err)
	}
}

// resetpassword
// @Resetpassword godoc
// @Summary reset user password
// @Description reset password by token
// @Tags Auth
// @Accept json
// @Param todo body handler.ResetPassword.resetPasswordReq true "new Password"
// @Success 200
// @Router /auth/reset-password [put]
func (handler *UserHandler) ResetPassword(c *fiber.Ctx) error {
	type resetPasswordReq struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}
	newPassword := new(resetPasswordReq)
	if err := c.BodyParser(newPassword); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	userId, err := util.ParseToken(newPassword.Token, []byte(handler.config.AuthConfig.JWTSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword.Password), 14)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if err := handler.usecase.UpdatePasswordById(string(hashPassword), userId); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}
	return c.Status(http.StatusOK).SendString("Password is reseted")
}

// FindUserById
// @FindUserById godoc
// @Summary find user profile
// @Tags User
// @Accept json
// @Param id path string true "User Id"
// @Success 200 {object} entities.UserResponse{}
// @Router /user/{id} [get]
func (handler *UserHandler) FindUserById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	user, err := handler.usecase.GetUserById(uint(userId))
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}
	userRes := entities.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		Fullname:    user.Fullname,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Avatar:      user.Avatar,
	}
	return c.Status(http.StatusOK).JSON(userRes)
}

// UpdatePassword
// @Update godoc
// @Description Update Password
// @Tags Auth
// @Accept json
// @Param todo body entities.UpdatePassword true "Updated Password"
// @Success 200
// @Security ApiKeyAuth
// @Router /auth/update-password [put]
func (handler *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	authId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Token is wrong or missing")
	}
	user, err := handler.usecase.GetUserById(authId)
	if err != nil {
		return fiber.ErrNotFound
	}
	var req entities.UpdatePassword
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "incorrect password")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 14)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "something bad happended")
	}
	if err := handler.usecase.UpdatePasswordById(string(hashPassword), authId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("update success")

}
