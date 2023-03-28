package handler

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	usecase usecase.CommentUsecase
}

func NewCommentHandler(usecase usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{usecase: usecase}
}

// @GetAllCommentsBycommentID godoc
// @Summary Show all comments from comment
// @Description get all comments from specfied comment
// @Tags Comments
// @Param  userID query int false "User ID"
// @Param  postID query int false "Post ID"
// @Param  parentID query int false "Parent ID"
// @Success 200 {array} entities.CommentRes{}
// @Failure 404
// @Failure 500
// @Router /comments/ [get]
func (handler *CommentHandler) GetAllComment(c *fiber.Ctx) error {
	userID := c.QueryInt("userID")
	postID := c.QueryInt("postID")
	parentID := c.QueryInt("parentID")

	toResponse := func(comments []entities.Comment) []entities.CommentRes {
		commentReponse := []entities.CommentRes{}

		for _, comment := range comments {
			commentRes := entities.CommentRes{
				ID:        comment.ID,
				UserID:    comment.UserID,
				PostID:    comment.PostID,
				ParentID:  comment.ParentID,
				Content:   comment.Content,
				CreatedAt: comment.CreatedAt,
				UpdatedAt: comment.UpdatedAt,
			}

			commentReponse = append(commentReponse, commentRes)
		}

		return commentReponse
	}

	if userID == 0 && parentID == 0 && postID == 0 {
		comments, err := handler.usecase.GetAllComments()

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "failed to get all comments")
		}

		return c.Status(fiber.StatusOK).JSON(toResponse(comments))
	}

	comments, err := handler.usecase.GetAllCommentsByQuery(uint(userID), uint(postID), uint(parentID))

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to get all comments")
	}

	return c.Status(fiber.StatusOK).JSON(toResponse(comments))
}

// @CreateComment godoc
// @Summary Create comment
// @Description Create comment
// @Tags Comments
// @Accept json
// @Param comment body entities.CommentReq true "Comment"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /comments/ [post]
func (handler *CommentHandler) CreateComment(c *fiber.Ctx) error {
	authID := c.Locals("uId").(uint)
	req := new(entities.CommentReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := handler.usecase.CreateComment(authID, req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new comment")
	}

	return c.Status(fiber.StatusCreated).SendString("Created successfully")
}

// @UpdateComment godoc
// @Description Update comment
// @Summary Update comment with new message
// @Param id path int true "Comment ID"
// @Param comment body handler.UpdateComment.commentReq true "Comment"
// @Tags Comments
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /comments/{id} [put]
func (handler *CommentHandler) UpdateComment(c *fiber.Ctx) error {
	type commentReq struct {
		Content string
	}

	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment ID")
	}

	req := new(commentReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := handler.usecase.UpdateComment(uint(id), req.Content); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update specfied comment")
	}

	return c.Status(fiber.StatusOK).SendString("Updated successfully")
}

// @DeleteComment godoc
// @Description Delete comment
// @Summary Delete specified comment
// @Param id path int true "Comment ID"
// @Tags Comments
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @security ApiKeyAuth
// @Router /comments/{id} [delete]
func (handler *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment ID")
	}

	if err := handler.usecase.DeleteComment(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete specfied comment")
	}

	return c.Status(fiber.StatusOK).SendString("Deleted successfully")
}
