package reaction

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"api-blog/src/notification"
	notificationEntities "api-blog/src/notification/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// @ReactToEntity godoc
// @Summary React to a blog post
// @Description Create a reaction to an article
// @Tags Reaction
// @Accept json
// @Param post body entities.WriteReactionBody true "Create reaction body"
// @Success 201 {object} entities.Reaction
// @Failure 400 {string} string
// @Failure 409 {string} string
// @Failure 422 {string} string
// @Failure 500 {string} string
// @Security ApiKeyAuth
// @Router /reaction/react [post]
func ReactToEntity(c *fiber.Ctx) error {
	req := new(entities.WriteReactionBody)
	UserID := c.Locals("uId").(uint)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err := common.Validator.Struct(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request body")
	}

	db := c.Locals("db").(*gorm.DB)
	notifyRepo := c.Locals("notifyRepository").(notification.NotifyRepository)
	userSvc := c.Locals("userService").(usecase.UserUsecase)
	user, err := userSvc.GetUserById(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "User does not exist")
	}

	var notiferID *uint
	var entityData *string
	if req.ReactableType == "posts" {
		post := entities.Post{ID: req.ReactableID}
		if err := db.First(&post).Error; err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Reacted entity does not exist")
		}
		notiferID = &post.UserID
		entityData = &post.Content
	}

	reaction := entities.Reaction{UserID: UserID, User: *user, ReactableID: req.ReactableID, ReactableType: req.ReactableType}
	if res := db.Joins(clause.Associations).Create(&reaction); res.Error != nil {
		if res.Error.(*pgconn.PgError).Code == "23505" {
			return fiber.NewError(fiber.StatusConflict, "Entity already exist")
		}
	}

	if req.ReactableType == "posts" && notiferID != nil && *notiferID != UserID {
		notificationRequest := notificationEntities.NotificationRequest{
			ActorID:    reaction.UserID,
			ActionType: "reacted to your post",
			EntityData: *entityData,
			EntityID:   reaction.ReactableID,
			EntityType: "post",
			NotifierID: *notiferID,
		}
		go notifyRepo.Notify(notificationRequest)
	}

	return c.Status(201).JSON(reaction)
}

// @DropEntityReaction godoc
// @Description Delete a reaction
// @Summary Delete user reaction to a post
// @Tags Reaction
// @Produce json
// @Param post body entities.WriteReactionBody true "Delete reaction body"
// @Success 200 {object} entities.Reaction
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 409 {string} string
// @Failure 422 {string} string
// @Failure 500 {string} string
// @security ApiKeyAuth
// @Router /reaction/drop [delete]
func DropEntityReaction(c *fiber.Ctx) error {
	req := new(entities.WriteReactionBody)
	UserID := c.Locals("uId").(uint)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := common.Validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request body")
	}

	db := c.Locals("db").(*gorm.DB)
	notifyRepo := c.Locals("notifyRepository").(notification.NotifyRepository)
	userSvc := c.Locals("userService").(usecase.UserUsecase)
	user, err := userSvc.GetUserById(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "User does not exist")
	}

	reaction := entities.Reaction{UserID: UserID,
		User:          *user,
		ReactableID:   req.ReactableID,
		ReactableType: req.ReactableType,
	}
	res := db.
		Delete(&reaction)
	if res.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}
	if res.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	var notiferID *uint
	if req.ReactableType == "posts" {
		post := entities.Post{ID: req.ReactableID}
		if err := db.First(&post).Error; err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Reacted entity does not exist")
		}
		notiferID = &post.UserID
	}

	if req.ReactableType == "posts" && notiferID != nil && *notiferID != UserID {
		notificationRequest := notificationEntities.NotificationRequest{
			ActorID:    reaction.UserID,
			ActionType: "reacted to your post",
			EntityID:   reaction.ReactableID,
			EntityType: "post",
			NotifierID: *notiferID,
		}
		go notifyRepo.DeleteNotification(notificationRequest)
	}

	return c.JSON(reaction)
}
