package usecase

import (
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
)

type CommentUsecase interface {
	GetAllComments(query *entities.CommentQuery) (common.BasePaginationResponse[entities.Comment], error)
	CreateComment(userID uint, body *entities.CommentRequest) (*entities.Comment, error)
	UpdateComment(id uint, content string) (*entities.Comment, error)
	DeleteComment(id uint) (*entities.Comment, error)
}

type commentUsecase struct {
	repo repository.CommentRepository
}

func NewCommentUseCase(repo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{repo: repo}
}

func (usecase *commentUsecase) GetAllComments(query *entities.CommentQuery) (common.BasePaginationResponse[entities.Comment], error) {
	return usecase.repo.GetAllComments(query)
}

func (usecase *commentUsecase) CreateComment(userID uint, body *entities.CommentRequest) (*entities.Comment, error) {
	comment := &entities.Comment{
		UserID:   userID,
		PostID:   body.PostID,
		ParentID: body.ParentID,
		Content:  body.Content,
	}

	return usecase.repo.CreateComment(comment)

}

func (usecase *commentUsecase) UpdateComment(id uint, content string) (*entities.Comment, error) {
	comment := &entities.Comment{
		ID:      id,
		Content: content,
	}

	return usecase.repo.UpdateComment(comment)
}

func (usecase *commentUsecase) DeleteComment(id uint) (*entities.Comment, error) {
	return usecase.repo.DeleteComment(id)
}
