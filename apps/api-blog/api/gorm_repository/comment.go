package gorm_repository

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"

	"gorm.io/gorm"
)

type CommentGormRepo struct {
	db *gorm.DB
}

func NewCommentGormRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentGormRepo{
		db: db,
	}
}

func (r *CommentGormRepo) GetAllComments() ([]entities.Comment, error) {
	var comments []entities.Comment

	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentGormRepo) GetAllCommentsByQuery(userID uint, postID uint, parentID uint) ([]entities.Comment, error) {
	var comments []entities.Comment
	query := &entities.Comment{UserID: userID, PostID: postID, ParentID: &parentID}

	if parentID == 0 {
		query = &entities.Comment{UserID: userID, PostID: postID}
	}

	if err := r.db.Where(&query).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentGormRepo) CreateComment(comment *entities.Comment) error {
	return r.db.Create(&comment).Error
}

func (r *CommentGormRepo) UpdateComment(comment *entities.Comment) error {
	return r.db.Model(&comment).Update("content", comment.Content).Error
}

func (r *CommentGormRepo) DeleteComment(id uint) error {
	return r.db.Where("parent_id = ? OR id = ?", id, id).Delete(&entities.Comment{}).Error
}
