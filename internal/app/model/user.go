package model

type User struct {
	ID   int    `json:"id,omitempty"`
	UUID string `json:"uuid"`
}

func NewUser(uuid string) *User {
	return &User{UUID: uuid}
}
