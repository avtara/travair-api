package users

import (
	"context"
	"github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/businesses"
	"github.com/avtara/travair-api/businesses/queue"
	"github.com/avtara/travair-api/helpers"
	"github.com/google/uuid"
	"strings"
	"time"
)

type userService struct {
	userRepository Repository
	contextTimeout time.Duration
	queueRepo      queue.Repository
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(rep Repository, timeout time.Duration, queueRep queue.Repository, jwt *middleware.ConfigJWT) Service {
	return &userService{
		userRepository: rep,
		contextTimeout: timeout,
		queueRepo:      queueRep,
		jwtAuth:        jwt,
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
		return nil, businesses.ErrEmailDuplicate
	}

	userDomain.Password, err = helpers.HashPassword(userDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.userRepository.StoreNewUsers(ctx, userDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	us.queueRepo.EmailUsers(res.UserID, res.Name, res.Email, "registration")

	return res, nil
}

func (us *userService) Activation(ctx context.Context, userID string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.userRepository.GetByUserID(ctx, uuidUserID)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		} else {
			return nil, businesses.ErrAccountNotFound
		}
	}

	if res.Status == 1 {
		return nil, businesses.ErrAccountActivated
	}

	if err = us.userRepository.UpdateStatus(ctx, uuidUserID); err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *userService) Login(ctx context.Context, email, password string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	password, err := helpers.HashPassword(password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.userRepository.GetByEmailAndPassword(ctx, email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
		return nil, businesses.ErrAccountNotFound
	}

	if helpers.ValidateHash(password, res.Password) {
		return nil, businesses.ErrInvalidCredential
	}

	if res.Status != 1 {
		return nil, businesses.ErrAccountUnactivated
	}
	res.Token = us.jwtAuth.GenerateToken(res.UserID, res.Role)
	return res, nil
}

func (us *userService) GetID(ctx context.Context, userID uuid.UUID) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	res, _ := us.userRepository.GetByUserID(ctx, userID)

	return res.ID, nil
}
