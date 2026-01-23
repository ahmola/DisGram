package internal

import (
	"services/pkg/common"
)

type User struct {
	common.BaseEntity
	Username     string `json:"username" gorm:"unique;not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"passwordHash" gorm:"not null"`
	Bio          string `json:"bio" gorm:"type:text"`
	AvatarUrl    string `json:"avatarUrl" gorm:"unique"`
}
