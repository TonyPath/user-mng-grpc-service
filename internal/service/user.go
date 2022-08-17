package service

import (
	"context"
	"fmt"
	"time"

	// 3rd party
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	pbevents "github.com/TonyPath/user-mng-grpc-service/proto/events/user"
)

//go:generate moq -out user_storage_mock_test.go . userStorage
type userStorage interface {
	InsertUser(ctx context.Context, user models.User) (uuid.UUID, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUsersByFilter(ctx context.Context, opts models.GetUsersOptions) ([]models.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error)
	ExistsByID(ctx context.Context, userID uuid.UUID) (bool, error)
}

//go:generate moq -out event_publisher_mock_test.go . eventPublisher
type eventPublisher interface {
	Publish(ctx context.Context, topic string, key string, pbMessage proto.Message) error
}

type UserService struct {
	repo           userStorage
	eventPublisher eventPublisher
}

func NewUserService(repo userStorage, publisher eventPublisher) *UserService {
	return &UserService{
		repo:           repo,
		eventPublisher: publisher,
	}
}

func (uSvc *UserService) CreateUser(ctx context.Context, nu models.NewUser) (uuid.UUID, error) {
	now := time.Now().UTC()

	hash, err := bcryptPassword(nu.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("generating password hash: %w", err)
	}

	user := models.User{
		ID:        uuid.New(),
		Email:     nu.Email,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		Nickname:  nu.Nickname,
		Country:   nu.Country,
		Password:  hash,
		CreatedAt: now,
		UpdateAt:  nil,
	}

	userID, err := uSvc.repo.InsertUser(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}

	go func() {
		ctx := context.Background()
		evt := pbevents.UserCreated{
			UserId:    user.ID.String(),
			CreatedAt: timestamppb.New(now),
		}
		_ = uSvc.eventPublisher.Publish(ctx, "UserCreated", userID.String(), &evt)
	}()

	return userID, nil
}

func (uSvc *UserService) UpdateUser(ctx context.Context, userID uuid.UUID, updateUser models.UpdateUser) error {
	user, err := uSvc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}

	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}

	if updateUser.Nickname != "" {
		user.Nickname = updateUser.Nickname
	}

	if updateUser.Country != "" {
		user.Country = updateUser.Country
	}

	if len(updateUser.Password) != 0 {
		hash, err := bcryptPassword(updateUser.Password)
		if err != nil {
			return err
		}
		user.Password = hash
	}

	now := time.Now().UTC()
	user.UpdateAt = &now

	if err := uSvc.repo.UpdateUser(ctx, userID, user); err != nil {
		return err
	}

	go func() {
		ctx := context.Background()
		evt := pbevents.UserUpdated{
			UserId:    user.ID.String(),
			UpdatedAt: timestamppb.New(now),
		}
		_ = uSvc.eventPublisher.Publish(ctx, "UserUpdated", userID.String(), &evt)
	}()

	return nil
}

func (uSvc *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	exists, err := uSvc.repo.ExistsByID(ctx, userID)
	if err != nil {
		return err
	}

	if !exists {
		return models.ErrUserNotFound
	}

	if err := uSvc.repo.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

func (uSvc *UserService) GetUsers(ctx context.Context, qu models.GetUsersOptions) ([]models.User, error) {
	return uSvc.repo.GetUsersByFilter(ctx, qu)
}

func bcryptPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generating password hash: %w", err)
	}
	return hash, nil
}
