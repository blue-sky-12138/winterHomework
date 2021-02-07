package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
)

//获取视频id
func GetVideoId(bvCode string) int64 {
	return database.VideoId(bvCode)
}

//func GetBriefVideoInformation(target string, start int, end int) []utilities.VideoInformation {
//
//}

//获取置顶评论
func GetTopVideoComments(videoId int64) []utilities.MetaComment {
	return database.VideoComments(videoId,"and where top = 1")
}

//获取热门评论
func GetHotVideoComments(videoId int64) []utilities.MetaComment {
	return database.VideoComments(videoId,"and likes > 10")
}

//获取平平无奇的评论
func GetCommonVideoComments(videoId int64) []utilities.MetaComment {
	return database.VideoComments(videoId,"and where top = 0 and where likes <=10")
}