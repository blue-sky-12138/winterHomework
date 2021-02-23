package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取视频评论。
func GetVideoComments(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	var(
		topComments *[]utilities.MetaComment		//置顶评论
		hotComments *[]utilities.MetaComment			//热门评论
		commonComments *[]utilities.MetaComment		//平平无奇的评论
	)
	bvCode := context.Query("bv_code")
	topComments, err := model.GetTopVideoComments(bvCode)		//获取置顶评论
	if err != nil {
		resp.Code = 50001
		resp.Message = err.Error()
		context.JSON(200,resp)
		return
	}

	hotComments, commonComments, err = model.GetHotVideoComments(bvCode)	//获取非置顶评论
	if err != nil {
		resp.Code = 50001
		resp.Message = err.Error()
		context.JSON(200,resp)
		return
	}

	resp.Code = 500
	resp.Message ="响应成功"
	resp.Data["top_comment"] = topComments
	resp.Data["hot_comments"] = hotComments
	resp.Data["common_comments"] = commonComments
	context.JSON(200,resp)
}

//获取视频信息。
func GetVideoInformation(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	videoType := context.Query("type")

	if videoType == "1" {			//视频的简要信息
		target := context.Query("target")
		count,_ := strconv.Atoi(context.Query("count"))

		data, err := model.GetBriefVideoInformation(target,count)
		if err == nil  {
			resp.Code = 600
			resp.Message = "响应成功"
			resp.Data["data"] = data
			context.JSON(200,resp)
		}else {
			resp.Code = 60002
			resp.Message = err.Error()
			context.JSON(200,resp)
		}
	}else if videoType == "2" {		//视频的详细信息
		bvCode := context.Query("bv_code")
		id, _ := strconv.ParseInt(context.Query("user_id"),10,64)
		data, err := model.GetDetailedVideoInformation(bvCode, id)
		if err == nil  {
			resp.Code = 600
			resp.Message = "响应成功"
			resp.Data["data"] = data
			context.JSON(200,resp)
		}else {
			resp.Code = 60002
			resp.Message = err.Error()
			context.JSON(200,resp)
		}
	}else {					//视频的类型无法识别
		resp.Code = 60001
		resp.Message = "视频信息的类型不合法"
		context.JSON(200,resp)
	}
}

//获取视频弹幕。
func GetVideoBarrages(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	bvCode := context.Query("bv_code")
	p, _ := strconv.Atoi(context.Query("p"))

	data, err := model.GetVideoBarrages(bvCode,p)
	if err == nil  {
		resp.Code = 700
		resp.Message = "响应成功"
		resp.Data["data"] = data
		context.JSON(200,resp)
	}else {
		resp.Code = 70001
		resp.Message = err.Error()
		context.JSON(200,resp)
	}
}

//获取视频地址。
func GetVideoPath(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	bvCode := context.Query("bv_code")

	data, err := model.GetVideoPath(bvCode)
	if err == nil  {
		resp.Code = 800
		resp.Message = "响应成功"
		resp.Data["data"] = data
		context.JSON(200,resp)
	}else {
		resp.Code = 80001
		resp.Message = err.Error()
		context.JSON(200,resp)
	}
}

//进行视频操作。
func OperateVideo(context *gin.Context) {
	var (
		resp utilities.Resp
		err error
	)

	videoId, _ := strconv.ParseInt(context.Query("video_id"),10,64)
	userId, _ := strconv.ParseInt(context.Query("user_id"),10,64)
	operateType := context.Query("type")						//进行了什么操作
	value, _ := strconv.Atoi(context.Query("value"))	//操作后的数值一

	if videoId == 0 || userId == 0 {
		resp.Code = 90001
		resp.Message = "信息不完整"
		context.JSON(200,resp)
		return
	}
	
	if operateType == "1" {				//点赞操作
		if value > 1 || value < -1 {		//如果点赞数值不合法
			resp.Code = 90002
			resp.Message = "点赞数值不合法"
			context.JSON(200,resp)
			return	
		}

		err = model.UpdateVideoOperation(userId, videoId, "likes", value)

	}else if operateType == "2" {			//投币操作
		if value > 2 || value < 1 {		//如果投币数值不合法
			resp.Code = 90003
			resp.Message = "投币数值不合法"
			context.JSON(200,resp)
			return	
		}

		err = model.UpdateVideoOperation(userId, videoId, "coins", value)
		
	}else if operateType == "3" {			//收藏操作
		if value > 1 || value < 0 {		//如果收藏数值不合法
			resp.Code = 90004
			resp.Message = "收藏数值不合法"
			context.JSON(200,resp)
			return	
		}

		err = model.UpdateVideoOperation(userId, videoId, "collections", value)
		
	}else if operateType == "4" {			//分享操作
		if value > 1 || value < 0 {		//如果分享数值不合法
			resp.Code = 90005
			resp.Message = "分享数值不合法"
			context.JSON(200,resp)
			return	
		}

		err = model.UpdateVideoOperation(userId, videoId, "shares", value)
	
	}else if operateType == "6" {			//一键三连操作

		err = model.UpdateVideoOperation(userId, videoId, "likes,coins,collections", 1, 1, 1)
	
	}else {								//操作类型不合法
		resp.Code = 90002
		resp.Message = "操作类型不合法"
		context.JSON(200,resp)
		return
	}

	if err == nil  {
		resp.Code = 900
		resp.Message = "响应成功"
		context.JSON(200,resp)
	}else {
		resp.Code = 90006
		resp.Message = err.Error()
		context.JSON(200,resp)
	}	
}

//添加评论
func AddComment(context *gin.Context) {
	var (
		data utilities.NewComment
		resp utilities.Resp
	)
	err := context.ShouldBindJSON(&data)
	if err != nil {
		utilities.LogError("GetNewComment Error",err)
		resp.Code = 13001
		resp.Message = "未知错误13001"
		context.JSON(200,resp)
		return
	}

	err = model.AddNewComment(&data)
	if err != nil {
		utilities.LogError("AddNewComment Error",err)
		resp.Code = 13002
		resp.Message = "未知错误13002"
		context.JSON(200,resp)
	}else {
		resp.Code = 1300
		resp.Message = "评论成功"
		context.JSON(200,resp)
	}
}