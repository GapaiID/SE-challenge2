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

func (b CommentRepository) Query(params *dto.CommentQueryParams) (*models.Comment, *dto.Pagination, error) {
	return nil, nil, nil
}

func (b CommentRepository) Get(id uint) (*models.Comment, error) {
	return nil, nil
}

func (b CommentRepository) Create(post *models.Comment) error {
	return nil
}

func (b CommentRepository) Delete(post *models.Comment) error {
	return nil
}
