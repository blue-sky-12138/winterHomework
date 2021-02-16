package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
	"fmt"
)

//获取视频简要信息。
//count表示一共要获取几条数据，最多获取1000条，若为0则默认获取1000条。
//target为搜索关键词。
func GetBriefVideoInformation(target string, count int) (*[]utilities.BriefVideoInformation, error) {
	if count == 0 {
		data , err := database.BriefVideoInformation(target,"limit 0,1000")
		if err != nil {
			return nil, err
		}else {
			return data, nil
		}
	}else if count > 0 && count <= 1000{
		limit := fmt.Sprintf("limit 0,%d",count)
		data , err := database.BriefVideoInformation(target,limit)
		if err != nil {
			return nil, err
		}else {
			return data, nil
		}
	}else {					//如果count不在0~1000区间内
		return nil, fmt.Errorf("输入的数值不合法")
	}
}

//获取视频详情。
func GetDetailedVideoInformation(bvCode string) (*utilities.DetailedVideoInformation, error) {
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