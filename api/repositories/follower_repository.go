package repositories

import (
	"errors"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/lib"
)

type FollowerRepository struct {
	Db lib.Database
}

func NewFollowerRepository(db lib.Database) FollowerRepository {
	return FollowerRepository{
		Db: db,
	}
}

func (r FollowerRepository) HasFollowing(userID uint, followingID uint) error {
	var follower models.Follower

	if err := r.Db.ORM.Where("user_id = ?", followingID).Where("follower_id = ?", userID).First(&follower).Error; err != nil {
		return nil
	}
	return errors.New("already followed")
}

func (r FollowerRepository) Create(follower *models.Follower) error {
	if err := r.Db.ORM.Create(follower).Error; err != nil {
		return err
	}
	return nil
}
