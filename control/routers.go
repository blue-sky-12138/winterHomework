package control

import (
	"WinterHomework/serve"
	"WinterHomework/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func RoutersEntrance(){
	//同时输出到终端和日志文件
	file,_:=os.Create("ginLog.md")
	gin.DefaultWriter=io.MultiWriter(file,os.Stdout)

	//http://121.196.155.183:8000/serve
	router:=gin.Default()
	router.Static("/static","./static/")		//加载静态文件夹
	router.Use(middleware.Cors())								//跨域中间件

	user:=router.Group("")							//用户服务
	{
		user.POST("/serve/user/login", serve.PostLogin)		//登录
		user.POST("/serve/user/register", serve.Register)	//注册
		//user.PUT("/serve/user/update", serve.Update)			//更新个人信息
	}

	video:=router.Group("")							//视频服务
	{
		video.GET("/serve/video/comment", serve.GetVideoComments)	//获取视频评论
		video.GET("/serve/video", serve.GetVideoInformation)			//获取视频的元数据
	}

	uploadAndDownload:=router.Group("")
	{
		uploadAndDownload.GET("/serve/video/file/:bvCode/:fileName",serve.GetVideoFile)	//获取视频文件本体
	}

	router.Run(":8000")
}