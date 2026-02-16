package valueobject

type RoleCode string

const (
	UserRole      RoleCode = "user"
	AdminRole     RoleCode = "admin"
	ModeratorRole RoleCode = "moderator"
	SupportRole   RoleCode = "support"
)

func (code RoleCode) String() string {
	return string(code)
}
