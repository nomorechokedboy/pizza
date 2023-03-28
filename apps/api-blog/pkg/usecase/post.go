package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
)

type PostUsecase interface {
	GetAllPosts() ([]entities.Post, error)
	GetAllPostsByUserID(userID uint) ([]entities.Post, error)
	GetPostBySlug(slug string) (*entities.Post, error)
	GetPostsByParentID(parentID uint) ([]entities.Post, error)
	CreatePost(userID uint, body *entities.PostReq) (uint, error)
	UpdatePost(id uint, body *entities.PostReq) error
	DeletePost(id uint) error
}

type postUsecase struct {
	repo repository.PostRepository
}

func NewPostUseCase(repo repository.PostRepository) PostUsecase {
	return &postUsecase{repo: repo}
}

func (usecase *postUsecase) GetAllPosts() ([]entities.Post, error) {
	return usecase.repo.GetAllPosts()
}

func (usecase *postUsecase) GetAllPostsByUserID(userID uint) ([]entities.Post, error) {
	return usecase.repo.GetAllPostsByUserID(userID)
}

func (usecase *postUsecase) GetPostBySlug(slug string) (*entities.Post, error) {
	return usecase.repo.GetPostBySlug(slug)
}

func (usecase *postUsecase) CreatePost(userID uint, body *entities.PostReq) (uint, error) {
	post := &entities.Post{
		Title:       body.Title,
		ParentID:    body.ParentID,
		Content:     body.Content,
		PublishedAt: body.PublishedAt,
		UserID:      userID,
	}

	return usecase.repo.CreatePost(post)

}

func (usecase *postUsecase) GetPostsByParentID(parentID uint) ([]entities.Post, error) {
	return usecase.repo.GetAllPostsByParentID(parentID)
}

func (usecase *postUsecase) UpdatePost(id uint, body *entities.PostReq) error {
	post := &entities.Post{
		ID:          id,
		Title:       body.Title,
		ParentID:    body.ParentID,
		Content:     body.Content,
		PublishedAt: body.PublishedAt,
	}

	return usecase.repo.UpdatePost(post)
}

func (usecase *postUsecase) DeletePost(id uint) error {
	return usecase.repo.DeletePost(id)
}
