package internal

func OriginToResponse(follow *Follow) (*FollowResponse, error) {
	res := &FollowResponse{
		FollowerID: follow.FollowerID,
		FolloweeID: follow.FolloweeID,
	}

	return res, nil
}
