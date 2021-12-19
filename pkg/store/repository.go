package store

type UserRepository interface {
	CreateUser(login, password string) error
	FindByLoginAndPassword(login, password string) bool
}
