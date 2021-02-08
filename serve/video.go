package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
)

//获取视频评论
func GetVideoComments(context *gin.Context) {
	var(
		topComments []utilities.MetaComment			//置顶评论
		hotComments []utilities.MetaComment			//热门评论
		commonComments []utilities.MetaComment		//平平无奇的评论
	)
	bvCode:=context.Query("bv_code")
	topComments = model.GetTopVideoComments(bvCode)
	normalComments:=model.GetHotVideoComments(bvCode)	//获取非置顶评论
	for _, value := range normalComments{			//遍历分离热门评论(这里热门的判断条件为点赞数>=10)
		if value.Likes >= 10{
			hotComments = append(hotComments, value)
		}else {
			commonComments = append(commonComments, value)
		}
	}

	var resp utilities.Resp
	resp.Code=500
	resp.Message="成功接受请求"
	resp.Data["top_comment"] = topComments
	resp.Data["hot_comments"] = hotComments
	resp.Data["common_comments"] = commonComments
	context.JSON(200,resp)
}

//获取视频信息
func GetVideoInformation(context *gin.Context) {
	//tem:=model.GetBriefVideoInformation()
}

