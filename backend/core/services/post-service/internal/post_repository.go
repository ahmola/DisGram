package internal

import (
	"services/pkg/common"
)

type PostRepository struct {
	PostRepo  common.GormRepository[Post]
	ImageRepo common.GormRepository[PostImage]
	LikeRepo  common.GormRepository[Like]
}

func (r *PostRepository) FindAllImagesByPostID(postID uint) ([]*PostImage, error) {
	var images []*PostImage
	err := r.ImageRepo.DB.Where("post_id = ?", postID).Find(images).Error
	return images, err
}

func (r *PostRepository) DeleteAllImagesByPostID(postID uint) error {
	err := r.ImageRepo.DB.Delete("post_id = ?", postID).Error
	return err
}

func (r *PostRepository) FindAllLikesByPostID(postID uint) ([]*Like, error) {
	var likes []*Like
	err := r.LikeRepo.DB.Where("post_id = ?", postID).Find(likes).Error
	return likes, err
}
