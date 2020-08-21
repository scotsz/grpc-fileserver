package server

import "github.com/google/uuid"

type NameGenerator interface {
	Generate() string
}

type uuidNameGen struct {
}

func (u *uuidNameGen) Generate() string {
	return uuid.New().String()
}
func NewUuidGenerator() *uuidNameGen {
	return &uuidNameGen{}
}
