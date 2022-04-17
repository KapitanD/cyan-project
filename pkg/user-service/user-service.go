package user_service

import (
	"context"
	"errors"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/KapitanD/cyan-project/pkg/model"
	"github.com/KapitanD/cyan-project/pkg/store"
	"github.com/KapitanD/cyan-project/pkg/store/teststore"
)

type userService struct {
	v1.UnimplementedUserServiceServer
	store        store.Store
	tokenService *TokenService
}

func NewUserService() *userService {
	return &userService{
		store:        teststore.New(),
		tokenService: &TokenService{},
	}
}

func (us *userService) CreateUser(ctx context.Context, r *v1.CreateRequest) (*v1.CreateResponse, error) {
	user := &model.User{
		Email:    r.Email,
		Password: r.Password,
	}
	err := us.store.User().CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &v1.CreateResponse{
		User: &v1.UserIdentity{
			UserId: user.Id,
			Email:  user.Email,
		},
		Created: true,
	}, nil
}

func (us *userService) AuthenticateUser(ctx context.Context, r *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	user, err := us.store.User().FindByEmail(r.Email)
	if err != nil {
		return nil, err
	}

	if !user.ComparePassword(r.Password) {
		return nil, errors.New("invalid password")
	}

	tokenPair, err := us.tokenService.TokenPair(user)
	if err != nil {
		return nil, err
	}

	return &v1.AuthenticateResponse{
		Token:        tokenPair["access_token"],
		RefreshToken: tokenPair["refresh_token"],
	}, nil
}

func (us *userService) AuthorizeUser(ctx context.Context, r *v1.AuthorizeRequest) (*v1.AuthorizeResponse, error) {
	claims, err := us.tokenService.Decode(r.Token)
	if err != nil {
		return nil, err
	}

	return &v1.AuthorizeResponse{
		User: &v1.UserIdentity{
			UserId: claims.User.Id,
			Email:  claims.User.Email,
		},
	}, nil
}
