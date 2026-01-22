package internal

import (
	"services/pkg/proto/comment"
	"strconv"
)

func ConvertStringToUint(str string) (uint, error) {
	convPostID, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(convPostID), nil
}

func OriginToResponse(comment Comment) (*CommentResponse, error) {
	res := &CommentResponse{
		UserID:    comment.UserID,
		PostID:    comment.PostID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.String(),
	}

	return res, nil
}

func RequestToOrigin(req CommentRequest) (*Comment, error) {
	res := &Comment{
		UserID:  req.UserID,
		PostID:  req.PostID,
		Content: req.Content,
	}

	return res, nil
}

// grpc
func grpcRequestToRequest(req comment.CommentRequest) *CommentRequest {
	res := &CommentRequest{
		UserID:  uint(req.UserId),
		PostID:  uint(req.PostId),
		Content: req.Content,
	}

	return res
}

func responseToGrpcResponse(res *CommentResponse) *comment.CommentResponse {
	grpcRes := &comment.CommentResponse{
		UserId:    uint32(res.UserID),
		PostId:    uint32(res.PostID),
		Content:   res.Content,
		CreatedAt: res.CreatedAt,
	}

	return grpcRes
}
