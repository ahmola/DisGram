package internal

import (
	pb "services/pkg/proto/user"
)

func OriginToResponse(user User) (*UserResponse, error) {
	res := &UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Bio:       user.Bio,
		AvatarUrl: user.AvatarUrl,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return res, nil
}

func RequestToOrigin(req UserRequest) (*User, error) {
	res := &User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		Bio:          req.Bio,
		AvatarUrl:    req.AvatarUrl,
	}

	return res, nil
}

// grpc

func grpcRequestToRequest(grpcReq *pb.UserRequest) *UserRequest {
	return &UserRequest{
		Username:     grpcReq.Username,
		Email:        grpcReq.Email,
		PasswordHash: grpcReq.PasswordHash,
		Bio:          grpcReq.Bio,
		AvatarUrl:    grpcReq.AvatarUrl,
	}
}

func responseToGrpcResponse(res *UserResponse) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        uint32(res.Id),
		Username:  res.Username,
		Email:     res.Email,
		Bio:       res.Bio,
		AvatarUrl: res.AvatarUrl,
		CreatedAt: res.CreatedAt,
	}
}
