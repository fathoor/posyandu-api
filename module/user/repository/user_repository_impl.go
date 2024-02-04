package repository

import (
	"github.com/fathoor/posyandu-api/module/user/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	*gorm.DB
}

func (repository *userRepositoryImpl) Insert(user *entity.User) error {
	return repository.DB.Create(&user).Error
}

func (repository *userRepositoryImpl) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := repository.DB.Find(&users).Error

	return users, err
}

func (repository *userRepositoryImpl) FindByRole(role string) ([]entity.User, error) {
	var users []entity.User
	err := repository.DB.Find(&users, "role = ?", role).Error

	return users, err
}

func (repository *userRepositoryImpl) FindByID(id int) (entity.User, error) {
	var user entity.User
	err := repository.DB.Take(&user, id).Error

	return user, err
}

func (repository *userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := repository.DB.Take(&user, "username = ?", username).Error

	return user, err
}

func (repository *userRepositoryImpl) Save(user *entity.User) error {
	return repository.DB.Save(&user).Error
}

func (repository *userRepositoryImpl) Delete(user *entity.User) error {
	return repository.DB.Delete(&user).Error
}

func ProvideUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}
