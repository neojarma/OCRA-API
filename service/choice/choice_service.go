package choice_service

import "ocra_server/model/entity"

type ChoiceService interface {
	CreateLikeRecord(req *entity.Likes) error
	IsUserLikeTheVideo(req *entity.Likes) bool
	CreateDislikeRecord(req *entity.Dislikes) error
	IsUserDislikeTheVideo(req *entity.Dislikes) bool
}
