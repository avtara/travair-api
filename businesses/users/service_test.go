package users_test

import (
	"errors"
	"github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/businesses"
	_queueMock "github.com/avtara/travair-api/businesses/queue/mocks"
	"github.com/avtara/travair-api/businesses/users"
	_usersMock "github.com/avtara/travair-api/businesses/users/mocks"
	"github.com/google/uuid"
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
	userService = users.NewUserService(&mockUsersRepository, time.Second*1, &mockQueueRepository, &middleware.ConfigJWT{})
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
		assert.Nil(t, err)
		assert.Equal(t, res.Email, domainTest.Email)
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

		_, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
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

		res, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac130002")
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
		assert.Nil(t, req)
		assert.Equal(t, businesses.ErrAccountNotFound, err)
	})

	t.Run("error parse uuid", func(t *testing.T) {
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

		req, err := userService.Activation(context.Background(), "c5c838ba-24f8-11ec-9621-0242ac1300")
		assert.Nil(t, req)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})
}

func TestUserService_GetID(t *testing.T) {
	t.Run("success get ID", func(t *testing.T) {
		activeAccount := &users.Domain{
			ID:       2,
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status:   0,
		}

		mockUsersRepository.On("GetByUserID", mock.Anything, mock.Anything).Return(activeAccount, nil).Once()

		uuidUserID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")

		_, err := userService.GetID(context.Background(), uuidUserID)
		assert.Nil(t, err)
	})
}

func TestUserService_Login(t *testing.T) {
	t.Run("fail login", func(t *testing.T) {
		activeAccount := &users.Domain{
			ID:       2,
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status:   0,
		}

		mockUsersRepository.On("GetByEmailAndPassword", mock.Anything, mock.Anything).Return(nil, errors.New("SQL: not found")).Once()


		res, err := userService.Login(context.Background(), activeAccount.Email, activeAccount.Password)
		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrAccountNotFound, err)
	})

	t.Run("error sql", func(t *testing.T) {
		activeAccount := &users.Domain{
			ID:       2,
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status:   0,
		}

		mockUsersRepository.On("GetByEmailAndPassword", mock.Anything, mock.Anything).Return(nil, errors.New("SQL: sql bla bla")).Once()


		res, _ := userService.Login(context.Background(), activeAccount.Email, activeAccount.Password)
		assert.Nil(t, res)
	})

	t.Run("unactivate account", func(t *testing.T) {
		account := &users.Domain{
			ID:       2,
			Email:    "avtara@gmail.com",
			Role:     "tenant",
			Password: "Avtara12!",
			Name:     "Muhammad Avtara Khrisna",
			Status:   0,
		}

		mockUsersRepository.On("GetByEmailAndPassword", mock.Anything, mock.Anything).Return(account, nil).Once()


		res, err := userService.Login(context.Background(), account.Email, account.Password)
		assert.Nil(t, res)
		assert.Equal(t, businesses.ErrAccountUnactivated, err)
	})
}