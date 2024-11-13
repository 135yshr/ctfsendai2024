package dto

type LoginRequest struct {
	UserID   string `binding:"required" json:"user_id"`
	Password string `binding:"required" json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}
