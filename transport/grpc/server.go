package grpc

import (
	"context"
	// 3rd party
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	pb "github.com/TonyPath/user-mng-grpc-service/proto/services/user"
)

const defaultPageSize = 10

//go:generate moq -out user_service_mock_test.go . UserService
type userService interface {
	CreateUser(ctx context.Context, nu models.NewUser) (uuid.UUID, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUsers(ctx context.Context, qu models.GetUsersOptions) ([]models.User, error)
}

type GRPC struct {
	pb.UnimplementedUserServer

	logger *zap.SugaredLogger
	svc    userService
}

func New(logger *zap.SugaredLogger, svc userService) *GRPC {
	return &GRPC{
		logger: logger,
		svc:    svc,
	}
}

func (g *GRPC) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := models.NewUser{
		Email:     req.GetEmail(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Nickname:  req.GetNickname(),
		Password:  req.GetPassword(),
		Country:   req.GetCountry(),
	}

	userID, err := g.svc.CreateUser(ctx, newUser)
	if err != nil {
		return nil, g.mapError(err)
	}

	return &pb.CreateUserResponse{
		UserId: userID.String(),
	}, nil
}

func (g *GRPC) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, errInvalidUserID
	}

	updateUser := models.UpdateUser{
		Email:     req.GetFields().GetEmail(),
		FirstName: req.GetFields().GetFirstName(),
		LastName:  req.GetFields().GetLastName(),
		Nickname:  req.GetFields().GetNickname(),
		Country:   req.GetFields().GetCountry(),
		Password:  req.GetFields().GetPassword(),
	}

	if err = g.svc.UpdateUser(ctx, userID, updateUser); err != nil {
		return nil, g.mapError(err)
	}

	return &pb.UpdateUserResponse{
		Success: true,
	}, nil
}

func (g *GRPC) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, errInvalidUserID
	}

	if err := g.svc.DeleteUser(ctx, userID); err != nil {
		return nil, g.mapError(err)
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (g *GRPC) QueryUsers(ctx context.Context, req *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {
	quOpts := mapQueryOptions(req)
	users, err := g.svc.GetUsers(ctx, quOpts)
	if err != nil {
		return nil, g.mapError(err)
	}

	return mapUsersInfo(users)
}

func mapUsersInfo(users []models.User) (*pb.QueryUsersResponse, error) {
	items := make([]*pb.UserInfo, len(users))

	for i, u := range users {
		items[i] = &pb.UserInfo{
			Id:        u.ID.String(),
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			Country:   u.Country,
			Password:  string(u.Password),
			CreatedAt: timestamppb.New(u.CreatedAt),
			//UpdateAt:  timestamppb.New(u.UpdateAt),
		}
	}

	return &pb.QueryUsersResponse{
		Users: items,
	}, nil
}

func mapQueryOptions(req *pb.QueryUsersRequest) models.GetUsersOptions {
	pageNumber := req.GetPageNumber()
	if pageNumber == 0 {
		pageNumber = 1
	}

	pageSize := req.GetPageSize()
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	quOpts := models.GetUsersOptions{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}

	if req.GetFilter() != nil {
		if cnt := req.GetFilter().GetCountry(); cnt != "" {
			quOpts.Filter.Country = cnt
		}

		if email := req.GetFilter().GetEmail(); email != "" {
			quOpts.Filter.Email = email
		}

		if nickname := req.GetFilter().GetNickname(); nickname != "" {
			quOpts.Filter.Nickname = nickname
		}
	}

	return quOpts
}
