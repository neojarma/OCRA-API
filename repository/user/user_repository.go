package user_repository

import (
	"ocra_server/model/entity"
)

type UserRepository interface {
	ValidateCredentials(user *entity.Users) error
	CreateUser(user *entity.Users) error
	DetailUser(user *entity.Users) (*entity.Users, error)
	UpdateUser(user *entity.Users) error
}
