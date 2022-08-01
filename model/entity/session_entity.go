package entity

type Sessions struct {
	SessionId string
	UserId    string
	CreatedAt uint64
	ExpiresAt uint64
}
