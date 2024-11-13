package models

import (
	"time"
)

type Plan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Duration    int    `json:"duration"`
}

type PlanSearchParams struct {
	StartDate *time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate   *time.Time `form:"end_date"   time_format:"2006-01-02"`
	Status    *string    `form:"status"`
}
