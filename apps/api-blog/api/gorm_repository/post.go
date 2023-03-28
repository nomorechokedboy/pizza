package gorm_repository

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"gorm.io/gorm"
)

type PostGormRepo struct {
	db *gorm.DB
}

func NewPostGormRepository(db *gorm.DB) repository.PostRepository {
	return &PostGormRepo{
		db: db,
	}
}

func (r *PostGormRepo) CreateSlug(slug *entities.Slug) error {
	return r.db.Create(&slug).Error
}

func (r *PostGormRepo) CreatePost(post *entities.Post) (uint, error) {
	createdPost := r.db.Create(&post)

	return post.ID, createdPost.Error
}

func (r *PostGormRepo) GetPostBySlug(slug string) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.Joins("JOIN slugs ON slugs.post_id = posts.id AND slugs.slug = ?", slug).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostGormRepo) GetAllPosts() ([]entities.Post, error) {
	var posts []entities.Post

	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostGormRepo) GetAllPostsByParentID(parentID uint) ([]entities.Post, error) {
	var posts []entities.Post

	if err := r.db.Find(&posts, "parent_id = ?", parentID).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostGormRepo) GetAllPostsByUserID(userID uint) ([]entities.Post, error) {
	var posts []entities.Post

	if err := r.db.Find(&posts, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostGormRepo) UpdatePost(post *entities.Post) error {
	return r.db.Model(&post).Updates(&entities.Post{
		Title:    post.Title,
		ParentID: post.ParentID,
		Content:  post.Content,
	}).Error
}

func (r *PostGormRepo) DeletePost(id uint) error {
	post := entities.Post{ID: id}

	return r.db.Delete(&post).Error
}
