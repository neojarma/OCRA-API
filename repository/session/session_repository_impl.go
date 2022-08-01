package repository

import (
	"ocra_server/model/entity"

	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
	Db *gorm.DB
}

func (repository *SessionRepositoryImpl) CheckActiveSession(currentDateTime uint64, session *entity.Sessions) error {
	err := repository.Db.Where("session_id = ?", session.SessionId).Where("expires_at >= ?", currentDateTime).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}

	return nil
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
