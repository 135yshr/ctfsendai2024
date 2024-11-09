package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
)

type ReservationPresenter interface {
	PresentReservations([]*dto.ReservationResponse) interface{}
	PresentError(error) interface{}
}

type JSONReservationPresenter struct{}

func NewJSONReservationPresenter() ReservationPresenter {
	return &JSONReservationPresenter{}
}

// ReservationsResponse 予約一覧レスポンス
// @Description 予約一覧のレスポンス
type ReservationsResponse struct {
	// ステータス
	// @Example "success"
	Status string `json:"status" example:"success"`

	// 予約データ
	Data []*dto.ReservationResponse `json:"data"`
}

// ErrorResponse エラーレスポンス
// @Description エラー情報のレスポンス
type ErrorResponse struct {
	// ステータス
	// @Example "error"
	Status string `json:"status" example:"error"`

	// エラーメッセージ
	// @Example "ユーザーIDは必須項目です"
	Error string `json:"error" example:"ユーザーIDは必須項目です"`
}

func (p *JSONReservationPresenter) PresentReservations(reservations []*dto.ReservationResponse) interface{} {
	return ReservationsResponse{
		Status: "success",
		Data:   reservations,
	}
}

func (p *JSONReservationPresenter) PresentError(err error) interface{} {
	return ErrorResponse{
		Status: "error",
		Error:  err.Error(),
	}
}
