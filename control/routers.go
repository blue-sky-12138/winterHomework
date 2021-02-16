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
	file,_ := os.Create("ginLog.md")
	gin.DefaultWriter = io.MultiWriter(file,os.Stdout)

	//http://121.196.155.183:8000/serve
	router := gin.Default()
	router.Static("static","./static/")			//加载静态文件夹
	router.StaticFile("favicon.ico", "./static/favicon.ico")	//加载网页图标
	router.Use(middleware.Cors())								//跨域中间件

	user := router.Group("/serve/user")							//用户服务
	{
		user.POST("/login", serve.Login)				//登录
		user.POST("/register", serve.Register)		//注册
		//user.PUT("/serve/user/update", serve.Update)			//更新个人信息
	}

	video := router.Group("/serve/video")					//视频服务
	{
		video.GET("/comment", serve.GetVideoComments)			//获取视频评论
		video.GET("/information", serve.GetVideoInformation)		//获取视频的元数据
		video.GET("/barrage",serve.GetVideoBarrages)				//获取视频弹幕
		video.GET("/path",serve.GetVideoPath)					//获取视频地址
	}

	download := router.Group("/serve/download")
	{
		download.GET("/user/head/:id/:fileName",serve.GetUserHead)			//获取头像
		download.GET("/video/cover/:bvCode/:fileName",serve.GetVideoFile)	//获取视频封面
		download.GET("/video/file/:bvCode/:fileName",serve.GetVideoFile)		//获取视频文件本体
	}

	upload := router.Group("/serve/upload")
	{
		upload.PUT("/user/head")
	}

	router.Run(":8000")
}