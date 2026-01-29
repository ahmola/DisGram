package internal

import (
	pb "services/pkg/proto/post"
)

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

func grpcPostRequestToPostRequest(grpcReq pb.PostRequest) *PostRequest {
	images := make([]PostImageRequest, len(grpcReq.PostImages))
	for i, image := range grpcReq.PostImages {
		images[i] = PostImageRequest{
			PostID:    uint(image.PostId),
			FileKey:   image.FileKey,
			Extension: image.Extension,
			Url:       image.Url,
			Seq:       uint(image.Seq),
		}
	}
	return &PostRequest{
		UserID:     uint(grpcReq.UserId),
		Caption:    grpcReq.Caption,
		PostImages: images,
	}
}

func postResponseToGrpcPostResponse(res *PostResponse) *pb.PostResponse {
	return &pb.PostResponse{
		Id:            uint32(res.ID),
		UserId:        uint32(res.UserID),
		Caption:       res.Caption,
		PostImageUrls: res.PostImageUrls,
	}
}

func grpcLikeRequestToLikeRequest(grpcReq pb.LikeRequest) *LikeRequest {
	return &LikeRequest{
		PostID: uint(grpcReq.PostId),
		UserID: uint(grpcReq.UserId),
	}
}

func likeListResponseToGrpcLikeResponse(res []*LikeResponse) *pb.LikeListResponse {
	likes := make([]*pb.LikeResponse, len(res))
	for i, like := range res {
		likes[i] = &pb.LikeResponse{
			Id:     uint32(like.ID),
			PostId: uint32(like.PostID),
			UserId: uint32(like.UserID),
		}
	}

	return &pb.LikeListResponse{
		Likes: likes,
	}
}
