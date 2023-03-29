package handler

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type PostHandler struct {
	usecase     usecase.PostUsecase
	slugUsecase usecase.SlugUsecase
	userUsecase usecase.UserUsecase
}

func NewPostHandler(usecase usecase.PostUsecase, slugUsecase usecase.SlugUsecase, userUsecase usecase.UserUsecase) *PostHandler {
	return &PostHandler{
		usecase:     usecase,
		slugUsecase: slugUsecase,
		userUsecase: userUsecase,
	}
}

// @GetAllPostsByUserID godoc
// @Summary Show all posts
// @Description get all posts from specified user or all if user is empty
// @Tags Posts
// @Param  userID query string false "User ID"
// @Param  parentID query string false "Parent ID"
// @Param  page query string false "Page"
// @Param  pageSize query string false "Page Size"
// @Success 200 {array} entities.PostRes{}
// @Failure 404
// @Failure 500
// @Router /posts [get]
func (handler *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	userID := c.QueryInt("userID")
	parentID := c.QueryInt("parentID")
	page := c.QueryInt("page")
	pageSize := c.QueryInt("pageSize")

	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	var posts []entities.Post

	toResponse := func(posts []entities.Post) entities.PostPaginationResponse {
		postReponse := []entities.PostRes{}

		for _, post := range posts {
			user, _ := handler.userUsecase.GetUserById(post.UserID)
			parent, _ := handler.usecase.GetPostByID(*post.ParentID)

			postRes := entities.PostRes{
				ID:          post.ID,
				User:        user,
				Parent:      parent,
				Slug:        post.Slug,
				Title:       post.Title,
				Content:     post.Content,
				Published:   post.Published,
				PublishedAt: post.PublishedAt,
				CreatedAt:   post.CreatedAt,
				UpdatedAt:   post.UpdatedAt,
			}

			postReponse = append(postReponse, postRes)
		}

		postPaginationReponse := entities.PostPaginationResponse{
			Posts:    postReponse,
			Page:     page,
			PageSize: pageSize,
			Total:    len(postReponse),
		}

		return postPaginationReponse
	}

	if userID == 0 && parentID == 0 {
		posts, err := handler.usecase.GetAllPosts(page, pageSize)

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "failed to get all posts")
		}

		return c.Status(fiber.StatusOK).JSON(toResponse(posts))
	}

	posts, err := handler.usecase.GetAllPostsByQuery(uint(userID), uint(parentID), page, pageSize)

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

	user, _ := handler.userUsecase.GetUserById(post.UserID)
	parent, _ := handler.usecase.GetPostByID(*post.ParentID)

	postRes := entities.PostRes{
		ID:          post.ID,
		User:        user,
		Parent:      parent,
		Slug:        post.Slug,
		Title:       post.Title,
		Content:     post.Content,
		Published:   post.Published,
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

	if userPosts, err := handler.usecase.GetAllPosts(0, 1); err == nil {
		for _, post := range userPosts {
			if post.Title == req.Title {
				return fiber.NewError(fiber.StatusConflict, "duplicate post title")
			}
		}
	}

	postSlug := slug.Make(req.Title)

	if slugCount, err := handler.slugUsecase.GetSlugCount(postSlug); slugCount != 0 && err == nil {
		postSlug = fmt.Sprintf("%s-%d", postSlug, slugCount)
	}

	postID, err := handler.usecase.CreatePost(authID, postSlug, req)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create new post")
	}

	if err := handler.slugUsecase.CreateSlug(postID, postSlug); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create post slug")
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

	postSlug := slug.Make(req.Title)

	if slugCount, err := handler.slugUsecase.GetSlugCount(postSlug); slugCount != 0 && err == nil {
		postSlug = fmt.Sprintf("%s-%d", postSlug, slugCount)
	}

	if err := handler.usecase.UpdatePost(uint(id), postSlug, req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update specfied post")
	}

	if err := handler.slugUsecase.CreateSlug(uint(id), postSlug); err != nil {
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
