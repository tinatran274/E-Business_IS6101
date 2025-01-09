package models

var (
	ActiveStatus   Status = "active"
	InactiveStatus Status = "inactive"
	DeletedStatus  Status = "deleted"
)

type Status string
