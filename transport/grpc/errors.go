package grpc

import (
	"errors"

	// 3rd party
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
)

var (
	errInvalidUserID = status.Errorf(codes.InvalidArgument, "invalid user id")
	errUserNotFound  = status.Errorf(codes.NotFound, "user not found")
	errEmailTaken    = status.Errorf(codes.AlreadyExists, "email is already used")
	errInternal      = status.Errorf(codes.Internal, "internal server error")
)

func (g *GRPC) mapError(err error) error {
	switch {
	case errors.Is(err, models.ErrUserNotFound):
		return errUserNotFound
	case errors.Is(err, models.ErrEmailTaken):
		return errEmailTaken
	default:
		g.logger.Error(err)
		return errInternal
	}
}
