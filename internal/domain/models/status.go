package models

type Status string

const (
	StatusReserved  Status = "reserved"
	StatusCanceled  Status = "canceled"
	StatusConfirmed Status = "confirmed"
	StatusPending   Status = "pending"
)
