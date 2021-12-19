package teststore

import "errors"

type UserRepository struct {
	store *Store
	users map[string]string
}

func (ur *UserRepository) CreateUser(login, password string) error {
	_, ok := ur.users[login]
	if ok {
		return errors.New("user already exist")
	}

	ur.users[login] = password
	return nil
}

func (ur *UserRepository) FindByLoginAndPassword(login, password string) bool {
	realPassword, ok := ur.users[login]
	if ok {
		return realPassword == password
	} else {
		return false
	}
}
