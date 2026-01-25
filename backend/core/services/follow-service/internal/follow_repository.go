package internal

import (
	"services/pkg/common"
)

type FollowRepository struct {
	common.GormRepository[Follow]
}

func (r *FollowRepository) FindFolloweesByFollowerID(followerID uint) ([]*uint, error) {
	var follows []Follow
	err := r.DB.Model(&Follow{}).
		Where("follower_id = ?", followerID).
		Find(&follows).Error

	followees := make([]*uint, len(follows))
	for i, follow := range follows {
		followees[i] = &follow.FolloweeID
	}

	return followees, err
}

func (r *FollowRepository) FindFollowersByFolloweeID(followeeID uint) ([]uint, error) {
	var follows []Follow
	err := r.DB.Model(&Follow{}).
		Where("followee_id = ?", followeeID).
		Find(&follows).Error

	followers := make([]uint, len(follows))
	for i, follow := range follows {
		followers[i] = follow.FollowerID
	}

	return followers, err
}
