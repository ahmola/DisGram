package internal

import (
	"services/comment-service/internal"

	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

func (r *commentRepository) Save(comment *internal.Comment) error {
	return r.DB.Create(comment).Error
}

func (r *commentRepository) FindByPostID(postID uint) ([]internal.Comment, error) {
	var comments []internal.Comment
	err := r.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
