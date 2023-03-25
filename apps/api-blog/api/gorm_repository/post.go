package gorm_repository

import (
	"api-blog/pkg/entities"
	"api-blog/pkg/repository"
	"time"

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
	return r.db.Create(slug).Error
}

func (r *PostGormRepo) CreatePost(post *entities.Post) (uint, error) {
	post.CreatedAt = time.Now()

	createdPost := r.db.Create(&post)

	return post.ID, createdPost.Error
}

func (r *PostGormRepo) GetPostBySlug(slug string) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.Joins("JOIN slug ON slug.postID = post.id AND slug.name = ?", slug).First(&post).Error; err != nil {
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

func (r *PostGormRepo) GetAllPostsByUserID(userID uint) ([]entities.Post, error) {
	var posts []entities.Post

	if err := r.db.Find(&posts, "userID = ?", userID).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostGormRepo) UpdatePost(post *entities.Post) error {
	return r.db.Where("id = ?", post.ID).Updates(entities.Post{
		Title:     post.Title,
		Content:   post.Content,
		UpdatedAt: time.Now(),
	}).Error
}

func (r *PostGormRepo) DeletePost(id uint) error {
	return r.db.Delete("id = ?", id).Error
}
