package repository

import "github.com/itsLeonB/posyandu-api/module/user/entity"

type UserRepository interface {
	Insert(user *entity.User) error
	FindAll() ([]entity.User, error)
	FindByRole(role string) ([]entity.User, error)
	FindByID(id int) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	Save(user *entity.User) error
	Delete(user *entity.User) error
}
