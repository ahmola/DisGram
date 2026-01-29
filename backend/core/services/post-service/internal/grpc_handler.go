package internal

import (
	"context"
	pb "services/pkg/proto/post"
)

// gRPC Server Interface
type PostGrpcHandler struct {
	pb.UnimplementedPostServiceServer
	Svc *PostService // buisiness logic
}

// Implementaions
func (h *PostGrpcHandler) CreatePost(ctx context.Context, grpcReq *pb.PostRequest) (*pb.PostResponse, error) {
	req := grpcPostRequestToPostRequest(*grpcReq)
	res, err := h.Svc.CreatePost(*req)
	if err != nil {
		return nil, err
	}

	grpcRes := postResponseToGrpcPostResponse(res)
	return grpcRes, nil
}

func (h *PostGrpcHandler) GetPostById(ctx context.Context, req *pb.RequestById) (*pb.PostResponse, error) {
	res, err := h.Svc.GetPostById(req.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := postResponseToGrpcPostResponse(res)

	return grpcRes, nil
}

func (h *PostGrpcHandler) UpdatePost(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	grpcReq := grpcPostRequestToPostRequest(*req)
	res, err := h.Svc.UpdatePost(*grpcReq)
	if err != nil {
		return nil, err
	}

	grpcRes := postResponseToGrpcPostResponse(res)
	return grpcRes, nil
}

func (h *PostGrpcHandler) DeletePost(ctx context.Context, req *pb.RequestById) (*pb.ResponseByBoolean, error) {
	res, err := h.Svc.DeletePostById(req.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := pb.ResponseByBoolean{
		Success: res,
	}

	return &grpcRes, nil
}

// like service

func (h *PostGrpcHandler) CreateLike(ctx context.Context, req *pb.LikeRequest) (*pb.ResponseByBoolean, error) {
	grpcReq := grpcLikeRequestToLikeRequest(*req)
	res, err := h.Svc.CreateLike(*grpcReq)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseByBoolean{
		Success: res,
	}, nil
}

func (h *PostGrpcHandler) GetAllLikesByPostID(ctx context.Context, req *pb.RequestById) (*pb.LikeListResponse, error) {
	res, err := h.Svc.GetAllLikesByPostID(req.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := likeListResponseToGrpcLikeResponse(res)

	return grpcRes, nil
}
