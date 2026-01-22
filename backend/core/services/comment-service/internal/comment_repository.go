package internal

import (
	"services/pkg/common"
)

type CommentRepository struct {
	common.GormRepository[Comment]
}

func (r *CommentRepository) FindByPostID(postID uint) ([]Comment, error) {
	var comments []Comment
	err := r.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
