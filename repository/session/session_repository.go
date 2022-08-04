package session_repository

import "ocra_server/model/entity"

type SessionRepository interface {
	CheckActiveSession(session *entity.Sessions) (*entity.Sessions, error)
	CreateNewSession(session *entity.Sessions) error
	UpdateExpiresSession(newExpires int64, session *entity.Sessions) error
	DeleteSession(session *entity.Sessions) error
}
