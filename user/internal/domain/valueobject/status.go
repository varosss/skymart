package valueobject

type Status string

const (
	StatusActive  Status = "active"
	StatusBlocked Status = "blocked"
	StatusDeleted Status = "deleted"
)

func (status Status) String() string {
	return string(status)
}
