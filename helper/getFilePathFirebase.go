package helper

import "fmt"

func GetThumbnailFilePath(channelId, videoId string) string {
	return fmt.Sprintf("data/channels/%s/videos/%s/thumbnail/", channelId, videoId)
}

func GetVideoFilePath(channelId, videoId string) string {
	return fmt.Sprintf("data/channels/%s/videos/%s/video/", channelId, videoId)
}
