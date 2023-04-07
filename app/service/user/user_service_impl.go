package user_service

import (
	"context"
	"mime/multipart"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
	"ocra_server/model/response"
	user_repository "ocra_server/repository/user"
	firebase_service "ocra_server/service/firebase"
	session_service "ocra_server/service/session"
	verification_service "ocra_server/service/verification"
	"sync"
	"time"
)

type UserServiceImpl struct {
	UserRepo        user_repository.UserRepository
	VerifService    verification_service.VerificationService
	SessionService  session_service.SessionService
	FirebaseService firebase_service.FirebaseService
}

func NewUserService(
	userRepo user_repository.UserRepository,
	verifService verification_service.VerificationService,
	sessionService session_service.SessionService,
	firebaseService firebase_service.FirebaseService) UserService {
	var doOnce sync.Once
	service := new(UserServiceImpl)

	doOnce.Do(func() {
		service = &UserServiceImpl{
			UserRepo:        userRepo,
			VerifService:    verifService,
			FirebaseService: firebaseService,
			SessionService:  sessionService,
		}
	})

	return service
}

func (service *UserServiceImpl) Logout(sessionId string) error {
	return service.SessionService.DeleteSession(sessionId)
}

func (service *UserServiceImpl) ValidateLogin(reqUser *request.AuthRequest) (*joins_model.UserChannelJoin, error) {
	domainReq := &entity.Users{
		Email:    reqUser.Email,
		Password: reqUser.Password,
	}

	if err := service.UserRepo.ValidateCredentials(domainReq); err != nil {
		return nil, err
	}

	detailReq := &request.UserRequest{
		Email: reqUser.Email,
	}

	result, err := service.GetDetailUser(detailReq)
	if err != nil {
		return nil, err
	}

	sessionId, err := service.SessionService.CreateNewSession(result.UserId)
	if err != nil {
		return nil, err
	}

	result.SessionId = &sessionId

	return result, nil
}

func (service *UserServiceImpl) Register(reqUser *request.UserRequest) (*response.UserResponse, error) {

	newUserId := helper.GetRandomString(22)
	timeCreated := time.Now().UnixMilli()
	hashedPassword, err := helper.GetHashingPassword(reqUser.Password)
	if err != nil {
		return nil, err
	}

	domainUserReq := &entity.Users{
		UserId:    newUserId,
		FullName:  reqUser.FullName,
		Email:     reqUser.Email,
		Password:  hashedPassword,
		CreatedAt: timeCreated,
		UpdatedAt: timeCreated,
	}

	if err := service.UserRepo.CreateUser(domainUserReq); err != nil {
		return nil, err
	}

	if err := service.VerifService.CreateVerifToken(reqUser.Email); err != nil {
		return nil, err
	}

	return &response.UserResponse{
		CreatedAt:        timeCreated,
		UserId:           newUserId,
		FullName:         reqUser.FullName,
		UserProfileImage: nil,
		Email:            reqUser.Email,
	}, nil
}

func (service *UserServiceImpl) VerifyEmail(email, token string) error {

	if err := service.VerifService.ValidateVerifToken(email, token); err != nil {
		return err
	}

	domainUserReq := &entity.Users{
		Email: email,
	}
	if err := service.UserRepo.UpdateUserEmailStatus(domainUserReq); err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) GetDetailUser(reqUser *request.UserRequest) (*joins_model.UserChannelJoin, error) {

	domainUserReq := &entity.Users{
		UserId: reqUser.UserId,
		Email:  reqUser.Email,
	}
	result, err := service.UserRepo.DetailUser(domainUserReq)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) UpdateUser(reqUser *request.UserRequest, profileImage *multipart.FileHeader) (*response.UserResponse, error) {

	domainUserReq := new(entity.Users)
	domainUserReq.UserId = reqUser.UserId
	domainUserReq.FullName = reqUser.FullName

	if profileImage != nil {
		path := helper.GetUserProfileImagePath(reqUser.UserId)
		resourcePath, err := service.FirebaseService.CreateResource(context.Background(), path, profileImage)
		if err != nil {
			return nil, err
		}

		domainUserReq.ProfileImage = &resourcePath
	}

	if err := service.UserRepo.UpdateUser(domainUserReq); err != nil {
		return nil, err
	}

	return &response.UserResponse{
		UserId:           reqUser.UserId,
		FullName:         reqUser.FullName,
		UserProfileImage: reqUser.ProfileImage,
		Email:            reqUser.Email,
	}, nil
}
