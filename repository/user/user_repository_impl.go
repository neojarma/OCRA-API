package user_repository

import (
	"errors"
	"ocra_server/helper"
	"ocra_server/model/entity"
	"sync"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	var doOnce sync.Once
	result := new(userRepositoryImpl)

	doOnce.Do(func() {
		result = &userRepositoryImpl{
			Db: db,
		}
	})

	return result
}

func (repository *userRepositoryImpl) ValidateCredentials(user *entity.Users) error {
	prevPass := user.Password
	result := repository.Db.Select("password").Where("email = ?", user.Email).First(user)
	if result.Error != nil {
		return result.Error
	}

	if ok := helper.CompareHashingPassword(user.Password, prevPass); !ok {
		return errors.New("invalid")
	}

	return nil
}

func (repository *userRepositoryImpl) isEmailAlreadyExist(user *entity.Users) bool {
	result := repository.Db.Where("email = ?", user.Email).First(user)
	return result.RowsAffected == 1
}

func (repository *userRepositoryImpl) CreateUser(user *entity.Users) error {

	if exist := repository.isEmailAlreadyExist(user); exist {
		return errors.New("email already exist")
	}

	if err := repository.Db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (repository *userRepositoryImpl) DetailUser(user *entity.Users) (*entity.Users, error) {
	result := repository.Db.First(user, user.UserId)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return user, nil
}

func (repository *userRepositoryImpl) UpdateUser(user *entity.Users) error {
	err := repository.Db.Model(user).Where("user_id = ?", user.UserId).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}
