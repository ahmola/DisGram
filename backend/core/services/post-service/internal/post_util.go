package internal

func PostRequestToOrigins(req PostRequest) (*Post, []*PostImage, error) {
	post := &Post{
		UserID:  req.UserID,
		Caption: req.Caption,
	}

	postImages := make([]*PostImage, len(req.PostImages))
	for _, image := range req.PostImages {
		postImage := &PostImage{
			PostID:    image.PostID,
			FileKey:   image.FileKey,
			Extension: image.Extension,
			Url:       image.Url,
			Seq:       image.Seq,
		}
		postImages = append(postImages, postImage)
	}

	return post, postImages, nil
}

func OriginsToResponse(post *Post, postImages []*PostImage) (*PostResponse, error) {
	imageUrls := make([]string, len(postImages))
	for _, image := range postImages {
		imageUrls = append(imageUrls, image.Url)
	}

	return &PostResponse{
		ID:            post.ID,
		UserID:        post.UserID,
		Caption:       post.Caption,
		PostImageUrls: imageUrls,
	}, nil
}

func LikeRequestToOrigin(likeRequest LikeRequest) (*Like, error) {
	return &Like{
		PostID: likeRequest.PostID,
		UserID: likeRequest.UserID,
	}, nil
}
