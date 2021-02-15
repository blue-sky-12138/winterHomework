package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取视频评论
func GetVideoComments(context *gin.Context) {
	var(
		topComments []utilities.MetaComment			//置顶评论
		hotComments []utilities.MetaComment			//热门评论
		commonComments []utilities.MetaComment		//平平无奇的评论
	)
	bvCode := context.Query("bv_code")
	topComments = *model.GetTopVideoComments(bvCode)
	normalComments := *model.GetHotVideoComments(bvCode)	//获取非置顶评论
	for _, value := range normalComments{			//遍历分离热门评论(这里热门的判断条件为点赞数>=10)
		if value.Likes >= 10{
			hotComments = append(hotComments, value)
		}else {
			commonComments = append(commonComments, value)
		}
	}

	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	resp.Code = 500
	resp.Message ="成功接受请求"
	resp.Data["top_comment"] = topComments
	resp.Data["hot_comments"] = hotComments
	resp.Data["common_comments"] = commonComments
	context.JSON(200,resp)
}

//获取视频信息
func GetVideoInformation(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	videoType := context.Query("type")
	if videoType == "1" {			//视频的简要信息
		data := model.GetBriefVideoInformation("",0,0)
		resp.Code = 600
		resp.Message = "成功接受请求"
		resp.Data["data"] = data
		context.JSON(200,resp)
	}else if videoType == "2" {		//视频的详细信息
		bvCode := context.Query("bv_code")
		data := model.GetDetailedVideoInformation(bvCode)
		resp.Code = 600
		resp.Message = "成功接受请求"
		resp.Data["data"] = data
		context.JSON(200,resp)
	}else {
		resp.Code = 601
		resp.Message = "视频信息的类型不合法"
		context.JSON(200,resp)
	}
}

//获取视频弹幕
func GetVideoBarrages(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	bvCode := context.Query("bv_code")
	p,_ := strconv.Atoi(context.Query("p"))

	data := *model.GetVideoBarrages(bvCode,p)

	resp.Code = 700
	resp.Message = "成功接受请求"
	resp.Data["data"] = data
	context.JSON(200,resp)
}

//获取视频地址
func GetVideoPath(context *gin.Context) {
	var resp utilities.Resp
	resp.Data = make(map[string]interface{})	//防止map为空

	bvCode := context.Query("bv_code")

	data := *model.GetVideoPath(bvCode)

	resp.Code = 800
	resp.Message = "成功接受请求"
	resp.Data["data"] = data
	context.JSON(200,resp)
}