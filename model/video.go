package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
)

//获取视频简要信息。
func GetBriefVideoInformation(target string, start int, end int) (_ *utilities.VideoInformation, err string) {

	return nil, ""
}

//获取视频详情。
func GetDetailedVideoInformation(bvCode string) (*utilities.VideoInformation, error) {
	data, err := database.DetailedVideoInformation(bvCode)
	if err == nil {			//如果为空
		return data, nil
	}
	return nil, err
}

//获取置顶评论。
func GetTopVideoComments(bvCode string) (*[]utilities.MetaComment, error) {
	videoId, err1 := database.VideosId(bvCode)
	if err1 == nil {			//如果为空
		data, err2 := database.VideoComments(videoId," and top = 1 ")
		if err2 == nil {			//如果为空
			return data, nil
		}else {
			return nil, err2
		}
	}else {
		return nil, err1
	}
}

//获取非置顶评论。
func GetHotVideoComments(bvCode string) (*[]utilities.MetaComment, error) {
	videoId, err1 := database.VideosId(bvCode)
	if err1 == nil {			//如果为空
		data, err2 := database.VideoComments(videoId," and top = 0 ")
		if err2 == nil {			//如果为空
			return data, nil
		}else {
			return nil, err2
		}
	}else {
		return nil, err1
	}
}

//获取视频弹幕。
func GetVideoBarrages(bvCode string,p int) (*[]utilities.VideoBarrage, error) {
	videoId, err1 := database.VideosId(bvCode)
	if err1 == nil {			//如果为空
		data, err2 := database.VideoBarrages(videoId,p)
		if err2 == nil {			//如果为空
			return data, nil
		}else {
			return nil, err2
		}
	}else{
		return nil, err1
	}
}

//获取视频地址及相关信息。
func GetVideoPath(bvCode string) (*[]utilities.VideoPathInformation, error) {
	videoId, err1 := database.VideosId(bvCode)
	if err1 == nil {			//如果为空
		data, err2 := database.VideoPath(videoId)
		if err2 == nil {			//如果为空
			return data,  nil
		}else {
			return nil, err2
		}
	}else {
		return nil,  err1
	}
}