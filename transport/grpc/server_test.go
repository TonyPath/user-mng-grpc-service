package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	// 3rd party
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	"github.com/TonyPath/user-mng-grpc-service/proto/services/user"
)

func TestGRPC_CreateUser(t *testing.T) {
	type fields struct {
		svc userService
	}
	type args struct {
		req *user.CreateUserRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		checkFn func(t *testing.T, resp *user.CreateUserResponse, err error)
	}{
		{
			name: "happy path",
			fields: fields{
				svc: &UserServiceMock{
					CreateUserFunc: func(ctx context.Context, nu models.NewUser) (uuid.UUID, error) {
						return uuid.MustParse("1c8f21c1-c8d0-401c-89b5-3f577c54679e"), nil
					},
				},
			},
			args: args{
				req: &user.CreateUserRequest{
					Email:     "antonis.test@mail.com",
					FirstName: "antonis",
					LastName:  "papath",
					Nickname:  "TonyPath",
					Password:  "password",
					Country:   "GR",
				},
			},
			checkFn: func(t *testing.T, resp *user.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, &user.CreateUserResponse{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
				}, resp)
			},
		},
		{
			name: "email has been taken",
			fields: fields{
				svc: &UserServiceMock{
					CreateUserFunc: func(ctx context.Context, nu models.NewUser) (uuid.UUID, error) {
						return uuid.Nil, models.ErrEmailTaken
					},
				},
			},
			args: args{
				req: &user.CreateUserRequest{
					Email:     "antonis.test@mail.com",
					FirstName: "antonis",
					LastName:  "papath",
					Nickname:  "TonyPath",
					Password:  "password",
					Country:   "GR",
				},
			},
			checkFn: func(t *testing.T, resp *user.CreateUserResponse, err error) {
				require.ErrorIs(t, err, errEmailTaken)
			},
		},
		{
			name: "internal server error",
			fields: fields{
				svc: &UserServiceMock{
					CreateUserFunc: func(ctx context.Context, nu models.NewUser) (uuid.UUID, error) {
						return uuid.Nil, errors.New("internal error")
					},
				},
			},
			args: args{
				req: &user.CreateUserRequest{
					Email:     "antonis.test@mail.com",
					FirstName: "antonis",
					LastName:  "papath",
					Nickname:  "TonyPath",
					Password:  "password",
					Country:   "GR",
				},
			},
			checkFn: func(t *testing.T, resp *user.CreateUserResponse, err error) {
				require.ErrorIs(t, err, errInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				svc:    tt.fields.svc,
				logger: zap.NewNop().Sugar(),
			}
			got, err := g.CreateUser(context.Background(), tt.args.req)
			tt.checkFn(t, got, err)
		})
	}
}

func TestGRPC_UpdateUser(t *testing.T) {
	type fields struct {
		svc userService
	}
	type args struct {
		req *user.UpdateUserRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		checkFn func(t *testing.T, resp *user.UpdateUserResponse, err error)
	}{
		{
			name: "happy path",
			fields: fields{
				svc: &UserServiceMock{
					UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error {
						return nil
					},
				},
			},
			args: args{
				req: &user.UpdateUserRequest{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
					Fields: &user.UpdateUserRequest_Fields{
						Nickname: "user_nickname",
					},
				},
			},
			checkFn: func(t *testing.T, resp *user.UpdateUserResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, &user.UpdateUserResponse{
					Success: true,
				}, resp)
			},
		},
		{
			name: "invalid uuid",
			fields: fields{
				svc: &UserServiceMock{},
			},
			args: args{
				req: &user.UpdateUserRequest{
					UserId: "invalid uuid",
					Fields: &user.UpdateUserRequest_Fields{
						Nickname: "user_nickname",
					},
				},
			},
			checkFn: func(t *testing.T, resp *user.UpdateUserResponse, err error) {
				require.ErrorIs(t, err, errInvalidUserID)
			},
		},
		{
			name: "user not found",
			fields: fields{
				svc: &UserServiceMock{
					UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error {
						return models.ErrUserNotFound
					},
				},
			},
			args: args{
				req: &user.UpdateUserRequest{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
					Fields: &user.UpdateUserRequest_Fields{
						Nickname: "user_nickname",
					},
				},
			},
			checkFn: func(t *testing.T, resp *user.UpdateUserResponse, err error) {
				require.ErrorIs(t, err, errUserNotFound)
			},
		},
		{
			name: "internal server error",
			fields: fields{
				svc: &UserServiceMock{
					UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error {
						return errors.New("internal server error")
					},
				},
			},
			args: args{
				req: &user.UpdateUserRequest{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
					Fields: &user.UpdateUserRequest_Fields{
						Nickname: "user_nickname",
					},
				},
			},
			checkFn: func(t *testing.T, resp *user.UpdateUserResponse, err error) {
				require.ErrorIs(t, err, errInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				svc:    tt.fields.svc,
				logger: zap.NewNop().Sugar(),
			}
			got, err := g.UpdateUser(context.Background(), tt.args.req)
			tt.checkFn(t, got, err)
		})
	}
}

func TestGRPC_DeleteUser(t *testing.T) {
	type fields struct {
		svc userService
	}
	type args struct {
		req *user.DeleteUserRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		checkFn func(t *testing.T, resp *user.DeleteUserResponse, err error)
	}{
		{
			name: "happy path",
			fields: fields{
				svc: &UserServiceMock{
					DeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
						return nil
					},
				},
			},
			args: args{
				req: &user.DeleteUserRequest{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
				},
			},
			checkFn: func(t *testing.T, resp *user.DeleteUserResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, &user.DeleteUserResponse{
					Success: true,
				}, resp)
			},
		},
		{
			name: "invalid uuid",
			fields: fields{
				svc: &UserServiceMock{},
			},
			args: args{
				req: &user.DeleteUserRequest{
					UserId: "invalid uuid",
				},
			},
			checkFn: func(t *testing.T, resp *user.DeleteUserResponse, err error) {
				require.ErrorIs(t, err, errInvalidUserID)
			},
		},
		{
			name: "user not found",
			fields: fields{
				svc: &UserServiceMock{
					DeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
						return models.ErrUserNotFound
					},
				},
			},
			args: args{
				req: &user.DeleteUserRequest{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
				},
			},
			checkFn: func(t *testing.T, resp *user.DeleteUserResponse, err error) {
				require.ErrorIs(t, err, errUserNotFound)
			},
		},
		{
			name: "internal server error",
			fields: fields{
				svc: &UserServiceMock{
					DeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
						return errors.New("internal server error")
					},
				},
			},
			args: args{
				req: &user.DeleteUserRequest{
					UserId: "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
				},
			},
			checkFn: func(t *testing.T, resp *user.DeleteUserResponse, err error) {
				require.ErrorIs(t, err, errInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				svc:    tt.fields.svc,
				logger: zap.NewNop().Sugar(),
			}
			got, err := g.DeleteUser(context.Background(), tt.args.req)
			tt.checkFn(t, got, err)
		})
	}
}

func TestGRPC_QueryUsers(t *testing.T) {

	now := time.Now()

	type fields struct {
		svc userService
	}
	type args struct {
		req *user.QueryUsersRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		checkFn func(t *testing.T, resp *user.QueryUsersResponse, err error)
	}{
		{
			name: "happy path",
			fields: fields{
				svc: &UserServiceMock{
					GetUsersFunc: func(ctx context.Context, qu models.GetUsersOptions) ([]models.User, error) {
						require.Equal(t, uint64(1), qu.PageNumber)
						require.Equal(t, uint64(defaultPageSize), qu.PageSize)

						return []models.User{
							{
								ID:        uuid.MustParse("1c8f21c1-c8d0-401c-89b5-3f577c54679e"),
								Email:     "antonis@mail.com",
								FirstName: "antonis",
								LastName:  "papath",
								Nickname:  "TonyPath",
								Country:   "GR",
								Password:  []byte(`super_secret`),
								CreatedAt: now,
								UpdateAt:  nil,
							},
						}, nil
					},
				},
			},
			args: args{
				req: &user.QueryUsersRequest{
					PageNumber: 1,
					PageSize:   10,
				},
			},
			checkFn: func(t *testing.T, resp *user.QueryUsersResponse, err error) {
				require.NoError(t, err)
				if !reflect.DeepEqual(&user.QueryUsersResponse{
					Users: []*user.UserInfo{
						{
							Id:        "1c8f21c1-c8d0-401c-89b5-3f577c54679e",
							Email:     "antonis@mail.com",
							FirstName: "antonis",
							LastName:  "papath",
							Nickname:  "TonyPath",
							Country:   "GR",
							Password:  "super_secret",
							CreatedAt: timestamppb.New(now),
							UpdateAt:  nil,
						},
					},
				}, resp) {
					t.Errorf("GRPC.QueryUsers()")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				svc:    tt.fields.svc,
				logger: zap.NewNop().Sugar(),
			}
			got, err := g.QueryUsers(context.Background(), tt.args.req)
			tt.checkFn(t, got, err)
		})
	}
}
