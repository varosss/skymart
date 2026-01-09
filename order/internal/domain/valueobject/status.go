package valueobject

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusCancelled Status = "cancelled"
	StatusCompleted Status = "completed"
)
