package models

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("ErrUserNotFound")
	ErrEmailTaken   = errors.New("ErrEmailTaken")
)
