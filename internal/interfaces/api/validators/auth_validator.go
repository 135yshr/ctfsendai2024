package validators

// LoginRequest ログインリクエスト.
type LoginRequest struct {
	UserID   string `binding:"required,min=3,max=50"  json:"user_id"`
	Password string `binding:"required,min=3,max=100" json:"password"`
}

// SecretLoginRequest 秘密の質問によるログインリクエスト.
type SecretLoginRequest struct {
	UserID       string `binding:"required,min=3,max=50" json:"user_id"`
	SecretAnswer string `binding:"required,min=1"        json:"secret_answer"`
}

// SecretQuestionRequest 秘密の質問取得リクエスト.
type SecretQuestionRequest struct {
	UserID string `binding:"required,min=3,max=50" form:"user_id"`
}
