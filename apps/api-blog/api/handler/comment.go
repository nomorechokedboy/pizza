package handler

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"api-blog/src/notification"
	notificationEntities "api-blog/src/notification/entities"
	"context"
	"fmt"
	"log"
	"time"

	_ "api-blog/pkg/common"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type CommentHandler struct {
	usecase usecase.CommentUsecase
	rdb     *redis.Client
}

func NewCommentHandler(usecase usecase.CommentUsecase, rdb *redis.Client) *CommentHandler {
	return &CommentHandler{usecase: usecase, rdb: rdb}
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
// @Tags Comments
// @Router /comments/ [get]
func (handler *CommentHandler) GetAllComments(c *fiber.Ctx) error {
	query := new(entities.CommentQuery)

	if err := c.QueryParser(query); err != nil {
		return err
	}

	// key := fmt.Sprintf(
	// 	"comments-%d-page:%d-pageSize:%d-sort:%s-sortBy:%s-userID:%d",
	// 	query.PostID,
	// 	query.Page,
	// 	query.PageSize,
	// 	query.Sort,
	// 	query.SortBy,
	// 	query.UserID,
	// )
	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	// defer cancel()
	//
	// commentsString, err := handler.rdb.Get(ctx, key).Result()
	// if err == redis.Nil {
	// 	comments, err := handler.usecase.GetAllComments(query)
	// 	if err != nil {
	// 		return fiber.NewError(fiber.StatusInternalServerError, "failed to get all comments")
	// 	}
	//
	// 	commentsJSON, err := json.Marshal(comments)
	// 	err = handler.rdb.Set(ctx, key, commentsJSON, time.Hour*5).Err()
	// 	if err != nil {
	// 		log.Println("There is some problem with redis")
	// 	}
	//
	// 	return c.JSON(comments)
	// } else if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	// }
	//
	// cacheComments := new(common.BasePaginationResponse[entities.Comment])
	// if err := json.Unmarshal([]byte(commentsString), cacheComments); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	// }

	// return c.JSON(cacheComments)

	comments, err := handler.usecase.GetAllComments(query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get all comments")
	}

	return c.JSON(comments)
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

	go invalidateCommentCache(handler.rdb, int(req.PostID))

	isOwner := authID == comment.Post.UserID
	notifyRepo := c.Locals("notifyRepository").(notification.NotifyRepository)
	if req.ParentID == nil && !isOwner {
		notificationRequest := notificationEntities.NotificationRequest{
			ActionType: "commented on your post",
			ActorID:    authID,
			EntityData: comment.Content,
			// EntityDataID: comment.ID,
			EntityID:   comment.ID,
			EntityType: "comment",
			NotifierID: comment.Post.UserID,
		}
		go notifyRepo.Notify(notificationRequest)
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

	go invalidateCommentCache(handler.rdb, int(comment.PostID))

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

	notifyRepo := c.Locals("notifyRepository").(notification.NotifyRepository)
	isOwner := comment.UserID == comment.Post.UserID
	if !isOwner {
		req := notificationEntities.NotificationRequest{
			ActorID:    comment.UserID,
			ActionType: "commented on your post",
			EntityData: comment.Content,
			EntityType: "comment",
			EntityID:   comment.ID,
			NotifierID: comment.Post.UserID,
		}
		go notifyRepo.DeleteNotification(req)
	}

	go invalidateCommentCache(handler.rdb, int(comment.PostID))

	return c.Status(fiber.StatusOK).JSON(comment)
}

func invalidateCommentCache(rdb *redis.Client, postID int) {
	key := fmt.Sprintf("comments-%d*", postID)
	invalidateCache(rdb, key)
}

func invalidateCache(rdb *redis.Client, key string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	pipe := rdb.Pipeline()
	match := fmt.Sprintf(key)
	iter := rdb.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		pipe.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		log.Println("REDIS ERR: ", err)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		log.Println("REDIS ERR: ", err)
	}
}
