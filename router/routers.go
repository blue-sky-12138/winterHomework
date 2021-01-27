package router

import(
	"WinterHomework/controller"
	"WinterHomework/middleware"
	"github.com/gin-gonic/gin"
)

func RoutersEntrance(){
	//http://121.196.155.183:8000/serve
	router:=gin.Default()
	router.Use(middleware.Cors())

	router.POST("/serve/login",controller.PostLogin)

	router.POST("/serve/register",controller.Register)

	router.Run(":8000")
}