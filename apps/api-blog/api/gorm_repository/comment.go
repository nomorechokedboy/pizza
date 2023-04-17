package gorm_repository

import (
	"api-blog/api/scopes"
	"api-blog/pkg/common"
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

func (r *CommentGormRepo) GetAllComments(query *entities.CommentQuery) (common.BasePaginationResponse[entities.Comment], error) {
	var comments []entities.Comment
	res := common.BasePaginationResponse[entities.Comment]{}

	parentIDaddr := &query.ParentID

	if query.ParentID == 0 {
		parentIDaddr = nil
	}

	cond := &entities.Comment{UserID: query.UserID, PostID: query.PostID, ParentID: parentIDaddr}

	if err := r.db.Scopes(scopes.Pagination(r.db, entities.Comment{}, query.BaseQuery, &res)).
		Preload("Replies.User").
		Preload("Replies.Replies").
		Preload(clause.Associations).
		Find(&comments, cond).Error; err != nil {
		return res, err
	}

	commentFilter := []entities.Comment{}

	for _, item := range comments {
		if len(item.Replies) > 0 {
			commentFilter = append(commentFilter, item)
		}
	}

	res.Items = commentFilter

	return res, nil
}

func (r *CommentGormRepo) CreateComment(comment *entities.Comment) (*entities.Comment, error) {
	tx := r.db.Create(&comment)

	r.db.Preload(clause.Associations).First(&comment)

	return comment, tx.Error
}

func (r *CommentGormRepo) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	tx := r.db.Model(&comment).Update("content", comment.Content)

	r.db.Preload(clause.Associations).First(&comment)

	return comment, tx.Error
}

func (r *CommentGormRepo) DeleteComment(id uint) (*entities.Comment, error) {
	comment := entities.Comment{ID: id}

	r.db.Preload(clause.Associations).First(&comment)

	tx := r.db.Delete(&entities.Comment{}, "id = ? OR parent_id = ?", id, id)

	return &comment, tx.Error
}
