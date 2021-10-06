package users_test

import (
	"errors"
	"fmt"
	"github.com/avtara/travair-api/businesses"
	_queueMock "github.com/avtara/travair-api/businesses/queue/mocks"
	"github.com/avtara/travair-api/businesses/users"
	_usersMock "github.com/avtara/travair-api/businesses/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"testing"
	"time"
)

var (
	mockUsersRepository _usersMock.Repository
	mockQueueRepository _queueMock.Repository
	userService         users.Service

	domainTest users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewUserService(&mockUsersRepository, time.Second*1, &mockQueueRepository, &mockJWTRepos)
	domainTest = users.Domain{
		Email:    "avtara@gmail.com",
		Role:     "tenant",
		Password: "Avtara12!",
		Name:     "Muhammad Avtara Khrisna",
	}
	m.Run()
}

func TestUserService_Registration(t *testing.T) {
	t.Run("Valid payload", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,
			mock.Anything).Return(nil, nil).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything,
			mock.Anything).Return(&domainTest, nil).Once()
		mockQueueRepository.On("EmailUsers", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Once()

		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
		}

		res, err := userService.Registration(context.Background(), req)
		fmt.Println(res)
		assert.Nil(t, err)
		assert.Equal(t, req.Email, domainTest.Email)
	})

	t.Run("Error while search account using email", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("SQL: cannot bla bla")).Once()

		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
		}

		res, err := userService.Registration(context.Background(), req)
		assert.Nil(t, res)
		assert.Equal(t, err, errors.New("SQL: cannot bla bla"))
	})

	t.Run("Error when using existed account", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()

		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
		}

		res, err := userService.Registration(context.Background(), req)
		assert.Nil(t, res)
		assert.Equal(t, err, businesses.ErrEmailDuplicate)
	})

	t.Run("Error when store user", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(nil, errors.New("SQL: not found")).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything, mock.Anything).Return(nil, errors.New("SQL: blabla")).Once()

		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
		}

		res, err := userService.Registration(context.Background(), req)
		assert.Nil(t, res)
		assert.Equal(t, err, businesses.ErrInternalServer)
	})
}

func TestUserService_Activation(t *testing.T) {
	t.Run("success while activation account", func(t *testing.T) {
		activeAccount := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 0,
		}

		mockUsersRepository.On("GetByUserID", mock.Anything, mock.Anything).Return(activeAccount, nil).Once()
		mockUsersRepository.On("UpdateStatus", mock.Anything, mock.Anything).Return(nil).Once()
		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 1,
		}

		_, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
		fmt.Println(req)
		assert.Nil(t, err)
	})

	t.Run("error when activating an account that has been activated", func(t *testing.T) {
		activeAccount := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 1,
		}

		mockUsersRepository.On("GetByUserID", mock.Anything, mock.Anything).Return(activeAccount, nil).Once()
		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 1,
		}

		res, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
		fmt.Println(req)
		assert.Equal(t, err, businesses.ErrAccountActivated)
		assert.Nil(t, res)
	})

	t.Run("error while activation account", func(t *testing.T) {
		activeAccount := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 0,
		}

		mockUsersRepository.On("GetByUserID", mock.Anything, mock.Anything).Return(activeAccount, nil).Once()
		mockUsersRepository.On("UpdateStatus", mock.Anything, mock.Anything).Return(errors.New("SQL: err")).Once()
		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 1,
		}

		req, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
		fmt.Println(req)
		assert.Nil(t, req)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("error while search account", func(t *testing.T) {
		activeAccount := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 0,
		}

		mockUsersRepository.On("GetByUserID", mock.Anything, mock.Anything).Return(activeAccount, errors.New("SQL: not found")).Once()
		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 1,
		}

		req, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
		fmt.Println(req)
		assert.Nil(t, req)
		assert.Equal(t, businesses.ErrAccountNotFound, err)
	})

	t.Run("error while search account", func(t *testing.T) {
		activeAccount := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 0,
		}

		mockUsersRepository.On("GetByUserID", mock.Anything, mock.Anything).Return(activeAccount, errors.New("SQL: bla bla")).Once()
		req := &users.Domain{
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status: 1,
		}

		req, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
		fmt.Println(req)
		assert.Nil(t, req)
		assert.Equal(t, errors.New("SQL: bla bla"), err)
	})
}