package entity

type Users struct {
	UserId       string
	FullName     string
	ProfileImage *string
	Email        string
	Password     string
	IsVerified   bool
	CreatedAt    int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt    int64 `gorm:"autoUpdateTime:milli"`
}
