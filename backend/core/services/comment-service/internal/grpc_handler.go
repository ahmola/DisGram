package internal

import (
	"context"
	"services/pkg/proto/comment"
)

// gRPC Server Interface
type CommentGrpcHandler struct {
	comment.UnimplementedCommetnServiceServer
	Svc *CommentService // buisiness logic
}

// Implementaions
func (h *CommentGrpcHandler) CreateComment(ctx context.Context, grpcReq *comment.CommentRequest) (*comment.CommentResponse, error) {
	req := grpcRequestToRequest(*grpcReq)
	res, err := h.Svc.CreateComment(*req)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)
	return grpcRes, nil
}

func (h *CommentGrpcHandler) GetCommentsByPostID(ctx context.Context, req *comment.CommentRequestByPostID) (*comment.CommentListResponse, error) {

}

func (h *CommentGrpcHandler) GetComment(ctx context.Context, req *comment.CommentRequestByID) (*comment.CommentResponse, error) {

}

func (h *CommentGrpcHandler) UpdateComment(ctx context.Context, req *comment.CommentRequest) (*comment.CommentResponse, error) {

}

func (h *CommentGrpcHandler) DeleteComment(ctx context.Context, req *comment.CommentRequestByID) (*comment.DeleteCommentResponse, error) {

}
