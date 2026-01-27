package internal

import (
	"services/pkg/common"
)

type Post struct {
	common.BaseEntity
	UserID  uint   `json:"userId" gorm:"not null"`
	Caption string `json:"caption" gorm:"type:text"`
}

type PostImage struct {
	common.BaseEntity
	PostID    uint   `json:"postId"`
	FileKey   string `json:"fileKey" gorm:"not null;unique"`
	Extension string `json:"extension" gorm:"not null"`
	Url       string `json:"string" gorm:"not null"`
	Seq       uint   `json:"seq"`
}

type Like struct {
	common.BaseEntity
	PostID uint `json:"postId" gorm:"not null"`
	UserID uint `json:"userId" gorm:"not null"`
}
