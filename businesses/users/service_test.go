package users_test

import (
	"errors"
	"fmt"
	"github.com/avtara/travair-api/businesses"
	_cacheMock "github.com/avtara/travair-api/businesses/cache/mocks"
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
	mockCacheRepository _cacheMock.Repository
	mockQueueRepository _queueMock.Repository
	userService users.Service

	domainTest users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewUserService(&mockUsersRepository,time.Second*1,&mockQueueRepository,&mockCacheRepository)
	domainTest = users.Domain{
		Email: "avtara@gmail.com",
		Role: "tenant",
		Password: "Avtara12!",
		Name: "Muhammad Avtara Khrisna",
	}
	m.Run()
}

func TestUserService_Registration(t *testing.T) {
	t.Run("error duplicate data", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.Anything).Return(&domainTest, nil).Once()
		req := &users.Domain{
			Email: "avtara@gmail.com",
			Role: "tenant",
			Password: "Avtara12!",
			Name: "Muhammad Avtara Khrisna",
		}
		_, err := userService.Registration(context.Background(), req)
		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("Valid payload", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.Anything).Return(nil, nil).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything,mock.Anything).Return(&domainTest, nil).Once()
		mockCacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", nil).Once()
		mockQueueRepository.On("Publish", mock.Anything, mock.Anything).Return(nil, nil).Once()

		req := &users.Domain{
			Email: "avtara@gmail.com",
			Role: "tenant",
			Password: "Avtara12!",
			Name: "Muhammad Avtara Khrisna",
		}

		res, err := userService.Registration(context.Background(), req)
		fmt.Println(res)
		assert.Nil(t, err)
		assert.Equal(t, req.Email, domainTest.Email)
	})

	t.Run("invalid payload", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.Anything).Return(nil, nil).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything,mock.Anything).Return(nil, errors.New("awd")).Once()

		req := &users.Domain{
			Email: "avtara@gmail.com",
			Role: "tenant",
			Password: "Avtara12!",
			Name: "Muhammad Avtara Khrisna",
		}

		_, err := userService.Registration(context.Background(), req)
		assert.Equal(t, businesses.ErrInternalServer,err)
	})

	t.Run("Valid payload", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.Anything).Return(nil, nil).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything,mock.Anything).Return(&domainTest, nil).Once()
		mockCacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", nil).Once()
		mockQueueRepository.On("Publish", mock.Anything, mock.Anything).Return(nil, nil).Once()

		req := &users.Domain{
			Email: "avtara@gmail.com",
			Role: "tenant",
			Password: "Avtara12!",
			Name: "Muhammad Avtara Khrisna",
		}

		res, err := userService.Registration(context.Background(), req)
		fmt.Println(res)
		assert.Nil(t, err)
		assert.Equal(t, req.Email, domainTest.Email)
	})

	t.Run("Cache Error", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.Anything).Return(nil, nil).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything,mock.Anything).Return(&domainTest, nil).Once()
		mockCacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("",errors.New("awd")).Once()

		req := &users.Domain{
			Email: "avtara@gmail.com",
		}

		_, err := userService.Registration(context.Background(), req)
		assert.Equal(t, businesses.ErrInternalServer, err)
	})

	t.Run("Queue Error", func(t *testing.T) {
		mockUsersRepository.On("GetByEmail", mock.Anything,mock.Anything).Return(nil, nil).Once()
		mockUsersRepository.On("StoreNewUsers", mock.Anything,mock.Anything).Return(&domainTest, nil).Once()
		mockCacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("",nil).Once()
		mockQueueRepository.On("Publish", mock.Anything, mock.Anything).Return(errors.New("awd")).Once()


		req := &users.Domain{
			Email: "avtara@gmail.com",
		}

		_, err := userService.Registration(context.Background(), req)
		assert.Equal(t, businesses.ErrInternalServer,err)
	})
}