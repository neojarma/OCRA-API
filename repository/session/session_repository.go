package repository

import "ocra_server/model/entity"

type SessionRepository interface {
	CheckActiveSession(currentDateTime uint64, session *entity.Sessions) error
	CreateNewSession(session *entity.Sessions) error
	DeleteSession(session *entity.Sessions) error
}
