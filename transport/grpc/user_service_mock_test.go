// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package grpc

import (
	"context"
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	"github.com/google/uuid"
	"sync"
)

// Ensure, that UserServiceMock does implement userService.
// If this is not the case, regenerate this file with moq.
var _ userService = &UserServiceMock{}

// UserServiceMock is a mock implementation of userService.
//
// 	func TestSomethingThatUsesUserService(t *testing.T) {
//
// 		// make and configure a mocked userService
// 		mockedUserService := &UserServiceMock{
// 			CreateUserFunc: func(ctx context.Context, nu models.NewUser) (uuid.UUID, error) {
// 				panic("mock out the CreateUser method")
// 			},
// 			DeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
// 				panic("mock out the DeleteUser method")
// 			},
// 			GetUsersFunc: func(ctx context.Context, qu models.GetUsersOptions) ([]models.User, error) {
// 				panic("mock out the GetUsers method")
// 			},
// 			UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error {
// 				panic("mock out the UpdateUser method")
// 			},
// 		}
//
// 		// use mockedUserService in code that requires userService
// 		// and then make assertions.
//
// 	}
type UserServiceMock struct {
	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(ctx context.Context, nu models.NewUser) (uuid.UUID, error)

	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(ctx context.Context, userID uuid.UUID) error

	// GetUsersFunc mocks the GetUsers method.
	GetUsersFunc func(ctx context.Context, qu models.GetUsersOptions) ([]models.User, error)

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Nu is the nu argument value.
			Nu models.NewUser
		}
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// GetUsers holds details about calls to the GetUsers method.
		GetUsers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Qu is the qu argument value.
			Qu models.GetUsersOptions
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
			// Uu is the uu argument value.
			Uu models.UpdateUser
		}
	}
	lockCreateUser sync.RWMutex
	lockDeleteUser sync.RWMutex
	lockGetUsers   sync.RWMutex
	lockUpdateUser sync.RWMutex
}

// CreateUser calls CreateUserFunc.
func (mock *UserServiceMock) CreateUser(ctx context.Context, nu models.NewUser) (uuid.UUID, error) {
	if mock.CreateUserFunc == nil {
		panic("UserServiceMock.CreateUserFunc: method is nil but userService.CreateUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Nu  models.NewUser
	}{
		Ctx: ctx,
		Nu:  nu,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	return mock.CreateUserFunc(ctx, nu)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//     len(mockedUserService.CreateUserCalls())
func (mock *UserServiceMock) CreateUserCalls() []struct {
	Ctx context.Context
	Nu  models.NewUser
} {
	var calls []struct {
		Ctx context.Context
		Nu  models.NewUser
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}

// DeleteUser calls DeleteUserFunc.
func (mock *UserServiceMock) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if mock.DeleteUserFunc == nil {
		panic("UserServiceMock.DeleteUserFunc: method is nil but userService.DeleteUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	mock.lockDeleteUser.Unlock()
	return mock.DeleteUserFunc(ctx, userID)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//     len(mockedUserService.DeleteUserCalls())
func (mock *UserServiceMock) DeleteUserCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	mock.lockDeleteUser.RUnlock()
	return calls
}

// GetUsers calls GetUsersFunc.
func (mock *UserServiceMock) GetUsers(ctx context.Context, qu models.GetUsersOptions) ([]models.User, error) {
	if mock.GetUsersFunc == nil {
		panic("UserServiceMock.GetUsersFunc: method is nil but userService.GetUsers was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Qu  models.GetUsersOptions
	}{
		Ctx: ctx,
		Qu:  qu,
	}
	mock.lockGetUsers.Lock()
	mock.calls.GetUsers = append(mock.calls.GetUsers, callInfo)
	mock.lockGetUsers.Unlock()
	return mock.GetUsersFunc(ctx, qu)
}

// GetUsersCalls gets all the calls that were made to GetUsers.
// Check the length with:
//     len(mockedUserService.GetUsersCalls())
func (mock *UserServiceMock) GetUsersCalls() []struct {
	Ctx context.Context
	Qu  models.GetUsersOptions
} {
	var calls []struct {
		Ctx context.Context
		Qu  models.GetUsersOptions
	}
	mock.lockGetUsers.RLock()
	calls = mock.calls.GetUsers
	mock.lockGetUsers.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *UserServiceMock) UpdateUser(ctx context.Context, userID uuid.UUID, uu models.UpdateUser) error {
	if mock.UpdateUserFunc == nil {
		panic("UserServiceMock.UpdateUserFunc: method is nil but userService.UpdateUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
		Uu     models.UpdateUser
	}{
		Ctx:    ctx,
		UserID: userID,
		Uu:     uu,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(ctx, userID, uu)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//     len(mockedUserService.UpdateUserCalls())
func (mock *UserServiceMock) UpdateUserCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
	Uu     models.UpdateUser
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
		Uu     models.UpdateUser
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
