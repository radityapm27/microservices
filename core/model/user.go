package model

// User ...
type User struct {
	UserID       string
	UserPassword string
	Permissions  []*UserPermission
}

// UserPermission ...
type UserPermission struct {
	Permission string
	IsWrite    bool
	IsRead     bool
}
