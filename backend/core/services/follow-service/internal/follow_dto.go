package internal

type FollowRequest struct {
	FolloweeID uint `json:"followeeID"`
	FollowerID uint `json:"followerID"`
}

type FollowResponse struct {
	ID         uint   `json:"id"`
	FolloweeID uint   `json:"followeeID"`
	FollowerID uint   `json:"followerID"`
	CreatedAt  string `json:"createdAt"`
}
