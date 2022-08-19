package helper

import "fmt"

func GetThumbnailFilePath(channelId, videoId string) string {
	return fmt.Sprintf("data/channels/%s/videos/%s/thumbnail/", channelId, videoId)
}

func GetVideoFilePath(channelId, videoId string) string {
	return fmt.Sprintf("data/channels/%s/videos/%s/video/", channelId, videoId)
}

func GetChannelProfileImagePath(channelId string) string {
	return fmt.Sprintf("data/channels/%s/profile/", channelId)
}

func GetUserProfileImagePath(userId string) string {
	return fmt.Sprintf("data/users/%s/profile/", userId)
}
