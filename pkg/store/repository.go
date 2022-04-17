package store

import (
	"github.com/KapitanD/cyan-project/pkg/model"
)

type UserRepository interface {
	CreateUser(*model.User) error
	FindByEmail(email string) (*model.User, error)
}
