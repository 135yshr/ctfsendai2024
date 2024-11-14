package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

type ReservationPresenter interface {
	PresentReservations(reservations []*dto.ReservationResponse) ReservationsResponse
	PresentError(err error) response.ErrorResponse
	PresentReservation(reservation *dto.ReservationResponse) ReservationResponse
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

// ReservationResponse 予約レスポンス
// @Description 予約のレスポンス.
type ReservationResponse struct {
	// ステータス
	// @Example "success"
	Status string `example:"success" json:"status"`

	// 予約データ
	Data *dto.ReservationResponse `json:"data"`
}

func (p *JSONReservationPresenter) PresentReservation(reservation *dto.ReservationResponse) ReservationResponse {
	return ReservationResponse{
		Status: "success",
		Data:   reservation,
	}
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
