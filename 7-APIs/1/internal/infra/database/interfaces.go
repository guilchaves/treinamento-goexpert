package database

import "github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
