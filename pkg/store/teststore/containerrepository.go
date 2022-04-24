package teststore

import (
	"errors"

	"github.com/KapitanD/cyan-project/pkg/model"
)

type ContainerRepository struct {
	store *ContainerStore
	trees map[string]*containerNode
}

func (cr *ContainerRepository) CreateRoot(ownerEmail string) error {
	if _, ok := cr.trees[ownerEmail]; ok {
		return errors.New("root for user already exists")
	}

	cr.trees[ownerEmail] = newTree()
	return nil
}

func (cr *ContainerRepository) CreateContainer(owner *model.User, path string, new *model.Container) error {
	parent, err := cr.trees[owner.Email].getByPath(path)
	if err != nil {
		return err
	}

	err = parent.createContainer(new)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ContainerRepository) FindByPath(owner *model.User, path string) (*model.Container, error) {
	containerNode, err := cr.trees[owner.Email].getByPath(path)
	if err != nil {
		return nil, err
	}

	return containerNode.element, nil
}

func (cr *ContainerRepository) GetInners(owner *model.User, path string) ([]*model.Container, error) {
	containerNode, err := cr.trees[owner.Email].getByPath(path)
	if err != nil {
		return nil, err
	}

	inners := make([]*model.Container, 0)

	for _, v := range containerNode.leafs {
		inners = append(inners, v.element)
	}

	return inners, nil
}
