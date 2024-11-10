package models

import "time"

type Reservation struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	User      User      `json:"user"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
	Plan      Plan      `json:"plan"`
}
