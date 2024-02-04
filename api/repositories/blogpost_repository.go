package repositories

import (
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

func (b BlogPostRepository) Get(id uint) (*models.BlogPost, error) {
	var post models.BlogPost
	err := b.Db.ORM.Preload("User").Where("id = ?", id).Find(&post).Error
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

func (b BlogPostRepository) Delete(post *models.BlogPost) error {
	return b.Db.ORM.Delete(&post).Error
}
