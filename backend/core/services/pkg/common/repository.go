package common

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	// create
	Create(entity *T) error
	// read
	FindByID(id uint) (*T, error)
	FindAll() ([]*T, error)
	// update
	Update(entity *T) error
	// delete
	Delete(id uint) error
}

type GormRepository[T any] struct {
	DB *gorm.DB
}

func (r *GormRepository[T]) Save(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *GormRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	return &entity, err
}

func (r *GormRepository[T]) FindAll() ([]*T, error) {
	var entities []*T
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}
