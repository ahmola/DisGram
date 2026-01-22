package internal

import (
	"services/pkg/common"
)

type Comment struct {
	common.BaseEntity
	UserID  uint   `json:"userId" gorm:"not null"`
	PostID  uint   `json:"postId" gorm:"not null"`
	Content string `json:"content" gorm:"type:text;not null"`
}
