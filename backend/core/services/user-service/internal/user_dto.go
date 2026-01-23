package internal

type UserRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Bio          string `json:"bio"`
	AvatarUrl    string `json:"avatarUrl"`
}

type UserResponse struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatarUrl"`
}
