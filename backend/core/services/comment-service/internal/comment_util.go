package internal

import (
	// prevent namespace conflict
	pb "services/pkg/proto/comment"
)

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
func grpcRequestToRequest(req pb.CommentRequest) *CommentRequest {
	res := &CommentRequest{
		UserID:  uint(req.UserId),
		PostID:  uint(req.PostId),
		Content: req.Content,
	}

	return res
}

func responseToGrpcResponse(res *CommentResponse) *pb.CommentResponse {
	grpcRes := &pb.CommentResponse{
		UserId:    uint32(res.UserID),
		PostId:    uint32(res.PostID),
		Content:   res.Content,
		CreatedAt: res.CreatedAt,
	}

	return grpcRes
}

func commentListToGrpcResponse(comments []*CommentResponse) *pb.CommentListResponse {
	grpcRes := &pb.CommentListResponse{}

	for _, comment := range comments {
		grpcRes.Comments = append(grpcRes.Comments, &pb.CommentResponse{
			Id:        uint32(comment.ID),
			UserId:    uint32(comment.UserID),
			PostId:    uint32(comment.PostID),
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		})

	}

	return grpcRes
}
