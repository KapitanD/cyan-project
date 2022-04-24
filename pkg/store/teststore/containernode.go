package teststore

import (
	"errors"
	"strings"

	"github.com/KapitanD/cyan-project/pkg/model"
)

type containerNode struct {
	parent *containerNode
	leafs  map[string]*containerNode

	element *model.Container
}

func newTree() *containerNode {
	return &containerNode{
		parent: nil,
		leafs:  make(map[string]*containerNode),

		element: &model.Container{
			Name:   "",
			IsLeaf: false,
			Meta:   make(map[string]string),
		},
	}
}

func (node *containerNode) getByPath(path string) (*containerNode, error) {
	pathSlice := strings.Split(path, "/")
	curNode := node

	if len(pathSlice) < 2 || pathSlice[0] != "" {
		return nil, errors.New("invalid path")
	}

	if len(pathSlice[1:]) == 1 && pathSlice[1] == "" {
		return curNode, nil
	}

	for _, elem := range pathSlice[1:] {
		if _, ok := curNode.leafs[elem]; !ok {
			return nil, errors.New("invalid path")
		}

		curNode = curNode.leafs[elem]
	}

	return curNode, nil
}

func (node *containerNode) createContainer(new *model.Container) error {
	if _, ok := node.leafs[new.Name]; ok {
		return errors.New("duplicate name")
	}

	newNode := &containerNode{
		parent:  node,
		leafs:   make(map[string]*containerNode),
		element: new,
	}

	node.leafs[new.Name] = newNode

	return nil
}
