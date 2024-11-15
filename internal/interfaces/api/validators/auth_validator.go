package validators

// LoginRequest ログインリクエスト.
// @Description ユーザーIDとパスワードによるログインリクエスト.
type LoginRequest struct {
	UserID   string `binding:"required,min=3,max=50"  json:"user_id"`
	Password string `binding:"required,min=3,max=100" json:"password"`
}

// SecretLoginRequest 秘密の質問によるログインリクエスト.
// @Description 秘密の質問の回答によるログインリクエスト.
type SecretLoginRequest struct {
	UserID       string `binding:"required,min=3,max=50" json:"user_id"`
	SecretAnswer string `binding:"required,min=1"        json:"secret_answer"`
}

// SecretQuestionRequest 秘密の質問取得リクエスト.
// @Description 秘密の質問を取得するためのリクエスト.
type SecretQuestionRequest struct {
	UserID string `binding:"required,min=3,max=50" form:"user_id"`
}
