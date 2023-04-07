package entity

type Sessions struct {
	SessionId string
	UserId    string
	CreatedAt int64
	ExpiresAt int64
}
