package users

import (
	"context"
	"fmt"
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/cache"
	"github.com/avtara/travair-api/businesses/queue"
	"github.com/avtara/travair-api/helpers"
	"strings"
	"time"
)

type userService struct {
	userRepository Repository
	contextTimeout time.Duration
	queueRepo queue.Repository
	cacheRepo cache.Repository
}

func NewUserService(rep Repository, timeout time.Duration, queueRep queue.Repository, cacheRep cache.Repository) Service {
	return &userService{
		userRepository: rep,
		contextTimeout: timeout,
		queueRepo: queueRep,
		cacheRepo: cacheRep,
	}
}

func (us *userService) Registration(ctx context.Context, userDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	existedUser, err := us.userRepository.GetByEmail(ctx, userDomain.Email)
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
		return nil, businesses.ErrInternalServer
	}

	token:= helpers.RandomToken(25)
	keyRedis := fmt.Sprintf("confirm_email:%s",res.Email)
	_, err = us.cacheRepo.Set(ctx, keyRedis, token, time.Duration(60 * 5))
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	err = us.queueRepo.Publish("email:registration", res)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}
	return res, nil
}
