package internal

import (
	"services/pkg/common"
)

type PostRepository struct {
	postRepo  common.GormRepository[Post]
	imageRepo common.GormRepository[PostImage]
	likeRepo  common.GormRepository[Like]
}

func (r *PostRepository) FindAllImagesByPostID(postID uint) ([]*PostImage, error) {
	var images []*PostImage
	err := r.imageRepo.DB.Where("post_id = ?", postID).Find(images).Error
	return images, err
}

func (r *PostRepository) DeleteAllImagesByPostID(postID uint) error {
	err := r.imageRepo.DB.Delete("post_id = ?", postID).Error
	return err
}
