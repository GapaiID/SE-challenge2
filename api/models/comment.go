package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	Body   string   `gorm:"column:body;not null;"`
	UserID uint     `gorm:"column:user_id;not null;"`
	User   User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostID uint     `gorm:"column:post_id;not null;"`
	Post   BlogPost `gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Comments []Comment
