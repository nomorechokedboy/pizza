package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
)

type SlugUsecase interface {
	GetSlugCount(slug string) (int64, error)
	CreateSlug(postId uint, slug string) error
}

type slugUsecase struct {
	repo repository.SlugRepository
}

func NewSlugUseCase(repo repository.SlugRepository) SlugUsecase {
	return &slugUsecase{repo: repo}
}

func (usecase *slugUsecase) CreateSlug(postID uint, slug string) error {
	post_slug := &entities.Slug{
		Slug:   slug,
		PostID: postID,
	}

	return usecase.repo.CreateSlug(post_slug)
}

func (usecase *slugUsecase) GetSlugCount(slug string) (int64, error) {
	return usecase.repo.GetSlugCount(slug)
}
