package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
)

//获取视频简要信息
func GetBriefVideoInformation(target string, start int, end int) []utilities.VideoInformation {

	return nil
}

//获取视频详情
func GetDetailedVideoInformation(bvCode string) *utilities.VideoInformation {
	return database.DetailedVideoInformation(bvCode)
}

//获取置顶评论
func GetTopVideoComments(bvCode string) *[]utilities.MetaComment {
	return database.VideoComments(bvCode," and top = 1 ")
}

//获取非置顶评论
func GetHotVideoComments(bvCode string) *[]utilities.MetaComment {
	return database.VideoComments(bvCode," and top = 0 ")
}