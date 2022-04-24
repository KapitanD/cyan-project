package store

import (
	"github.com/KapitanD/cyan-project/pkg/model"
)

type UserRepository interface {
	CreateUser(*model.User) error
	FindByEmail(email string) (*model.User, error)
}

type ContainerRepository interface {
	CreateRoot(ownerEmail string) error
	CreateContainer(owner *model.User, path string, new *model.Container) error
	FindByPath(owner *model.User, path string) (*model.Container, error)
	GetInners(owner *model.User, path string) ([]*model.Container, error)
}
