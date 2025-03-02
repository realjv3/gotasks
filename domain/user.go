package domain

import "time"

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	Get(id int) (*User, error)
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetUser(id int) (*User, error)
}
