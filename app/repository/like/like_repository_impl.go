package like_repository

import (
	"ocra_server/model/entity"
	"sync"

	"gorm.io/gorm"
)

type LikeRepositoryImpl struct {
	Db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	var doOnce sync.Once
	repo := new(LikeRepositoryImpl)

	doOnce.Do(func() {
		repo = &LikeRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

// to check if user already like this video or not
func (repository *LikeRepositoryImpl) IsUserAlreadyLikeThisVideo(req *entity.Likes) bool {
	newModel := new(entity.Likes)
	return repository.Db.Where("video_id = ?", req.VideoId).Where("user_id = ?", req.UserId).Find(newModel).RowsAffected == 1
}

func (repository *LikeRepositoryImpl) CreateLike(req *entity.Likes) error {
	return repository.Db.Create(req).Error
}

func (repository *LikeRepositoryImpl) DeleteLike(req *entity.Likes) error {
	return repository.Db.Where("video_id = ?", req.VideoId).Where("user_id = ?", req.UserId).Delete(req).Error
}
