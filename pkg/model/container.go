package model

import (
	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
)

type Container struct {
	Name   string            `json:"name"`
	Meta   map[string]string `json:"meta"`
	IsLeaf bool              `json:"is_leaf"`
	Owner  *User             `json:"-"`
}

func MarshallContainer(cont *Container) *v1.Container {
	return &v1.Container{
		Name:     cont.Name,
		IsLeaf:   cont.IsLeaf,
		MetaData: cont.Meta,
	}
}

func UnmarshalContainer(cont *v1.Container) *Container {
	return &Container{
		Name:   cont.Name,
		Meta:   cont.MetaData,
		IsLeaf: cont.IsLeaf,
	}
}

func MarshallContainers(conts []*Container) []*v1.Container {
	marshalled := make([]*v1.Container, 0)
	for _, cont := range conts {
		marshalled = append(marshalled, MarshallContainer(cont))
	}

	return marshalled
}

func UnmarshalContainers(conts []*v1.Container) []*Container {
	unmarshaled := make([]*Container, 0)
	for _, cont := range conts {
		unmarshaled = append(unmarshaled, UnmarshalContainer(cont))
	}

	return unmarshaled
}
