package userservice

import (
	"context"
	"errors"

	"github.com/google/uuid"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/KapitanD/cyan-project/pkg/store"
)

type userService struct {
	v1.UnimplementedUserServiceServer
	store        *store.Store
	tokenService *TokenService
}

func (us *userService) CreateUser(ctx context.Context, r *v1.CreateRequest) (*v1.CreateResponse, error) {
	if r.Password != r.RepeatPassword {
		return nil, errors.New("passwords are not equals")
	}

	err := us.store.User().CreateUser(r.Login, r.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()

	return &v1.CreateResponse{
		UserId: &v1.UserIdentity{
			UserId: id,
		},
		Created: true,
	}, nil
}

func (us *userService) LoginUser(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {

}
