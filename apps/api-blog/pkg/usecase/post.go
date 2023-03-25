package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
)

type PostUsecase interface {
	GetAllPosts() ([]entities.Post, error)
	GetAllPostsByUserID(userID uint) ([]entities.Post, error)
	GetPostBySlug(slug string) (*entities.Post, error)
	CreateSlug(postId uint, title string) error
	CreatePost(userID uint, title string, content string) (uint, error)
	UpdatePost(id uint, title string, content string) error
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

func (usecase *postUsecase) CreateSlug(postID uint, title string) error {
	slug := &entities.Slug{
		Slug:   title,
		PostID: postID,
	}

	return usecase.repo.CreateSlug(slug)

}

func (usecase *postUsecase) CreatePost(userID uint, title string, content string) (uint, error) {
	post := &entities.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	return usecase.repo.CreatePost(post)

}

func (usecase *postUsecase) UpdatePost(id uint, title string, content string) error {
	post := &entities.Post{
		ID:      id,
		Title:   title,
		Content: content,
	}

	return usecase.repo.UpdatePost(post)
}

func (usecase *postUsecase) DeletePost(id uint) error {
	return usecase.repo.DeletePost(id)
}
