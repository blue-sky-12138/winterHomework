package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
)

//func GetBriefVideoInformation(target string, start int, end int) []utilities.VideoInformation {
//
//}

//获取置顶评论
func GetTopVideoComments(bvCode string) []utilities.MetaComment {
	return database.VideoComments(bvCode,"and where top = 1")
}

//获取非置顶评论
func GetHotVideoComments(bvCode string) []utilities.MetaComment {
	return database.VideoComments(bvCode,"and where top = 0")
}