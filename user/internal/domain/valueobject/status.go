package valueobject

type Status string

const (
	StatusActive  Status = "active"
	StatusBlocked Status = "blocked"
	StatusDeleted Status = "deleted"
)
