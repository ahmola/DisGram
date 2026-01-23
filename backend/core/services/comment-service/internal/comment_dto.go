package internal

type CommentRequest struct {
	UserID  uint   `json:"userID"`
	PostID  uint   `json:"postID"`
	Content string `json:"content"`
}

type CommentResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userID"`
	PostID    uint   `json:"postID"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}
