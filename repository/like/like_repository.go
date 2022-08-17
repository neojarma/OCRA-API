package like_repository

import "ocra_server/model/entity"

type LikeRepository interface {
	// to check if user already like this video or not
	IsUserAlreadyLikeThisVideo(req *entity.Likes) bool
	CreateLike(req *entity.Likes) error
	DeleteLike(req *entity.Likes) error
}
