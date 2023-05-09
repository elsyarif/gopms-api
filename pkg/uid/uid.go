package uid

import "github.com/google/uuid"

type Generator interface {
	Uid() string
}

type generator struct{}

func New() Generator {
	return &generator{}
}

func (g *generator) Uid() string {
	return uuid.New().String()
}

func Parse(value string) (string, error) {
	val, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}
	return val.String(), nil
}
