package store

type UserStore interface {
	User() UserRepository
}

type ContainerStore interface {
	Container() ContainerRepository
}
