package dto

type LoginRequest struct {
	UserID   string `binding:"required" json:"user_id"`
	Password string `binding:"required" json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type SecretLoginRequest struct {
	UserID       string `binding:"required" json:"user_id"`
	SecretAnswer string `binding:"required" json:"secret_answer"`
}

type SecretQuestionRequest struct {
	UserID string `binding:"required" query:"user_id"`
}

type SecretQuestionResponse struct {
	SecretQuestion string `json:"secret_question"`
}
