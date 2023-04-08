package gorm_repository

import (
	"api-blog/api/scopes"
	"api-blog/pkg/common"
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostGormRepo struct {
	db *gorm.DB
}

func NewPostGormRepository(db *gorm.DB) repository.PostRepository {
	return &PostGormRepo{
		db: db,
	}
}

func (r *PostGormRepo) GetAllPosts(query *entities.PostQuery) (common.BasePaginationResponse[entities.Post], error) {
	var posts []entities.Post
	res := common.BasePaginationResponse[entities.Post]{}

	if query != nil {
		parentIDaddr := &query.ParentID

		if query.ParentID == 0 {
			parentIDaddr = nil
		}

		cond := entities.Post{UserID: query.UserID, ParentID: parentIDaddr}

		if err := r.db.Scopes(scopes.Pagination(r.db, entities.Post{}, query.BaseQuery, &res)).
			Preload(clause.Associations).
			Find(&posts, cond).Error; err != nil {
			return res, err
		}
	} else {
		if err := r.db.
			Preload(clause.Associations).
			Find(&posts).Error; err != nil {
			return res, err
		}
	}

	res.Items = posts

	return res, nil
}

func (r *PostGormRepo) GetPostBySlug(slug string) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.
		Preload(clause.Associations).
		Joins("JOIN slugs ON slugs.post_id = posts.id AND slugs.slug = ?", slug).
		First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostGormRepo) CreatePost(post *entities.Post) (uint, error) {
	createdPost := r.db.Create(&post)

	return post.ID, createdPost.Error
}

func (r *PostGormRepo) UpdatePost(post *entities.Post) error {
	return r.db.Model(&post).Updates(
		map[string]interface{}{
			"title":        post.Title,
			"parent_id":    post.ParentID,
			"slug":         post.Slug,
			"image":        post.Image,
			"content":      post.Content,
			"published":    post.Published,
			"published_at": post.PublishedAt,
		}).Error
}

func (r *PostGormRepo) DeletePost(id uint) error {
	return r.db.
		Delete(entities.Post{ID: id}).
		Model(entities.Post{ParentID: &id}).
		Update("parent_id", nil).Error
}
