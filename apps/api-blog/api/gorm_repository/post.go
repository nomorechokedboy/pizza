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

		cond := entities.Post{UserID: query.UserID, ParentID: parentIDaddr, Published: true}

		if err := r.db.Scopes(scopes.Pagination(r.db, entities.Post{}, query.BaseQuery, &res)).
			Preload(clause.Associations).
			Find(&posts, cond).
			Error; err != nil {
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

	for i, post := range res.Items {
		res.Items[i].CommentCount = len(post.Comments)
		res.Items[i].ReactionCount = uint(len(post.Reactions))
	}

	return res, nil
}

func (r *PostGormRepo) GetPostBySlug(slug string) (*entities.Post, error) {
	var post entities.Post

	tx := r.db.
		Preload(clause.Associations).
		Joins("JOIN slugs ON slugs.post_id = posts.id AND slugs.slug = ?", slug).
		First(&post)
	if res := r.db.
		Preload(clause.Associations).
		Find(&post.Reactions); res.Error != nil {
		return nil, res.Error
	}

	post.CommentCount = len(post.Comments)
	post.ReactionCount = uint(len(post.Reactions))

	return &post, tx.Error
}

func (r *PostGormRepo) CreatePost(post *entities.Post) (*entities.Post, error) {
	tx := r.db.Create(&post)

	r.db.Preload(clause.Associations).First(&post)

	return post, tx.Error
}

func (r *PostGormRepo) UpdatePost(post *entities.Post) (*entities.Post, error) {
	tx := r.db.Model(&post).Clauses(clause.Returning{}).Updates(
		map[string]interface{}{
			"title":        post.Title,
			"parent_id":    post.ParentID,
			"slug":         post.Slug,
			"image":        post.Image,
			"content":      post.Content,
			"published":    post.Published,
			"published_at": post.PublishedAt,
		})

	r.db.Preload(clause.Associations).First(&post)

	return post, tx.Error
}

func (r *PostGormRepo) DeletePost(id uint) (*entities.Post, error) {
	post := entities.Post{ID: id}

	r.db.Preload(clause.Associations).First(&post)

	tx := r.db.
		Delete(&post).
		Model(entities.Post{ParentID: &id}).
		Update("parent_id", nil)

	return &post, tx.Error
}
