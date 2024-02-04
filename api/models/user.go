package models

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"gorm.io/gorm"
)

var AnonymousUser = &User{}

type User struct {
	gorm.Model
	Email    string `gorm:"column:email;not null;index;" json:"email" validate:"required"`
	Name     string `gorm:"column:name;size:64;not null;" json:"name" validate:"required"`
	Password string `gorm:"column:password;not null;" json:"-" validate:"required"`
}

func (m *User) IsAnonymous() bool {
	return m == AnonymousUser
}

func (m *User) BeforeCreate(_ *gorm.DB) (err error) {
	m.Password = HashPassword(m.Password)
	return
}

func (m *User) CheckPassword(plain string) (bool, error) {
	storedHashBytes, err := hex.DecodeString(m.Password)
	if err != nil {
		return false, err
	}
	userHash := sha256.Sum256([]byte(plain))
	return subtle.ConstantTimeCompare(userHash[:], storedHashBytes) == 1, nil
}

func HashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}
