package mocks

import "golabs/lets-go/internal/models"

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "duplicate@gmail.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "yakushou730@mock.com" && password == "pa$$w0rd" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}
