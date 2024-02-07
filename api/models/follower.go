package models

import "gorm.io/gorm"

type Follower struct {
	gorm.Model

	UserID     uint `gorm:"column:user_id;not null;"`
	User       User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FollowerID uint `gorm:"column:follower_id;not null;"`
	Follower   User `gorm:"foreignKey:FollowerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
