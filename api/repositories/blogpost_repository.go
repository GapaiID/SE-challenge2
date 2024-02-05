package repositories

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/lib"
)

type BlogPostRepository struct {
	Db lib.Database
}

func NewBlogPostRepository(db lib.Database) BlogPostRepository {
	return BlogPostRepository{
		Db: db,
	}
}

func (b BlogPostRepository) Query(params *dto.BlogPostQueryParams) (*models.BlogPosts, *dto.Pagination, error) {
	db := b.Db.ORM.Preload("User").Model(&models.BlogPosts{})

	db = db.Where(params.GetSearch(params.SearchFields()))
	db = db.Order(params.ParseOrderFilter(params.OrderFields()))
	params.SetDefaultPageSize(params.DefaultPageSize())

	list := make(models.BlogPosts, 0)
	pagination, err := QueryPagination(db, params.PaginationParams, &list)
	if err != nil {
		return nil, nil, err
	}

	return &list, pagination, nil
}

func (b BlogPostRepository) Get(id uint) (*models.BlogPost, error) {
	var post models.BlogPost
	err := b.Db.ORM.Preload("User").Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (b BlogPostRepository) Create(post *models.BlogPost) error {
	err := b.Db.ORM.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BlogPostRepository) Update(postID uint, post *models.BlogPost) error {
	err := b.Db.ORM.Where("id = ?", postID).Updates(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BlogPostRepository) Delete(post *models.BlogPost) error {
	return b.Db.ORM.Delete(&post).Error
}

func (b BlogPostRepository) QueryByFollowing(user *models.User, params *dto.BlogPostQueryParams) (*models.BlogPosts, *dto.Pagination, error) {
	db := b.Db.ORM.Model(&models.BlogPosts{}).Preload("User").Where(
		"user_id IN (SELECT user_id FROM followers WHERE follower_id = ?)", user.ID,
	)

	if params.Q != "" {
		db = db.Where("posts.title = ?", params.Q)
	}

	if params.PaginationParams.PageSize == 0 {
		params.PaginationParams.PageSize = 5
	}

	list := make(models.BlogPosts, 0)
	pagination, err := QueryPagination(db, params.PaginationParams, &list)
	if err != nil {
		return nil, nil, err
	}

	return &list, pagination, nil
}
