package valueobject

const (
	Admin = iota
	Regular
)

type Role struct {
	Type int
}

func NewRole() Role {
	return Role{Type: Regular}
}

func Int2Role(roleInt int) Role {
	if roleInt == Admin {
		return Role{Type: Admin}
	}
	return Role{Type: Regular}
}

func (r Role) IsAdmin() bool {
	return r.Type == Admin
}
