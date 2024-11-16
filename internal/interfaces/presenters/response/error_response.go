package response

// ErrorResponse エラーレスポンス
// @Description エラー情報のレスポンス.
type ErrorResponse struct {
	// ステータス
	Status string `example:"error" json:"status"`

	// エラーメッセージ
	Error string `example:"プランが見つかりません" json:"error"`
}

// NewErrorResponse は新しいErrorResponseインスタンスを生成します.
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error:  err.Error(),
	}
}
