package dto

type LoginRequest struct {
	UserID   string `binding:"required" json:"user_id"`
	Password string `binding:"required" json:"password"`
}

// LoginResponse ログインレスポンス
// @Description ログイン処理のレスポンス.
type LoginResponse struct {
	// アクセストークン
	AccessToken string `json:"access_token"`

	// トークンの有効期限
	ExpiresAt int64 `json:"expires_at"`
}

type SecretLoginRequest struct {
	UserID       string `binding:"required" json:"user_id"`
	SecretAnswer string `binding:"required" json:"secret_answer"`
}

type SecretQuestionRequest struct {
	UserID string `binding:"required" form:"user_id"`
}

type SecretQuestionResponse struct {
	SecretQuestion string `json:"secret_question"`
}
