package repositories

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/lib"
)

type CommentRepository struct {
	Db lib.Database
}

func NewCommentRepository(db lib.Database) CommentRepository {
	return CommentRepository{
		Db: db,
	}
}

func (b CommentRepository) Query(params *dto.CommentQueryParams) (*models.Comments, *dto.Pagination, error) {
	db := b.Db.ORM.Preload("User").Preload("Post").Model(&models.Comments{})

	if params.Q != "" {
		db = db.Where("title = ?", params.Q)
	}

	if params.PaginationParams.PageSize == 0 {
		params.PaginationParams.PageSize = 5
	}

	list := make(models.Comments, 0)
	pagination, err := QueryPagination(db, params.PaginationParams, &list)
	if err != nil {
		return nil, nil, err
	}

	return &list, pagination, nil
}

func (b CommentRepository) Get(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := b.Db.ORM.Preload("User").Preload("Post").Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (b CommentRepository) Create(post *models.Comment) error {
	return nil
}

func (b CommentRepository) Delete(post *models.Comment) error {
	return b.Db.ORM.Delete(&post).Error
}
