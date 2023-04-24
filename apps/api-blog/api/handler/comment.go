package handler

import (
	"api-blog/pkg/common"
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

// @GetAllComments godoc
// @Summary Show all comments from comment
// @Description get all comments from specfied comment
// @Tags Comments
// @Param  userID query int false "User ID"
// @Param  postID query int false "Post ID"
// @Param  parentID query int false "Parent ID"
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param sortBy query string false "Sort by" Enums(id, user_id, parent_id) default(id)
// @Success 200 {object} common.BasePaginationResponse[entities.Comment]
// @Failure 404
// @Failure 500
// @Router /comments/ [get]
func (handler *CommentHandler) GetAllComments(c *fiber.Ctx) error {
	query := new(entities.CommentQuery)

	if err := c.QueryParser(query); err != nil {
		return err
	}

	comments, err := handler.usecase.GetAllComments(&entities.CommentQuery{
		UserID:   query.UserID,
		PostID:   query.PostID,
		ParentID: query.ParentID,
		BaseQuery: common.BaseQuery{
			Page:     query.Page,
			PageSize: query.PageSize,
			Sort:     query.Sort,
			SortBy:   query.SortBy,
		},
	})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get all comments")
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}

// @CreateComment godoc
// @Summary Create comment
// @Description Create comment
// @Tags Comments
// @Accept json
// @Param comment body entities.CommentRequest true "Comment"
// @Success 201 {object} entities.Comment
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /comments/ [post]
func (handler *CommentHandler) CreateComment(c *fiber.Ctx) error {
	authID := c.Locals("uId").(uint)
	req := new(entities.CommentRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	comment, err := handler.usecase.CreateComment(authID, req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new comment")
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

// @UpdateComment godoc
// @Description Update comment
// @Summary Update comment with new message
// @Param id path int true "Comment ID"
// @Param comment body handler.UpdateComment.commentRequest true "Comment"
// @Tags Comments
// @Success 200 {object} entities.Comment
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /comments/{id} [put]
func (handler *CommentHandler) UpdateComment(c *fiber.Ctx) error {
	type commentRequest struct {
		Content string
	}

	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment ID")
	}

	req := new(commentRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	comment, err := handler.usecase.UpdateComment(uint(id), req.Content)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update specfied comment")
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}

// @DeleteComment godoc
// @Description Delete comment
// @Summary Delete specified comment
// @Param id path int true "Comment ID"
// @Tags Comments
// @Produce json
// @Success 200 {object} entities.Comment
// @Failure 400
// @Failure 500
// @security ApiKeyAuth
// @Router /comments/{id} [delete]
func (handler *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment ID")
	}

	comment, err := handler.usecase.DeleteComment(uint(id))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete specfied comment")
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}
