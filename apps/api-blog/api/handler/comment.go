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

// @GetAllCommentsBycommentID godoc
// @Summary Show all comments from comment
// @Description get all comments from specfied comment
// @Tags Comments
// @Param  userID query int false "User ID"
// @Param  postID query int false "Post ID"
// @Param  parentID query int false "Parent ID"
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param sortBy query string false "Sort by" Enums(id, title, slug, user_id, parent_id) default(id)
// @Success 200 {array} common.BasePaginationResponse[entities.Comment]
// @Failure 404
// @Failure 500
// @Router /comments/ [get]
func (handler *CommentHandler) GetAllComment(c *fiber.Ctx) error {
	query := new(entities.CommentQuery)

	if err := c.QueryParser(query); err != nil {
		return err
	}

	comments, err := handler.usecase.GetAllComments(&entities.CommentQuery{
		BaseQuery: common.BaseQuery{
			Page:     query.Page,
			PageSize: query.PageSize,
			Sort:     query.Sort,
			SortBy:   query.SortBy,
		},
	})

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to get all comments")
	}

	commentRes := common.BasePaginationResponse[entities.CommentRes]{
		Items: []entities.CommentRes{},
	}
	commentRes.Page = comments.Page
	commentRes.PageSize = comments.PageSize
	commentRes.Total = comments.Total

	for _, comment := range comments.Items {
		commentRes.Items = append(commentRes.Items, comment.ToResponse())
	}

	return c.Status(fiber.StatusOK).JSON(commentRes)
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
