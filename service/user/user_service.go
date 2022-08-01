package user_service

import (
	"ocra_server/model/request"
	"ocra_server/model/response"
)

type UserService interface {
	ValidateLogin(reqUser *request.UserRequest) (*response.UserResponse, error)
	Register(reqUser *request.UserRequest) (*response.UserResponse, error)
	GetDetailUser(reqUser *request.UserRequest) (*response.UserResponse, error)
	UpdateUser(reqUser *request.UserRequest) (*response.UserResponse, error)
}
