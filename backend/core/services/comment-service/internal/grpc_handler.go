package internal

import (
	"context"
	pb "services/pkg/proto/comment"
)

// gRPC Server Interface
type CommentGrpcHandler struct {
	pb.UnimplementedCommentServiceServer
	Svc *CommentService // buisiness logic
}

// Implementaions
func (h *CommentGrpcHandler) CreateComment(ctx context.Context, grpcReq *pb.CommentRequest) (*pb.CommentResponse, error) {
	req := grpcRequestToRequest(*grpcReq)
	res, err := h.Svc.CreateComment(*req)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)
	return grpcRes, nil
}

func (h *CommentGrpcHandler) GetCommentsByPostID(ctx context.Context, grpcReq *pb.CommentRequestById) (*pb.CommentListResponse, error) {
	res, err := h.Svc.GetCommentsByPostID(grpcReq.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := commentListToGrpcResponse(res)

	return grpcRes, nil
}

func (h *CommentGrpcHandler) GetComment(ctx context.Context, req *pb.CommentRequestById) (*pb.CommentResponse, error) {
	res, err := h.Svc.GetCommentByID(req.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)

	return grpcRes, nil
}

func (h *CommentGrpcHandler) UpdateComment(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	grpcReq := grpcRequestToRequest(*req)
	res, err := h.Svc.UpdateComment(*grpcReq)

	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)

	return grpcRes, nil
}

func (h *CommentGrpcHandler) DeleteComment(ctx context.Context, req *pb.CommentRequestById) (*pb.DeleteCommentResponse, error) {
	res, err := h.Svc.DeleteComment(req.Id)

	if err != nil {
		return nil, err
	}

	grpcRes := pb.DeleteCommentResponse{
		Success: res,
	}

	return &grpcRes, nil
}
