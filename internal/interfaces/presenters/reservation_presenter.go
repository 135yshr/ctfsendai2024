package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

type ReservationPresenter interface {
	PresentReservations(reservations []*dto.ReservationResponse) ReservationsResponse
	PresentError(err error) response.ErrorResponse
}

type JSONReservationPresenter struct{}

func NewJSONReservationPresenter() ReservationPresenter {
	return &JSONReservationPresenter{}
}

// ReservationsResponse 予約一覧レスポンス
// @Description 予約一覧のレスポンス.
type ReservationsResponse struct {
	// ステータス
	// @Example "success"
	Status string `example:"success" json:"status"`

	// 予約データ
	Data []*dto.ReservationResponse `json:"data"`
}

func (p *JSONReservationPresenter) PresentReservations(reservations []*dto.ReservationResponse) ReservationsResponse {
	return ReservationsResponse{
		Status: "success",
		Data:   reservations,
	}
}

func (p *JSONReservationPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
