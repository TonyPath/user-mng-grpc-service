// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	"github.com/google/uuid"
	"sync"
)

// Ensure, that UserStorageMock does implement UserStorage.
// If this is not the case, regenerate this file with moq.
var _ UserStorage = &UserStorageMock{}

// UserStorageMock is a mock implementation of UserStorage.
//
// 	func TestSomethingThatUsesUserStorage(t *testing.T) {
//
// 		// make and configure a mocked UserStorage
// 		mockedUserStorage := &UserStorageMock{
// 			DeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
// 				panic("mock out the DeleteUser method")
// 			},
// 			ExistsByIDFunc: func(ctx context.Context, userID uuid.UUID) (bool, error) {
// 				panic("mock out the ExistsByID method")
// 			},
// 			GetUserByIDFunc: func(ctx context.Context, userID uuid.UUID) (models.User, error) {
// 				panic("mock out the GetUserByID method")
// 			},
// 			GetUsersByFilterFunc: func(ctx context.Context, opts models.GetUsersOptions) ([]models.User, error) {
// 				panic("mock out the GetUsersByFilter method")
// 			},
// 			InsertUserFunc: func(ctx context.Context, user models.User) (uuid.UUID, error) {
// 				panic("mock out the InsertUser method")
// 			},
// 			UpdateUserFunc: func(ctx context.Context, userID uuid.UUID, user models.User) error {
// 				panic("mock out the UpdateUser method")
// 			},
// 		}
//
// 		// use mockedUserStorage in code that requires UserStorage
// 		// and then make assertions.
//
// 	}
type UserStorageMock struct {
	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(ctx context.Context, userID uuid.UUID) error

	// ExistsByIDFunc mocks the ExistsByID method.
	ExistsByIDFunc func(ctx context.Context, userID uuid.UUID) (bool, error)

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(ctx context.Context, userID uuid.UUID) (models.User, error)

	// GetUsersByFilterFunc mocks the GetUsersByFilter method.
	GetUsersByFilterFunc func(ctx context.Context, opts models.GetUsersOptions) ([]models.User, error)

	// InsertUserFunc mocks the InsertUser method.
	InsertUserFunc func(ctx context.Context, user models.User) (uuid.UUID, error)

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(ctx context.Context, userID uuid.UUID, user models.User) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// ExistsByID holds details about calls to the ExistsByID method.
		ExistsByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// GetUsersByFilter holds details about calls to the GetUsersByFilter method.
		GetUsersByFilter []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts models.GetUsersOptions
		}
		// InsertUser holds details about calls to the InsertUser method.
		InsertUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User models.User
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
			// User is the user argument value.
			User models.User
		}
	}
	lockDeleteUser       sync.RWMutex
	lockExistsByID       sync.RWMutex
	lockGetUserByID      sync.RWMutex
	lockGetUsersByFilter sync.RWMutex
	lockInsertUser       sync.RWMutex
	lockUpdateUser       sync.RWMutex
}

// DeleteUser calls DeleteUserFunc.
func (mock *UserStorageMock) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if mock.DeleteUserFunc == nil {
		panic("UserStorageMock.DeleteUserFunc: method is nil but UserStorage.DeleteUser was just called")
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
//     len(mockedUserStorage.DeleteUserCalls())
func (mock *UserStorageMock) DeleteUserCalls() []struct {
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

// ExistsByID calls ExistsByIDFunc.
func (mock *UserStorageMock) ExistsByID(ctx context.Context, userID uuid.UUID) (bool, error) {
	if mock.ExistsByIDFunc == nil {
		panic("UserStorageMock.ExistsByIDFunc: method is nil but UserStorage.ExistsByID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockExistsByID.Lock()
	mock.calls.ExistsByID = append(mock.calls.ExistsByID, callInfo)
	mock.lockExistsByID.Unlock()
	return mock.ExistsByIDFunc(ctx, userID)
}

// ExistsByIDCalls gets all the calls that were made to ExistsByID.
// Check the length with:
//     len(mockedUserStorage.ExistsByIDCalls())
func (mock *UserStorageMock) ExistsByIDCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockExistsByID.RLock()
	calls = mock.calls.ExistsByID
	mock.lockExistsByID.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *UserStorageMock) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("UserStorageMock.GetUserByIDFunc: method is nil but UserStorage.GetUserByID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	mock.lockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(ctx, userID)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//     len(mockedUserStorage.GetUserByIDCalls())
func (mock *UserStorageMock) GetUserByIDCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	mock.lockGetUserByID.RUnlock()
	return calls
}

// GetUsersByFilter calls GetUsersByFilterFunc.
func (mock *UserStorageMock) GetUsersByFilter(ctx context.Context, opts models.GetUsersOptions) ([]models.User, error) {
	if mock.GetUsersByFilterFunc == nil {
		panic("UserStorageMock.GetUsersByFilterFunc: method is nil but UserStorage.GetUsersByFilter was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts models.GetUsersOptions
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockGetUsersByFilter.Lock()
	mock.calls.GetUsersByFilter = append(mock.calls.GetUsersByFilter, callInfo)
	mock.lockGetUsersByFilter.Unlock()
	return mock.GetUsersByFilterFunc(ctx, opts)
}

// GetUsersByFilterCalls gets all the calls that were made to GetUsersByFilter.
// Check the length with:
//     len(mockedUserStorage.GetUsersByFilterCalls())
func (mock *UserStorageMock) GetUsersByFilterCalls() []struct {
	Ctx  context.Context
	Opts models.GetUsersOptions
} {
	var calls []struct {
		Ctx  context.Context
		Opts models.GetUsersOptions
	}
	mock.lockGetUsersByFilter.RLock()
	calls = mock.calls.GetUsersByFilter
	mock.lockGetUsersByFilter.RUnlock()
	return calls
}

// InsertUser calls InsertUserFunc.
func (mock *UserStorageMock) InsertUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	if mock.InsertUserFunc == nil {
		panic("UserStorageMock.InsertUserFunc: method is nil but UserStorage.InsertUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User models.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockInsertUser.Lock()
	mock.calls.InsertUser = append(mock.calls.InsertUser, callInfo)
	mock.lockInsertUser.Unlock()
	return mock.InsertUserFunc(ctx, user)
}

// InsertUserCalls gets all the calls that were made to InsertUser.
// Check the length with:
//     len(mockedUserStorage.InsertUserCalls())
func (mock *UserStorageMock) InsertUserCalls() []struct {
	Ctx  context.Context
	User models.User
} {
	var calls []struct {
		Ctx  context.Context
		User models.User
	}
	mock.lockInsertUser.RLock()
	calls = mock.calls.InsertUser
	mock.lockInsertUser.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *UserStorageMock) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) error {
	if mock.UpdateUserFunc == nil {
		panic("UserStorageMock.UpdateUserFunc: method is nil but UserStorage.UpdateUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
		User   models.User
	}{
		Ctx:    ctx,
		UserID: userID,
		User:   user,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(ctx, userID, user)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//     len(mockedUserStorage.UpdateUserCalls())
func (mock *UserStorageMock) UpdateUserCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
	User   models.User
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
		User   models.User
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
