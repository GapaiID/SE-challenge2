package repositories

import (
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
	err := r.Db.ORM.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
