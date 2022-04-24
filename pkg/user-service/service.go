package user_service

import (
	"context"
	"time"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/KapitanD/cyan-project/pkg/model"
	"github.com/KapitanD/cyan-project/pkg/store"
	"github.com/KapitanD/cyan-project/pkg/store/teststore"
	"github.com/KapitanD/cyan-project/pkg/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	v1.UnimplementedUserServiceServer
	store        store.UserStore
	tokenService *TokenService
	commitMap    utils.TTLMapI
}

func NewUserService(ctx context.Context) *userService {
	return &userService{
		store:        teststore.NewUserStore(),
		tokenService: &TokenService{},
		commitMap:    utils.NewTTLMap(ctx),
	}
}

func (us *userService) CreateUserTry(ctx context.Context, r *v1.UserCreateReq) (*v1.Commit, error) {
	user := &model.User{
		Email:    r.Email,
		Password: r.Password,
	}

	commitId := uuid.New().String()

	err := user.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "validate while creating user")
	}

	us.commitMap.Set(commitId, user, 30*time.Minute)

	return &v1.Commit{
		CommitId: commitId,
	}, nil
}

func (us *userService) CreateUserCommit(ctx context.Context, r *v1.Commit) (*v1.UserCreateResp, error) {
	mapItem, err := us.commitMap.Get(r.CommitId)
	if err != nil {
		return nil, errors.Wrap(err, "commit map get")
	}

	user, ok := mapItem.(*model.User)
	if !ok {
		return nil, errors.New("cannot convert mapitem to user")
	}

	err = us.store.User().CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &v1.UserCreateResp{
		UserId: model.MarshallUser(user),
	}, nil
}

func (us *userService) CreateUserCancel(ctx context.Context, r *v1.Commit) (*emptypb.Empty, error) {
	err := us.commitMap.Del(r.CommitId)
	if err != nil {
		return nil, errors.Wrap(err, "commit map del")
	}

	return &emptypb.Empty{}, nil
}

func (us *userService) AuthenticateUser(ctx context.Context, r *v1.AuthenticateReq) (*v1.AuthenticateResp, error) {
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

	return &v1.AuthenticateResp{
		Token:        tokenPair["access_token"],
		RefreshToken: tokenPair["refresh_token"],
	}, nil
}

func (us *userService) AuthorizeUser(ctx context.Context, r *v1.AuthorizeReq) (*v1.AuthorizeResp, error) {
	claims, err := us.tokenService.Decode(r.Token)
	if err != nil {
		return nil, err
	}

	return &v1.AuthorizeResp{
		User: &v1.UserIdentity{
			UserId: claims.User.Id,
			Email:  claims.User.Email,
		},
	}, nil
}
