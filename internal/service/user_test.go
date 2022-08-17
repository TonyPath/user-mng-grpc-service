package service

import (
	"context"
	"errors"
	"testing"
	"time"

	// 3rd party
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
)

func TestUserService_CreateUser_Success(t *testing.T) {
	guard := make(chan struct{})
	uuidMock := uuid.MustParse("d79b55a7-0ab9-4a54-b5f7-33f56f9f16f5")

	newUser := models.NewUser{
		Email: "antonis.papath@mail.com	",
		FirstName: "antonis",
		LastName:  "papath",
		Nickname:  "TonyPath",
		Country:   "GR",
		Password:  "password",
	}

	repoMock := userStorageMock{
		InsertUserFunc: func(ctx context.Context, user models.User) (uuid.UUID, error) {
			return uuidMock, nil
		},
	}

	publisherMock := eventPublisherMock{
		PublishFunc: func(ctx context.Context, topic string, key string, pbMessage protoreflect.ProtoMessage) error {
			guard <- struct{}{}
			return nil
		},
	}

	s := NewUserService(&repoMock, &publisherMock)

	uID, err := s.CreateUser(context.TODO(), newUser)

	<-guard

	require.NoError(t, err)
	require.Equal(t, uuidMock, uID)
	require.Len(t, repoMock.InsertUserCalls(), 1)
	require.Len(t, publisherMock.PublishCalls(), 1)
}

func TestUserService_CreateUser_Fail(t *testing.T) {
	mockNewUser := models.NewUser{
		Email: "antonis.papath@mail.com	",
		FirstName: "antonis",
		LastName:  "papath",
		Nickname:  "TonyPath",
		Country:   "GR",
		Password:  "password",
	}

	type deps struct {
		repo *userStorageMock
	}
	type args struct {
		newUser models.NewUser
	}

	tests := []struct {
		name            string
		deps            deps
		args            args
		checkFn         func(t *testing.T, userID uuid.UUID, err error)
		insertUserCalls int
		publishCalls    int
	}{
		{
			name: "ErrEmailTaken",
			deps: deps{
				repo: &userStorageMock{
					InsertUserFunc: func(ctx context.Context, user models.User) (uuid.UUID, error) {
						return uuid.Nil, models.ErrEmailTaken
					},
				},
			},
			args: args{
				newUser: mockNewUser,
			},
			checkFn: func(t *testing.T, userID uuid.UUID, err error) {
				require.ErrorIs(t, err, models.ErrEmailTaken)
				require.Equal(t, uuid.Nil, userID)
			},
			insertUserCalls: 1,
		},
		{
			name: "Internal error",
			deps: deps{
				repo: &userStorageMock{
					InsertUserFunc: func(ctx context.Context, user models.User) (uuid.UUID, error) {
						return uuid.Nil, errors.New("internal error")
					},
				},
			},
			args: args{
				newUser: mockNewUser,
			},
			checkFn: func(t *testing.T, userID uuid.UUID, err error) {
				require.Error(t, err)
				require.Equal(t, uuid.Nil, userID)
			},
			insertUserCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publisherMock := eventPublisherMock{}

			s := NewUserService(tt.deps.repo, &publisherMock)

			uID, err := s.CreateUser(context.TODO(), tt.args.newUser)
			tt.checkFn(t, uID, err)
			require.Len(t, tt.deps.repo.InsertUserCalls(), tt.insertUserCalls)
			require.Len(t, publisherMock.PublishCalls(), 0)
		})
	}
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	guard := make(chan struct{})
	uuidMock := uuid.MustParse("d79b55a7-0ab9-4a54-b5f7-33f56f9f16f5")

	updateUser := models.UpdateUser{
		Email: "antonis.papath@mail.com	",
		FirstName: "antonis",
		LastName:  "papath",
		Nickname:  "TonyPath",
		Country:   "GR",
		Password:  "password",
	}

	repoMock := userStorageMock{
		UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, user models.User) error {
			return nil
		},
		GetUserByIDFunc: func(ctx context.Context, userID uuid.UUID) (models.User, error) {
			return models.User{
				ID: uuidMock,
				Email: "antonis.papath@mail.com	",
				FirstName: "antonis",
				LastName:  "papath",
				Nickname:  "TonyPath",
				Country:   "GR",
				Password:  []byte(`password`),
				CreatedAt: time.Now(),
				UpdateAt:  nil,
			}, nil
		},
	}

	publisherMock := eventPublisherMock{
		PublishFunc: func(ctx context.Context, topic string, key string, pbMessage protoreflect.ProtoMessage) error {
			guard <- struct{}{}
			return nil
		},
	}

	s := NewUserService(&repoMock, &publisherMock)

	err := s.UpdateUser(context.TODO(), uuidMock, updateUser)

	<-guard

	require.NoError(t, err)
	require.Len(t, repoMock.UpdateUserCalls(), 1)
	require.Len(t, repoMock.GetUserByIDCalls(), 1)
	require.Len(t, publisherMock.PublishCalls(), 1)
}

func TestUserService_UpdateUser_Fail(t *testing.T) {
	uuidMock := uuid.MustParse("d79b55a7-0ab9-4a54-b5f7-33f56f9f16f5")
	updateUser := models.UpdateUser{
		Email: "antonis.papath@mail.com	",
		FirstName: "antonis",
		LastName:  "papath",
		Nickname:  "TonyPath",
		Country:   "GR",
		Password:  "password",
	}

	type deps struct {
		repo *userStorageMock
	}
	type args struct {
		updateUser models.UpdateUser
	}

	tests := []struct {
		name             string
		deps             deps
		args             args
		checkFn          func(t *testing.T, err error)
		updateUserCalls  int
		getUserByIDCalls int
	}{
		{
			name: "ErrUserNotFound",
			deps: deps{
				repo: &userStorageMock{
					UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, user models.User) error {
						return nil
					},
					GetUserByIDFunc: func(ctx context.Context, userID uuid.UUID) (models.User, error) {
						return models.User{}, models.ErrUserNotFound
					},
				},
			},
			args: args{
				updateUser: updateUser,
			},
			checkFn: func(t *testing.T, err error) {
				require.ErrorIs(t, err, models.ErrUserNotFound)
			},
			updateUserCalls:  0,
			getUserByIDCalls: 1,
		},
		{
			name: "Internal error",
			deps: deps{
				repo: &userStorageMock{
					UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, user models.User) error {
						return errors.New("internal error")
					},
					GetUserByIDFunc: func(ctx context.Context, userID uuid.UUID) (models.User, error) {
						return models.User{
							ID: uuidMock,
							Email: "antonis.papath@mail.com	",
							FirstName: "antonis",
							LastName:  "papath",
							Nickname:  "TonyPath",
							Country:   "GR",
							Password:  []byte(`password`),
							CreatedAt: time.Now(),
							UpdateAt:  nil,
						}, nil
					},
				},
			},
			args: args{
				updateUser: updateUser,
			},
			checkFn: func(t *testing.T, err error) {
				require.Error(t, err)
			},
			updateUserCalls:  1,
			getUserByIDCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publisherMock := eventPublisherMock{}

			s := NewUserService(tt.deps.repo, &publisherMock)

			err := s.UpdateUser(context.TODO(), uuidMock, tt.args.updateUser)
			tt.checkFn(t, err)
			require.Len(t, tt.deps.repo.UpdateUserCalls(), tt.updateUserCalls)
			require.Len(t, tt.deps.repo.GetUserByIDCalls(), tt.getUserByIDCalls)
			require.Len(t, publisherMock.PublishCalls(), 0)
		})
	}
}
