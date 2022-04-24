package teststore

import (
	"github.com/KapitanD/cyan-project/pkg/model"
	"github.com/KapitanD/cyan-project/pkg/store"
)

type UserStore struct {
	userRepository *UserRepository
}

type ContainerStore struct {
	containerRepository *ContainerRepository
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}

func NewContainerStore() *ContainerStore {
	return &ContainerStore{}
}

func (s *ContainerStore) Container() store.ContainerRepository {
	if s.containerRepository != nil {
		return s.containerRepository
	}

	s.containerRepository = &ContainerRepository{
		store: s,
		trees: make(map[string]*containerNode),
	}

	return s.containerRepository
}
