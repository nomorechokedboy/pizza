package usecase

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"time"
)

type PostUsecase interface {
	GetAllPosts(query *entities.PostQuery) (common.BasePaginationResponse[entities.Post], error)
	GetPostBySlug(slug string) (*entities.Post, error)
	CreatePost(userID uint, postSlug string, body *entities.PostReq) (*entities.Post, error)
	UpdatePost(id uint, postSlug string, body *entities.PostReq) error
	DeletePost(id uint) error
}

type postUsecase struct {
	repo repository.PostRepository
}

func NewPostUseCase(repo repository.PostRepository) PostUsecase {
	return &postUsecase{repo: repo}
}

func (usecase *postUsecase) GetAllPosts(query *entities.PostQuery) (common.BasePaginationResponse[entities.Post], error) {
	return usecase.repo.GetAllPosts(query)
}

func (usecase *postUsecase) GetPostBySlug(slug string) (*entities.Post, error) {
	return usecase.repo.GetPostBySlug(slug)
}

func (usecase *postUsecase) CreatePost(userID uint, postSlug string, body *entities.PostReq) (*entities.Post, error) {
	var publishedAt *time.Time = nil

	if body.Published {
		t := time.Now()
		publishedAt = &t
	}

	post := &entities.Post{
		UserID:      userID,
		Title:       body.Title,
		ParentID:    body.ParentID,
		Image:       body.Image,
		Slug:        postSlug,
		Content:     body.Content,
		Published:   body.Published,
		PublishedAt: publishedAt,
	}

	return usecase.repo.CreatePost(post)

}

func (usecase *postUsecase) UpdatePost(id uint, postSlug string, body *entities.PostReq) error {
	var publishedAt *time.Time = nil

	if body.Published {
		t := time.Now()
		publishedAt = &t
	}

	post := &entities.Post{
		ID:          id,
		Title:       body.Title,
		ParentID:    body.ParentID,
		Slug:        postSlug,
		Content:     body.Content,
		Published:   body.Published,
		PublishedAt: publishedAt,
	}

	return usecase.repo.UpdatePost(post)
}

func (usecase *postUsecase) DeletePost(id uint) error {
	return usecase.repo.DeletePost(id)
}
