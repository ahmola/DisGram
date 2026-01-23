package internal

import (
	"services/pkg/common"
)

type UserRepository struct {
	common.GormRepository[User]
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	user := &User{}
	err := r.DB.Where("username = ?", username).First(&user).Error

	return user, err
}

func (r *UserRepository) existsByUsername(username string) bool {
	count := r.DB.Where("username = ?", username).Find(&User{})

	return count.RowsAffected > 0
}

func (r *UserRepository) existsById(ID uint) bool {
	var exists bool
	r.DB.Model(&User{}).
		Select("count(*) > 0").
		Where("id = ? ", ID).
		Find(&exists)

	return exists
}
