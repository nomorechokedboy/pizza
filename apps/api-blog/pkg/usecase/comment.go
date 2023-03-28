package usecase

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
)

type CommentUsecase interface {
	GetAllComments() ([]entities.Comment, error)
	GetAllCommentsByQuery(userID uint, postID uint, parentID uint) ([]entities.Comment, error)
	CreateComment(userID uint, body *entities.CommentReq) error
	UpdateComment(id uint, content string) error
	DeleteComment(id uint) error
}

type commentUsecase struct {
	repo repository.CommentRepository
}

func NewCommentUseCase(repo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{repo: repo}
}

func (usecase *commentUsecase) GetAllComments() ([]entities.Comment, error) {
	return usecase.repo.GetAllComments()
}

func (usecase *commentUsecase) GetAllCommentsByQuery(userID uint, postID uint, parentID uint) ([]entities.Comment, error) {
	return usecase.repo.GetAllCommentsByQuery(userID, postID, parentID)
}

func (usecase *commentUsecase) CreateComment(userID uint, body *entities.CommentReq) error {
	comment := &entities.Comment{
		UserID:   userID,
		PostID:   body.PostID,
		ParentID: body.ParentID,
		Content:  body.Content,
	}

	return usecase.repo.CreateComment(comment)

}

func (usecase *commentUsecase) UpdateComment(id uint, content string) error {
	comment := &entities.Comment{
		ID:      id,
		Content: content,
	}

	return usecase.repo.UpdateComment(comment)
}

func (usecase *commentUsecase) DeleteComment(id uint) error {
	return usecase.repo.DeleteComment(id)
}
