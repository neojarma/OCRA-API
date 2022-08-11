package session_service

import (
	"errors"
	"ocra_server/helper"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	session_repository "ocra_server/repository/session"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

const renewTimeInEpoch = 86400 // 1 day in epoch

type SessionServiceImpl struct {
	Repo         session_repository.SessionRepository
	SessionCache *cache.Cache
}

func NewSessionService(repository session_repository.SessionRepository, sessionCache *cache.Cache) SessionService {
	var doOnce sync.Once
	repo := new(SessionServiceImpl)

	doOnce.Do(func() {
		repo = &SessionServiceImpl{
			Repo:         repository,
			SessionCache: sessionCache,
		}
	})

	return repo
}

func (service *SessionServiceImpl) CheckActiveSession(sessionId string) (*entity.Sessions, error) {
	if userId, expired, found := service.SessionCache.GetWithExpiration(sessionId); found {
		return &entity.Sessions{
			SessionId: sessionId,
			UserId:    userId.(string),
			ExpiresAt: expired.UnixMilli(),
		}, nil
	}

	session, err := service.Repo.CheckActiveSession(&entity.Sessions{SessionId: sessionId})
	if err != nil {
		return nil, err
	}

	service.SessionCache.Set(sessionId, session.UserId, cache.DefaultExpiration)

	return session, nil
}

func (service *SessionServiceImpl) CreateNewSession(userId string) (string, error) {
	newSessionId := helper.GetRandomString(32)
	timeCreated := time.Now().UnixMilli()
	timeExpires := time.Now().Add(72 * time.Hour).UnixMilli()

	req := &entity.Sessions{
		SessionId: newSessionId,
		UserId:    userId,
		CreatedAt: timeCreated,
		ExpiresAt: timeExpires,
	}

	if err := service.Repo.CreateNewSession(req); err != nil {
		return "", err
	}

	service.SessionCache.Set(newSessionId, userId, cache.DefaultExpiration)

	return newSessionId, nil
}

func (service *SessionServiceImpl) UpdateExpiresSession(sessionId string) error {

	if err := service.Repo.UpdateExpiresSession(renewTimeInEpoch, &entity.Sessions{SessionId: sessionId}); err != nil {
		return err
	}

	if userId, found := service.SessionCache.Get(sessionId); found {
		service.SessionCache.Replace(sessionId, userId, (24 * time.Hour))
		return nil
	}

	return errors.New(response.MessageInvalidSession)
}

func (service *SessionServiceImpl) DeleteSession(sessionId string) error {
	go service.SessionCache.Delete(sessionId)

	if err := service.Repo.DeleteSession(&entity.Sessions{SessionId: sessionId}); err != nil {
		return err
	}

	return nil
}
