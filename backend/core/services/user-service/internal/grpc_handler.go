package internal

import (
	"context"
	pb "services/pkg/proto/user"
)

type UserGrpcHandler struct {
	pb.UnimplementedUserServiceServer
	Svc *UserService
}

func (h *UserGrpcHandler) GetUserByID(ctx context.Context, grpcReq *pb.UserRequestById) (*pb.UserResponse, error) {
	res, err := h.Svc.GetUserByID(grpcReq.Id)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)
	return grpcRes, nil
}

func (h *UserGrpcHandler) CreateUser(ctx context.Context, grpcReq *pb.UserRequest) (*pb.UserResponse, error) {
	req := grpcRequestToRequest(grpcReq)
	res, err := h.Svc.CreateUser(*req)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)
	return grpcRes, nil
}

func (h *UserGrpcHandler) UpdateUser(ctx context.Context, grpcReq *pb.UserRequest) (*pb.UserResponse, error) {
	req := grpcRequestToRequest(grpcReq)
	res, err := h.Svc.UpdateUser(*req)
	if err != nil {
		return nil, err
	}

	grpcRes := responseToGrpcResponse(res)
	return grpcRes, nil
}

func (h *UserGrpcHandler) DeleteUser(ctx context.Context, grpcReq *pb.UserRequestById) (*pb.DeleteUserResponse, error) {
	grpcRes, err := h.Svc.DeleteUser(grpcReq.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Success: grpcRes}, nil
}
