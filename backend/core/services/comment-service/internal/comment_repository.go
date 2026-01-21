package internal

import (
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func (r *CommentRepository) Save(comment *Comment) error {
	return r.DB.Create(comment).Error
}

func (r *CommentRepository) FindByPostID(postID uint) ([]Comment, error) {
	var comments []Comment
	err := r.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
