package response

// ErrorResponse エラーレスポンス
// @Description エラー情報のレスポンス.
type ErrorResponse struct {
	// ステータス
	// @Example "error"
	Status string `example:"error" json:"status"`

	// エラーメッセージ
	// @Example "プランが見つかりません"
	Error string `example:"プランが見つかりません" json:"error"`
}

// NewErrorResponse は新しいErrorResponseインスタンスを生成します.
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error:  err.Error(),
	}
}
