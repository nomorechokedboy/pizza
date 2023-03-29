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

func (r *PostGormRepo) GetPostByID(id uint) (*entities.Post, error) {
	post := entities.Post{ID: id}
	foundPost := r.db.First(&post)

	return &post, foundPost.Error
}

func (r *PostGormRepo) GetPostBySlug(slug string) (*entities.Post, error) {
	var post entities.Post

	if err := r.db.Joins("JOIN slugs ON slugs.post_id = posts.id AND slugs.slug = ?", slug).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostGormRepo) GetAllPosts(page int, pageSize int) ([]entities.Post, error) {
	var posts []entities.Post
	offset := (page - 1) * pageSize

	if err := r.db.Offset(offset).Limit(pageSize).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostGormRepo) GetAllPostsByQuery(userID uint, parentID uint, page int, pageSize int) ([]entities.Post, error) {
	var posts []entities.Post
	offset := (page - 1) * pageSize
	query := &entities.Post{UserID: userID, ParentID: &parentID}

	if parentID == 0 {
		query = &entities.Post{UserID: userID}
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&posts, query).Error; err != nil {
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
	return r.db.Model(&post).Updates(
		&entities.Post{
			Title:    post.Title,
			ParentID: post.ParentID,
			Slug:     post.Slug,
			Content:  post.Content,
		}).Updates(
		map[string]interface{}{
			"published":    post.Published,
			"published_at": post.PublishedAt,
		}).Error
}

func (r *PostGormRepo) DeletePost(id uint) error {
	return r.db.Delete(&entities.Post{ID: id}).Update("parent_id", nil).Error
}
