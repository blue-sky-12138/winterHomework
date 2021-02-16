package serve

import (
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
	"path"
)

//获取视频文件及封面
func GetVideoFile(context *gin.Context) {
	var tem utilities.GetVideoFile
	context.ShouldBindUri(&tem)
	if tem.BvCode == "" && tem.FileName ==  ""{
		var resp utilities.Resp
		resp.Code = 401
		resp.Message = "路径为空！"
		context.JSON(200,resp)
	}else{
		filePath := path.Join("./static","videos",tem.BvCode,tem.FileName)
		context.File(filePath)
	}
}

//获取用户头像
func GetUserHead(context *gin.Context) {
	var tem utilities.GetUserHead
	context.ShouldBindUri(&tem)
	if tem.Id == "" && tem.FileName ==  ""{
		var resp utilities.Resp
		resp.Code = 401
		resp.Message = "路径为空！"
		context.JSON(200,resp)
	}else{
		headPath := path.Join("./static","users",tem.Id,tem.FileName)
		context.File(headPath)
	}
}