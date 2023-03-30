package gorm_repository

import (
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

func (r *PostGormRepo) GetAllPosts(userID uint, parentID uint, page int, pageSize int) ([]entities.Post, error) {
	var posts []entities.Post
	offset := (page - 1) * pageSize
	addrParentID := &parentID

	if parentID == 0 {
		addrParentID = nil
	}

	cond := entities.Post{UserID: userID, ParentID: addrParentID}

	if err := r.db.
		Offset(offset).
		Limit(pageSize).
		Order("id ASC").
		Preload(clause.Associations).
		Find(&posts, cond).Error; err != nil {
		return nil, err
	}

	return posts, nil
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
			"content":      post.Content,
			"published":    post.Published,
			"published_at": post.PublishedAt,
		}).Error
}

func (r *PostGormRepo) DeletePost(id uint) error {
	return r.db.
		Delete(&entities.Post{ID: id}).
		Model(&entities.Post{ParentID: &id}).
		Update("parent_id", nil).Error
}
