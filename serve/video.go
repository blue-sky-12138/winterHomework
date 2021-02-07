package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
	"path"
)

//获取视频评论
func GetVideoComments(context *gin.Context) {
	bvCode:=context.Query("bv_code")
	videoId:=model.GetVideoId(bvCode)
	topComment:=model.GetHotVideoComments(videoId)
	hotComments:=model.GetHotVideoComments(videoId)
	commonComments:=model.GetCommonVideoComments(videoId)
	var resp utilities.Resp
	resp.Code=500
	resp.Message="成功接受请求"
	resp.Data["top_comment"]=topComment
	resp.Data["hot_comments"]=hotComments
	resp.Data["common_comments"]=commonComments
	context.JSON(200,resp)
}

//获取视频信息
func GetVideoInformation(context *gin.Context) {
	//tem:=model.GetBriefVideoInformation()
}

//获取视频文件
func GetVideoFile(context *gin.Context) {
	var tem utilities.GetVideoFile
	context.ShouldBindUri(&tem)
	if tem.BvCode == "" && tem.FileName ==  ""{
		var resp utilities.Resp
		resp.Code = 401
		resp.Message = "路径为空！"
		context.JSON(200,resp)
	}else{
		filePath:=path.Join("./static",tem.BvCode,tem.FileName)
		context.File(filePath)
	}
}