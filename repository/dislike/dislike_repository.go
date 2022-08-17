package dislike_repository

import "ocra_server/model/entity"

type DislikeRepository interface {
	// to check if user already dislike this video or not
	IsUserAlreadyDislikeThisVideo(req *entity.Dislikes) bool
	CreateDislike(req *entity.Dislikes) error
	DeleteDislike(req *entity.Dislikes) error
}
