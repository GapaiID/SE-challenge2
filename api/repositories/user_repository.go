package repositories

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/lib"
)

type UserRepository struct {
	Db lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		Db: db,
	}
}

func (r UserRepository) Create(user *models.User) error {
	err := r.Db.ORM.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Update(user *models.User) error {
	err := r.Db.ORM.Model(&user).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Query(params *dto.UserQueryParams) (*models.Users, *dto.Pagination, error) {
	db := r.Db.ORM.Model(&models.Users{})

	db = db.Where(params.GetSearch(params.SearchFields()))
	db = db.Order(params.ParseOrderFilter(params.OrderFields()))
	params.SetDefaultPageSize(params.DefaultPageSize())

	list := make(models.Users, 0)
	pagination, err := QueryPagination(db, params.PaginationParams, &list)
	if err != nil {
		return nil, nil, err
	}

	return &list, pagination, nil
}

func (r UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.Db.ORM.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) Get(id uint) (*models.User, error) {
	var user models.User
	err := r.Db.ORM.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
