package valueobject

type Status string

var (
	StatusDraft       Status = "draft"
	StatusPublished   Status = "published"
	StatusUnpublished Status = "unpublished"
	StatusArchived    Status = "archived"
)
