package handler

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"

	"fmt"

	"context"
	"crypto/sha256"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/minio/minio-go/v7"
)

type PostHandler struct {
	usecase     usecase.PostUsecase
	slugUsecase usecase.SlugUsecase
	userUsecase usecase.UserUsecase
	minioClient minio.Client
}

func NewPostHandler(usecase usecase.PostUsecase, slugUsecase usecase.SlugUsecase, userUsecase usecase.UserUsecase, mionioClient *minio.Client) *PostHandler {
	return &PostHandler{
		usecase:     usecase,
		slugUsecase: slugUsecase,
		userUsecase: userUsecase,
		minioClient: *mionioClient,
	}
}

// @GetAllPosts godoc
// @Summary Show all posts
// @Description get all posts
// @Tags Posts
// @Param  userID query int false "User ID"
// @Param  parentID query int false "Parent ID"
// @Param  page query int false "Page"
// @Param  pageSize query int false "Page Size"
// @Param sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param sortBy query string false "Sort by" Enums(id, title, slug, user_id, parent_id) default(id)
// @Success 200 {object} common.BasePaginationResponse[entities.PostResponse]
// @Failure 404
// @Failure 500
// @Router /posts [get]
func (handler *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	query := new(entities.PostQuery)

	if err := c.QueryParser(query); err != nil {
		return err
	}

	posts, err := handler.usecase.GetAllPosts(&entities.PostQuery{
		UserID:   uint(query.UserID),
		ParentID: uint(query.ParentID),
		BaseQuery: common.BaseQuery{
			Page:     query.Page,
			PageSize: query.PageSize,
			Sort:     query.Sort,
			SortBy:   query.SortBy,
		},
	})

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to get all posts")
	}

	postRes := common.BasePaginationResponse[entities.PostResponse]{
		Items: []entities.PostResponse{},
	}
	postRes.Page = posts.Page
	postRes.PageSize = posts.PageSize
	postRes.Total = posts.Total

	for _, post := range posts.Items {
		postRes.Items = append(postRes.Items, post.ToResponse())
	}

	return c.Status(fiber.StatusOK).JSON(postRes)
}

// @GetPostByID godoc
// @Summary Get post
// @Description Get post by slug
// @Tags Posts
// @Param slug path string true "Post Slug"
// @Success 200 {object} entities.PostResponse
// @Failure 400
// @Failure 404
// @Router /posts/{slug} [get]
func (handler *PostHandler) GetPostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	post, err := handler.usecase.GetPostBySlug(slug)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to get post")
	}

	return c.Status(fiber.StatusOK).JSON(post.ToResponse())
}

// @CreatePost godoc
// @Summary Create post
// @Description Create post
// @Tags Posts
// @Accept json
// @Param post body entities.PostRequest true "Post"
// @Success 201 {object} entities.PostResponse
// @Failure 400
// @Failure 409
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/ [post]
func (handler *PostHandler) CreatePost(c *fiber.Ctx) error {
	authID := c.Locals("uId").(uint)
	req := new(entities.PostRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if userPosts, err := handler.usecase.GetAllPosts(nil); err == nil {
		for _, post := range userPosts.Items {
			if post.Title == req.Title {
				return fiber.NewError(fiber.StatusConflict, "duplicate post title")
			}
		}
	}

	postSlug := slug.Make(req.Title)

	if slugCount, err := handler.slugUsecase.GetSlugCount(postSlug); slugCount != 0 && err == nil {
		postSlug = fmt.Sprintf("%s-%d", postSlug, slugCount)
	}

	post, err := handler.usecase.CreatePost(authID, postSlug, req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new post")
	}

	if err := handler.slugUsecase.CreateSlug(post.ID, postSlug); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create post slug")
	}

	return c.Status(fiber.StatusCreated).JSON(post.ToResponse())
}

// @UpdatePost godoc
// @Description Update post
// @Summary Update post with new info
// @Param id path int true "Post ID"
// @Param post body entities.PostRequest true "Post"
// @Tags Posts
// @Success 200 {object} entities.PostResponse
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/{id} [put]
func (handler *PostHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post ID")
	}

	req := new(entities.PostRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	postSlug := slug.Make(req.Title)

	if slugCount, err := handler.slugUsecase.GetSlugCount(postSlug); slugCount != 0 && err == nil {
		postSlug = fmt.Sprintf("%s-%d", postSlug, slugCount)
	}

	post, err := handler.usecase.UpdatePost(uint(id), postSlug, req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update specfied post")
	}

	if err := handler.slugUsecase.CreateSlug(uint(id), postSlug); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new slug")
	}

	return c.Status(fiber.StatusOK).JSON(post.ToResponse())
}

// @DeletePost godoc
// @Description Delete post
// @Summary Delete specified post
// @Param id path int true "Post ID"
// @Tags Posts
// @Produce json
// @Success 200 {object} entities.PostResponse
// @Failure 400
// @Failure 500
// @security ApiKeyAuth
// @Router /posts/{id} [delete]
func (handler *PostHandler) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post ID")
	}

	post, err := handler.usecase.DeletePost(uint(id))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete specfied post")
	}

	return c.Status(fiber.StatusOK).JSON(post.ToResponse())
}

// @GetPostAudio godoc
// @Description Convert post to speech
// @Summary Convert post to speech
// @Tags Posts
// @Param content path string true "Content"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /posts/t2s/{content} [get]
func (handler *PostHandler) GetPostAudio(c *fiber.Ctx) error {
	ctx := context.Background()
	content := c.Params("content")

	if content == "" {
		return fiber.NewError(fiber.StatusBadRequest, "empty or invalid post content")
	}

	hash := sha256.New()
	hash.Write([]byte(content))
	objectName := fmt.Sprintf("%x", hash.Sum(nil))

	if _, err := handler.minioClient.StatObject(ctx, "audio", objectName, minio.StatObjectOptions{}); err != nil {
		url := "https://api.voicerss.org/?key=817e51130c864a4ab0d6558d46cbee24&hl=en-us&c=MP3&src=" + content
		res, _ := http.Get(url)

		if exists, _ := handler.minioClient.BucketExists(ctx, "audio"); !exists {
			handler.minioClient.MakeBucket(ctx, "audio", minio.MakeBucketOptions{})
		}

		defer res.Body.Close()

		if _, err := handler.minioClient.PutObject(
			ctx, "audio", objectName, res.Body, res.ContentLength,
			minio.PutObjectOptions{ContentType: "audio/mpeg"}); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	audioObject, err := handler.minioClient.GetObject(ctx, "audio", objectName, minio.GetObjectOptions{})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Set("Content-Type", "audio/mpeg")

	return c.SendStream(audioObject)
}
