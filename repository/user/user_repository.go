package user_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type UserRepository interface {
	ValidateCredentials(user *entity.Users) error
	UpdateUserEmailStatus(user *entity.Users) error
	CreateUser(user *entity.Users) error
	DetailUser(user *entity.Users) (*joins_model.UserChannelJoin, error)
	UpdateUser(user *entity.Users) error
}
