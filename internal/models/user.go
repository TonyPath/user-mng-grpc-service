package models

import (
	"time"

	// 3rd party
	"github.com/google/uuid"
)

// User represents a user.
type User struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
	Nickname  string
	Country   string
	Password  []byte
	CreatedAt time.Time
	UpdateAt  *time.Time
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Email     string
	FirstName string
	LastName  string
	Nickname  string
	Country   string
	Password  string
}

// UpdateUser defines the information may be provided to modify an existing user.
type UpdateUser struct {
	Email     string
	FirstName string
	LastName  string
	Nickname  string
	Country   string
	Password  string
}

// GetUsersOptions defines the information may be provided to fetch users.
type GetUsersOptions struct {
	PageNumber uint64
	PageSize   uint64
	Filter     struct {
		Country  string
		Email    string
		Nickname string
	}
}
