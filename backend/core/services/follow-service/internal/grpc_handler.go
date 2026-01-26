package internal

import (
	"context"
	pb "services/pkg/proto/follow"
)

// gRPC Server Interface
type FollowGrpcHandler struct {
	pb.UnimplementedFollowServiceServer
	Svc *FollowService // Business Logic
}

// Implementations
func (hdl *FollowGrpcHandler) CreateFollow(ctx context.Context, grpcReq *pb.FollowRequest) (*pb.FollowResponse, error) {
	req := grpcRequestToRequest(grpcReq)
	res, err := hdl.Svc.CreateFollow(*req)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)
	return grpcRes, nil
}

func (hdl *FollowGrpcHandler) GetFolloweesByFollowerID(ctx context.Context, followerID *pb.FollowRequestById) (*pb.IdListResponse, error) {
	res, err := hdl.Svc.GetFolloweesByFollowerID(followerID.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := idListToGrpcResponse(res)

	return grpcRes, nil
}

func (hdl *FollowGrpcHandler) GetFollowersByFolloweeID(ctx context.Context, followeeID *pb.FollowRequestById) (*pb.IdListResponse, error) {
	res, err := hdl.Svc.GetFollowersByFolloweeID(followeeID.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := idListToGrpcResponse(res)

	return grpcRes, nil
}

func (hdl *FollowGrpcHandler) GetFollow(ctx context.Context, followId *pb.FollowRequestById) (*pb.FollowResponse, error) {
	res, err := hdl.Svc.GetFollow(followId.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)

	return grpcRes, nil
}

func (hdl *FollowGrpcHandler) DeleteFollow(ctx context.Context, followId *pb.FollowRequestById) (*pb.DeleteFollowResponse, error) {
	res, err := hdl.Svc.DeleteFollow(followId.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := &pb.DeleteFollowResponse{Success: res}

	return grpcRes, nil
}
