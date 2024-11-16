package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

// ReservationPresenter インターフェース
//
//	@Description	予約情報のプレゼンター
//	@interface.
type ReservationPresenter interface {
	// PresentReservations 予約情報一覧をレスポンス形式に変換する
	//	@Summary		予約情報一覧を変換
	//	@Description	予約情報一覧を成功レスポンスの形式に変換する
	//	@Return			ReservationsResponse 変換された予約一覧レスポンス
	PresentReservations(reservations []*dto.ReservationResponse) ReservationsResponse

	// PresentReservation 単一の予約情報を変換
	//	@Summary		予約情報を変換
	//	@Description	単一の予約情報をレスポンス形式に変換する
	//	@Return			ReservationResponse 変換された予約レスポンス
	PresentReservation(reservation *dto.ReservationResponse) ReservationResponse

	// PresentError エラー情報を変換
	//	@Summary		エラー情報を変換
	//	@Description	エラー情報をエラーレスポンスの形式に変換する
	//	@Return			response.ErrorResponse エラーレスポンス
	PresentError(err error) response.ErrorResponse
}

type reservationPresenter struct{}

func NewReservationPresenter() ReservationPresenter {
	return &reservationPresenter{}
}

// ReservationResponse 予約レスポンス
//
//	@Description	予約のレスポンス
//	@Object			ReservationResponse.
type ReservationResponse struct {
	// Status レスポンスのステータス
	Status string `example:"success" json:"status"`

	// Data 予約データ
	Data *dto.ReservationResponse `json:"data"`
}

func (p *reservationPresenter) PresentReservation(reservation *dto.ReservationResponse) ReservationResponse {
	return ReservationResponse{
		Status: "success",
		Data:   reservation,
	}
}

// ReservationsResponse 予約一覧レスポンス
//
//	@Description	予約一覧のレスポンス
//	@Object			ReservationsResponse.
type ReservationsResponse struct {
	// Status レスポンスのステータス
	Status string `example:"success" json:"status"`

	// Data 予約データ一覧
	Data []*dto.ReservationResponse `json:"data"`
}

func (p *reservationPresenter) PresentReservations(reservations []*dto.ReservationResponse) ReservationsResponse {
	return ReservationsResponse{
		Status: "success",
		Data:   reservations,
	}
}

func (p *reservationPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
