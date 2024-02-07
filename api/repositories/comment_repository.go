package repositories

import (
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/GapaiID/SE-challenge2/lib"
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

	db = db.Where(params.GetSearch(params.SearchFields()))
	db = db.Order(params.ParseOrderFilter(params.OrderFields()))
	params.SetDefaultPageSize(params.DefaultPageSize())

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

func (b CommentRepository) Create(comment *models.Comment) error {
	err := b.Db.ORM.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (b CommentRepository) Update(commentID uint, comment *models.Comment) error {
	err := b.Db.ORM.Where("id = ?", commentID).Updates(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (b CommentRepository) Delete(comment *models.Comment) error {
	return b.Db.ORM.Delete(&comment).Error
}
