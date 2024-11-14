package models

import "time"

type Reservation struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	User      *User     `json:"user"`
	PlanID    string    `json:"plan_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    Status    `json:"status"`
	Plan      *Plan     `json:"plan"`
}
