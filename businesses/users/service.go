package users

import (
	"context"
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/queue"
	"github.com/avtara/travair-api/helpers"
	"strings"
	"time"
)

type userService struct {
	userRepository Repository
	contextTimeout time.Duration
	queueRepo queue.Repository
}

func NewUserService(rep Repository, timeout time.Duration, queueRep queue.Repository) Service {
	return &userService{
		userRepository: rep,
		contextTimeout: timeout,
		queueRepo: queueRep,
	}
}

func (us *userService) Registration(ctx context.Context, userDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	existedUser, err := us.userRepository.GetByEmail(ctx, userDomain)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}
	if existedUser != nil {
		return nil, businesses.ErrDuplicateData
	}

	userDomain.Password, err = helpers.HashPassword(userDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.userRepository.StoreNewUsers(ctx, userDomain)
	if err != nil {
		return nil, err
	}


	us.queueRepo.Publish("email:registration", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}