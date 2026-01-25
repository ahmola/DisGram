package internal

import (
	"services/pkg/common"
)

type Follow struct {
	common.BaseEntity
	FolloweeID uint `json:"followee_id" gorm:"not null"`
	FollowerID uint `json:"follower_id" gorm:"not null"`
}
