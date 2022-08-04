package user_repository

import (
	"errors"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
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
	result := repository.Db.Select("password", "is_verified").Where("email = ?", user.Email).First(user)
	if result.Error != nil {
		return result.Error
	}

	if ok := helper.CompareHashingPassword(user.Password, prevPass); !ok {
		return errors.New(response.MessageWrongCredentials)
	}

	if !user.IsVerified {
		return errors.New(response.MessageNotVerifedUser)
	}

	return nil
}

func (repository *userRepositoryImpl) isEmailAlreadyExist(user *entity.Users) bool {
	result := repository.Db.Where("email = ?", user.Email).First(user)
	return result.RowsAffected == 1
}

func (repository *userRepositoryImpl) UpdateUserEmailStatus(user *entity.Users) error {
	ok := repository.Db.Model(user).Where("email = ?", user.Email).Update("is_verified", true).RowsAffected >= 1
	if !ok {
		return errors.New("no record with this email")
	}

	return nil
}

func (repository *userRepositoryImpl) CreateUser(user *entity.Users) error {

	if exist := repository.isEmailAlreadyExist(user); exist {
		return errors.New(response.MessageFailedRegisterEmailExist)
	}

	if err := repository.Db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (repository *userRepositoryImpl) DetailUser(user *entity.Users) (*joins_model.UserChannelJoin, error) {
	result := new(joins_model.UserChannelJoin)
	err := repository.Db.Model(user).Select("users.user_id", "users.full_name", "users.profile_image", "users.email", "users.created_at").Joins("LEFT JOIN channels on users.user_id = channels.user_id").Where("users.email = ?", user.Email).Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository *userRepositoryImpl) UpdateUser(user *entity.Users) error {
	err := repository.Db.Model(user).Where("user_id = ?", user.UserId).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}
