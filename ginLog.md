[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
[GIN-debug] HEAD   /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
[GIN-debug] POST   /serve/user/login         --> WinterHomework/serve.Login (4 handlers)
[GIN-debug] POST   /serve/user/register      --> WinterHomework/serve.Register (4 handlers)
[GIN-debug] GET    /serve/video/comment      --> WinterHomework/serve.GetVideoComments (4 handlers)
[GIN-debug] GET    /serve/video/information  --> WinterHomework/serve.GetVideoInformation (4 handlers)
[GIN-debug] GET    /serve/video/barrage      --> WinterHomework/serve.GetVideoBarrages (4 handlers)
[GIN-debug] GET    /serve/video/path         --> WinterHomework/serve.GetVideoPath (4 handlers)
[GIN-debug] GET    /serve/download/user/head/:id/:fileName --> WinterHomework/serve.GetUserHead (4 handlers)
[GIN-debug] GET    /serve/download/video/cover/:bvCode/:fileName --> WinterHomework/serve.GetVideoFile (4 handlers)
[GIN-debug] GET    /serve/download/video/file/:bvCode/:fileName --> WinterHomework/serve.GetVideoFile (4 handlers)
[GIN-debug] PUT    /serve/upload/user/head   --> WinterHomework/middleware.Cors.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8000
[GIN] 2021/02/23 - 09:02:21 | 200 |    1.0696271s |             ::1 | GET      "/serve/video/information?bv_code=BV1No4y1d7tA&type=2"
