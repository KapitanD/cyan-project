package teststore

import (
	"errors"

	"github.com/KapitanD/cyan-project/pkg/model"
)

type UserRepository struct {
	store *UserStore
	users map[string]*model.User
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	_, ok := ur.users[user.Email]
	if ok {
		return errors.New("user already exist")
	}

	ur.users[user.Email] = user
	return nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	_, ok := ur.users[email]
	if ok {
		return ur.users[email], nil
	} else {
		return nil, errors.New("user do not exists")
	}
}
