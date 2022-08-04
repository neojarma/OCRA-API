package session_repository

import (
	"errors"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	"sync"
	"time"

	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
	Db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	var doOnce sync.Once
	repo := new(SessionRepositoryImpl)

	doOnce.Do(func() {
		repo = &SessionRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *SessionRepositoryImpl) UpdateExpiresSession(newExpires int64, session *entity.Sessions) error {
	result := repository.Db.Model(session).Where("session_id = ?", session.SessionId).Update("expires_at", gorm.Expr("expires_at + ?", newExpires))
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New(response.MessageInvalidSession)
		}

		return result.Error
	}

	return nil

}

func (repository *SessionRepositoryImpl) CheckActiveSession(session *entity.Sessions) (*entity.Sessions, error) {
	currentDateTime := time.Now().UnixMilli()
	err := repository.Db.Where("session_id = ?", session.SessionId).Where("expires_at >= ?", currentDateTime).Select("user_id", "expires_at").First(session).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New(response.MessageInvalidSession)
	}

	return session, nil
}

func (repository *SessionRepositoryImpl) CreateNewSession(session *entity.Sessions) error {
	if err := repository.Db.Create(session).Error; err != nil {
		return err
	}

	return nil
}

func (repository *SessionRepositoryImpl) DeleteSession(session *entity.Sessions) error {
	if err := repository.Db.Where("session_id = ?", session.SessionId).Delete(session).Error; err != nil {
		return err
	}

	return nil
}
