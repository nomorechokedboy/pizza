package gorm_repository

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentGormRepo struct {
	db *gorm.DB
}

func NewCommentGormRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentGormRepo{
		db: db,
	}
}

func (r *CommentGormRepo) GetAllComments(userID uint, postID uint, parentID uint, page int, pageSize int) ([]entities.Comment, error) {
	var comments []entities.Comment
	offset := (page - 1) * pageSize
	addrParentID := &parentID

	if parentID == 0 {
		addrParentID = nil
	}

	cond := &entities.Comment{UserID: userID, PostID: postID, ParentID: addrParentID}

	if err := r.db.Debug().
		Offset(offset).
		Limit(pageSize).
		Order("id ASC").
		Preload(clause.Associations).
		Find(&comments, cond).Error; err != nil {
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
