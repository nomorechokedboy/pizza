package handler

import (
	"api-blog/api/config"
	"api-blog/api/presenter"
	"api-blog/api/util"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"bytes"
	"net/http"
	"net/smtp"
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
// @Tags Users
// @Param todo body entities.UserLogin true "New User"
// @Accept json
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 409
// @Router /api/v1/users/register [post]
func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := new(entities.UserLogin)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if _, err := handler.usecase.GetUserByIdentifier(req.Identifier); err == nil {
		return fiber.NewError(fiber.StatusConflict, "indentifier already existed")
	}
	err := handler.usecase.CreateUser(req.Password, req.Identifier)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create New user")
	}
	return c.JSON(presenter.Response{
		Status:  http.StatusCreated,
		Message: "Create success",
		Data:    nil,
	})
}

// Login
// @Login godoc
// @Summary User Login
// @Description Use for login response the refresh_token and accessToken
// @Tags Users
// @Accept json
// @Param todo body entities.UserLogin true "Login"
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 403
// @Router /api/v1/users/login [post]
func (handler *UserHandler) Login(c *fiber.Ctx) error {

	req := new(entities.UserLogin)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	user, err := handler.usecase.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Identifier does not exist")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "incorrect password")
	}

	accessToken, refreshToken := util.GenerateToken(user.Id, []byte(handler.config.AuthConfig.JWTSecret), []byte(handler.config.AuthConfig.JWTRefreshToken))

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(presenter.Response{
		Status:  http.StatusOK,
		Message: "Login Success",
		Data: &fiber.Map{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"user":          user,
		},
	})
}

// GetAuthUserByToken
// @GetUserByAuthId godoc
// @Summary Get user infor by token
// @Description Get UserInfo by accessToken
// @Tags Users
// @Accept json
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/users/ [get]
func (handler *UserHandler) GetAuthUserById(c *fiber.Ctx) error {
	authId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not parse Id from token")
	}
	user, err := handler.usecase.GetUserById(authId)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(presenter.Response{
		Status:  200,
		Message: "Success",
		Data: &fiber.Map{
			"user": user,
		},
	})
}

// UpdateUserByToken
// @Update godoc
// @Summary Update user infor by token
// @Description Update UserInfo by Id from accessToken
// @Tags Users
// @Accept json
// @Param todo body entities.UserReq true "Updated User"
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/users/update [put]
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
	if _, err := handler.usecase.GetUserByUsername(req.Username); err == nil {
		return fiber.NewError(fiber.StatusConflict, "Username already existed")
	}
	if err := handler.usecase.UpdateUserInfo(req.Password, req.Fullname, req.Username, req.PhoneNumber, req.Email, req.Avatar, user.Id); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not update")
	}
	newUser, err := handler.usecase.GetUserById(user.Id)

	return c.JSON(presenter.Response{
		Status:  http.StatusOK,
		Message: "update success",
		Data: &fiber.Map{
			"user": newUser,
		},
	})
}

// ForgotPassword
// @ForgotPassword godoc
// @Summary option when user forgot password
// @Description send email to user for reset password
// @Tags Users
// @Accept json
// @Param todo body entities.UserEmail true "user email"
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/v1/users/forgot-password [post]
func (handler *UserHandler) ForgotPassword(c *fiber.Ctx) error {
	reqEmail := new(entities.UserEmail)
	if err := c.BodyParser(reqEmail); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	user, err := handler.usecase.GetUserByEmail(reqEmail.Email)
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, err.Error())
	}

	accessToken, _ := util.GenerateToken(user.Id, []byte(handler.config.AuthConfig.JWTSecret), []byte(handler.config.AuthConfig.JWTRefreshToken))

	auth := smtp.PlainAuth(
		"",
		handler.config.AuthEmail.Email,
		handler.config.AuthEmail.Password,
		"smtp.gmail.com",
	)
	var emailReponse entities.ResponseEmail
	emailReponse.Link = handler.config.AppAPI.Link + "/userId:" + accessToken
	emailReponse.Username = user.Username
	emailReponse.Sender = "blog team"
	tmpl := template.Must(template.ParseFiles("template/emailTemp.html"))
	buff := new(bytes.Buffer)
	tmpl.Execute(buff, emailReponse)

	subject := "Subject: Reset Password \n"
	mine := "MINE-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mine + buff.String())

	smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		handler.config.AuthEmail.Email,
		[]string{reqEmail.Email},
		msg,
	)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return c.JSON(presenter.Response{
		Status:  200,
		Message: "Please go to your email to reset",
	})
}

// resetpassword
// @Resetpassword godoc
// @Summary reset user password
// @Description reset password by token
// @Tags Users
// @Accept json
// @Param todo body entities.UserPassword true "new Password"
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/users/reset-password [put]
func (handler *UserHandler) ResetPassword(c *fiber.Ctx) error {
	userId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not parse Id from token")
	}
	newPassword := new(entities.UserPassword)
	if err := c.BodyParser(newPassword); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	if err := handler.usecase.UpdatePasswordById(newPassword.Password, userId); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}
	return c.JSON(presenter.Response{
		Status:  http.StatusCreated,
		Message: "Password is reseted",
	})
}

