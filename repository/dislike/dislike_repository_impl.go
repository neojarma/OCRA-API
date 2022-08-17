package dislike_repository

import (
	"ocra_server/model/entity"
	"sync"

	"gorm.io/gorm"
)

type DislikeRepositoryImpl struct {
	Db *gorm.DB
}

func NewDislikeRepository(db *gorm.DB) DislikeRepository {
	var doOnce sync.Once
	repo := new(DislikeRepositoryImpl)

	doOnce.Do(func() {
		repo = &DislikeRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

// to check if user already dislike this video or not
func (repository *DislikeRepositoryImpl) IsUserAlreadyDislikeThisVideo(req *entity.Dislikes) bool {
	newModel := new(entity.Dislikes)
	return repository.Db.Where("video_id = ?", req.VideoId).Where("user_id = ?", req.UserId).Find(newModel).RowsAffected == 1
}

func (repository *DislikeRepositoryImpl) CreateDislike(req *entity.Dislikes) error {
	return repository.Db.Create(req).Error
}

func (repository *DislikeRepositoryImpl) DeleteDislike(req *entity.Dislikes) error {
	return repository.Db.Where("video_id = ?", req.VideoId).Where("user_id = ?", req.UserId).Delete(req).Error
}
