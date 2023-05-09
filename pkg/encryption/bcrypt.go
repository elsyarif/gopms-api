package encryption

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	Hash(password string) string
	ComparePassword(password, hash string) error
}

type password struct{}

func PasswordHash() Password {
	return &password{}
}

func (p *password) Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (p *password) ComparePassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
