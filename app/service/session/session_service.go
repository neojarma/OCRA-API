package session_service

import "ocra_server/model/entity"

type SessionService interface {
	CheckActiveSession(sessionId string) (*entity.Sessions, error)
	CreateNewSession(userId string) (string, error)
	UpdateExpiresSession(sessionId string) error
	DeleteSession(sessionId string) error
}
