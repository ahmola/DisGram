package internal

import (
	pb "services/pkg/proto/follow"
)

func OriginToResponse(follow *Follow) (*FollowResponse, error) {
	res := &FollowResponse{
		FollowerID: follow.FollowerID,
		FolloweeID: follow.FolloweeID,
	}

	return res, nil
}

// grpc
func grpcRequestToRequest(grpcReq *pb.FollowRequest) *FollowRequest {
	req := &FollowRequest{
		FolloweeID: uint(grpcReq.FolloweeId),
		FollowerID: uint(grpcReq.FollowerId),
	}

	return req
}

func responseToGrpcResponse(res *FollowResponse) *pb.FollowResponse {
	grpcRes := &pb.FollowResponse{
		FollowerId: uint32(res.FollowerID),
		FolloweeId: uint32(res.FolloweeID),
	}

	return grpcRes
}

func idListToGrpcResponse(idList []*uint) *pb.IdListResponse {
	grpcRes := &pb.IdListResponse{}

	for _, id := range idList {
		grpcRes.IdList = append(grpcRes.IdList, uint32(*id))
	}

	return grpcRes
}
