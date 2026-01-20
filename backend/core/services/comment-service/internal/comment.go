package internal

import (
	"services/pkg/common"
)

type Comment struct {
	common.BaseEntity
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
	Content string `json:"content" gorm:"type:text"`
}
