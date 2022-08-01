package user_service

import (
	"ocra_server/helper"
	"ocra_server/model/entity"
	"ocra_server/model/request"
	"ocra_server/model/response"
	user_repository "ocra_server/repository/user"
	"sync"
	"time"
)

type UserServiceImpl struct {
	UserRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	var doOnce sync.Once
	service := new(UserServiceImpl)

	doOnce.Do(func() {
		service = &UserServiceImpl{
			UserRepo: userRepo,
		}
	})

	return service
}

func (service *UserServiceImpl) ValidateLogin(reqUser *request.UserRequest) (*response.UserResponse, error) {
	domainReq := &entity.Users{
		Email:    reqUser.Email,
		Password: reqUser.Password,
	}
	if err := service.UserRepo.ValidateCredentials(domainReq); err != nil {
		return nil, err
	}

	return &response.UserResponse{
		Email: reqUser.Email,
	}, nil
}

func (service *UserServiceImpl) Register(reqUser *request.UserRequest) (*response.UserResponse, error) {

	newUserId := helper.GetRandomString(24)
	hashedPassword, err := helper.GetHashingPassword(reqUser.Password)
	if err != nil {
		return nil, err
	}

	domainUserReq := &entity.Users{
		UserId:   newUserId,
		FullName: reqUser.FullName,
		Email:    reqUser.Email,
		Password: hashedPassword,
	}

	if err := service.UserRepo.CreateUser(domainUserReq); err != nil {
		return nil, err
	}

	return &response.UserResponse{
		CreatedAt:        time.Now().String(),
		UserId:           newUserId,
		FullName:         reqUser.FullName,
		UserProfileImage: nil,
		Email:            reqUser.Email,
	}, nil
}

func (service *UserServiceImpl) GetDetailUser(reqUser *request.UserRequest) (*response.UserResponse, error) {

	domainUserReq := &entity.Users{
		UserId: reqUser.UserId,
	}
	result, err := service.UserRepo.DetailUser(domainUserReq)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		UserId:           reqUser.UserId,
		FullName:         result.FullName,
		UserProfileImage: &result.UserProfileImage,
		Email:            result.Email,
	}, nil
}

func (service *UserServiceImpl) UpdateUser(reqUser *request.UserRequest) (*response.UserResponse, error) {

	domainUserReq := &entity.Users{
		UserId:           reqUser.UserId,
		FullName:         reqUser.FullName,
		UserProfileImage: reqUser.UserProfileImage,
		Email:            reqUser.Email,
	}
	if err := service.UserRepo.UpdateUser(domainUserReq); err != nil {
		return nil, err
	}

	return &response.UserResponse{
		UserId:           reqUser.UserId,
		FullName:         reqUser.FullName,
		UserProfileImage: &reqUser.UserProfileImage,
		Email:            reqUser.Email,
	}, nil
}
