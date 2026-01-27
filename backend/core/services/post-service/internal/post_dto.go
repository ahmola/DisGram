package internal

// PostImage
type PostImageRequest struct {
	PostID    uint   `json:"postId"`
	FileKey   string `json:"fileKey"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
	Seq       uint   `json:"seq"`
}

type PostImageResponse struct {
	ID        uint   `json:"Id"`
	PostID    uint   `json:"postId"`
	FileKey   string `json:"fileKey"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
	Seq       uint   `json:"seq"`
}

// Like
type LikeRequest struct {
	PostID uint `json:"postId"`
	UserID uint `json:"userId"`
}

type LikeResponse struct {
	ID     uint `json:"Id"`
	PostID uint `json:"postId"`
	UserID uint `json:"userId"`
}

// Post
type PostRequest struct {
	UserID     uint               `json:"userId"`
	Caption    string             `json:"caption"`
	PostImages []PostImageRequest `json:"postImage"`
}

type PostResponse struct {
	ID            uint     `json:"Id"`
	UserID        uint     `json:"userId"`
	Caption       string   `json:"caption"`
	PostImageUrls []string `json:"postImageUrls"`
}
