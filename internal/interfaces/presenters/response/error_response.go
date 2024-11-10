package response

// ErrorResponse エラーレスポンス
// @Description エラー情報のレスポンス
type ErrorResponse struct {
	// ステータス
	// @Example "error"
	Status string `json:"status" example:"error"`

	// エラーメッセージ
	// @Example "プランが見つかりません"
	Error string `json:"error" example:"プランが見つかりません"`
}

// NewErrorResponse は新しいErrorResponseインスタンスを生成します
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error:  err.Error(),
	}
}
