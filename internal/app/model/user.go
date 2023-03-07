package model

// User entity.
type User struct {
	ID   int    `json:"id,omitempty"`
	UUID string `json:"uuid"`
}

// NewUser User entity constructor.
func NewUser(uuid string) *User {
	return &User{UUID: uuid}
}
