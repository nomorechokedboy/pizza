package handler

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type PostHandler struct{ usecase usecase.PostUsecase }

func NewPostHandler(usecase usecase.PostUsecase) *PostHandler {
	return &PostHandler{usecase: usecase}
}

// @GetAllPostsByUserID godoc
// @Summary Show all posts
// @Description get all posts from specified user or all if user is empty
// @Tags Posts
// @Param  userID query string false "User ID"
// @Success 200 {array} entities.PostRes{}
// @Failure 404
// @Failure 500
// @Router /posts [get]
func (handler *PostHandler) GetAllPostsByUserID(c *fiber.Ctx) error {
	userID := c.QueryInt("userID")
	var posts []entities.Post

	toResponse := func(posts []entities.Post) []entities.PostRes {
		postReponse := []entities.PostRes{}

		for _, post := range posts {
			postRes := entities.PostRes{
				ID:          post.ID,
				UserID:      post.UserID,
				ParentID:    post.ParentID,
				Title:       post.Title,
				Content:     post.Content,
				PublishedAt: post.PublishedAt,
				CreatedAt:   post.CreatedAt,
				UpdatedAt:   post.UpdatedAt,
			}

			postReponse = append(postReponse, postRes)
		}

		return postReponse

	}

	if userID == 0 {
		posts, err := handler.usecase.GetAllPosts()

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "failed to get all posts")
		}

		return c.Status(fiber.StatusOK).JSON(toResponse(posts))
	}

	posts, err := handler.usecase.GetAllPostsByUserID(uint(userID))

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to get all posts")
	}

	return c.Status(fiber.StatusOK).JSON(toResponse(posts))
}

// @GetPostByID godoc
// @Summary Get post
// @Description Get post by slug
// @Tags Posts
// @Param slug path string true "Post Slug"
// @Success 200 {object} entities.PostReq{}
// @Failure 400
// @Router /posts/{slug} [get]
func (handler *PostHandler) GetPostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	if slug == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post slug")
	}

	post, err := handler.usecase.GetPostBySlug(slug)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "failed to get post")
	}

	postRes := entities.PostRes{
		ID:          post.ID,
		UserID:      post.UserID,
		ParentID:    post.ParentID,
		Title:       post.Title,
		Content:     post.Content,
		PublishedAt: post.PublishedAt,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(postRes)
}

// @CreatePost godoc
// @Summary Create post
// @Description Create post
// @Tags Posts
// @Accept json
// @Param post body entities.PostReq true "Post"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/ [post]
func (handler *PostHandler) CreatePost(c *fiber.Ctx) error {
	authID := c.Locals("uId").(uint)
	req := new(entities.PostReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	post_slug := slug.Make(req.Title)

	if _, err := handler.usecase.GetSlug(post_slug); err == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "duplicate post title")
	}

	postID, err := handler.usecase.CreatePost(authID, req.Title, req.Content)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new post")
	}

	if err := handler.usecase.CreateSlug(postID, post_slug); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new slug")
	}

	return c.Status(fiber.StatusCreated).SendString("Created successfully")
}

// @UpdatePost godoc
// @Description Update post
// @Summary Update post with new info
// @Param id path int true "Post ID"
// @Param post body entities.PostReq true "Post"
// @Tags Posts
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/{id} [put]
func (handler *PostHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post ID")
	}

	req := new(entities.PostReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	post_slug := slug.Make(req.Title)

	if _, err := handler.usecase.GetSlug(post_slug); err == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "duplicate post title")
	}

	if err := handler.usecase.UpdatePost(uint(id), req.Title, req.Content); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update specfied post")
	}

	if err := handler.usecase.CreateSlug(uint(id), post_slug); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new slug")
	}

	return c.Status(fiber.StatusOK).SendString("Updated successfully")
}

// @DeletePost godoc
// @Description Delete post
// @Summary Delete specified post
// @Param id path int true "Post ID"
// @Tags Posts
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @security ApiKeyAuth
// @Router /posts/{id} [delete]
func (handler *PostHandler) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post ID")
	}

	if err := handler.usecase.DeletePost(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete specfied post")
	}

	return c.Status(fiber.StatusOK).SendString("Deleted successfully")
}
