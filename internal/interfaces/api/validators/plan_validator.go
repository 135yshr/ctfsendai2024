package validators

import "time"

// GetPlansRequest プラン一覧取得リクエスト.
type GetPlansRequest struct {
	StartDate *time.Time `form:"startDate"                          time_format:"2006-01-02"`
	EndDate   *time.Time `form:"endDate"                            time_format:"2006-01-02"`
	Status    string     `binding:"omitempty,oneof=active inactive" form:"status"`
}
