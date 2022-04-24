package container_service

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/KapitanD/cyan-project/internal/config"
	"github.com/KapitanD/cyan-project/internal/constants"
	"github.com/KapitanD/cyan-project/pkg/model"
	"github.com/KapitanD/cyan-project/pkg/store/teststore"
	"github.com/KapitanD/cyan-project/pkg/utils"
)

type containerService struct {
	v1.UnimplementedContainerServiceServer
	store       *teststore.ContainerStore
	userService v1.UserServiceClient
}

func NewContainerService(config *config.ContainerServiceConfig) (*containerService, error) {
	userConn, err := grpc.Dial(config.UserServiceEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &containerService{
		store:       teststore.NewContainerStore(),
		userService: v1.NewUserServiceClient(userConn),
	}, nil
}

func (cs *containerService) CreateRootContainer(ctx context.Context, in *emptypb.Empty) (*v1.ContainerRootResponse, error) {
	userEmail, err := utils.UnaryValueFromMetadata(ctx, constants.MdKeyUserEmail)
	if err != nil {
		return nil, errors.Wrap(err, "get email from context")
	}

	err = cs.store.Container().CreateRoot(userEmail)
	if err != nil {
		return nil, errors.Wrap(err, "create root err")
	}

	return &v1.ContainerRootResponse{
		Created: true,
	}, nil
}

func (cs *containerService) CreateContainer(ctx context.Context, req *v1.ContainerCreateRequest) (*v1.ContainerCreateResponse, error) {
	user, err := cs.authReq(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user from context")
	}

	container := model.UnmarshalContainer(req.Container)
	container.Owner = user
	if err := cs.store.Container().CreateContainer(user, req.Path, container); err != nil {
		return nil, errors.Wrap(err, "container creation error")
	}

	return &v1.ContainerCreateResponse{
		Created:   true,
		Container: model.MarshallContainer(container),
	}, nil
}

func (cs *containerService) GetContainer(ctx context.Context, req *v1.ContainerGetRequest) (*v1.ContainerGetResponse, error) {
	user, err := cs.authReq(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user from context")
	}

	container, err := cs.store.Container().FindByPath(user, req.Path)
	if err != nil {
		return nil, errors.Wrapf(err, "error while looking path %s", req.Path)
	}

	return &v1.ContainerGetResponse{
		Container: model.MarshallContainer(container),
	}, nil
}

func (cs *containerService) GetInners(ctx context.Context, req *v1.ContainerGetRequest) (*v1.ContainerInnersResponse, error) {
	user, err := cs.authReq(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user from context")
	}

	inners, err := cs.store.Container().GetInners(user, req.Path)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting inners")
	}

	return &v1.ContainerInnersResponse{
		Inners: model.MarshallContainers(inners),
	}, nil
}

func (cs *containerService) authReq(ctx context.Context) (*model.User, error) {
	token, err := utils.UnaryValueFromMetadata(ctx, constants.MdKeyAuthToken)
	if err != nil {
		return nil, errors.Wrap(err, "getting token from context")
	}

	resp, err := cs.userService.AuthorizeUser(
		ctx,
		&v1.AuthorizeReq{
			Token: token,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while authorizing in container service")
	}

	return model.UnmarshalUser(resp.User), nil
}
