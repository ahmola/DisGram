package internal

import (
	"services/pkg/common"
)

type Comment struct {
	common.BaseEntity
	UserID  uint   `json:"userId" gorm:"not null"`            // User ID
	PostID  uint   `json:"postId" gorm:"not null"`            // Post ID
	Content string `json:"content" gorm:"type:text;not null"` // Content of Comment
}
