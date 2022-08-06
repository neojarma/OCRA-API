package user_service

import (
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
	"ocra_server/model/response"
)

type UserService interface {
	ValidateLogin(reqUser *request.AuthRequest) (*joins_model.UserChannelJoin, error)
	Register(reqUser *request.UserRequest) (*response.UserResponse, error)
	Logout(sessionId string) error
	GetDetailUser(reqUser *request.UserRequest) (*joins_model.UserChannelJoin, error)
	UpdateUser(reqUser *request.UserRequest) (*response.UserResponse, error)
}
