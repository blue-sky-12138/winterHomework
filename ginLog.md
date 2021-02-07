[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] POST   /serve/login              --> WinterHomework/serve.PostLogin (4 handlers)
[GIN-debug] POST   /serve/register           --> WinterHomework/serve.Register (4 handlers)
[GIN-debug] GET    /serve/video/comment      --> WinterHomework/serve.GetVideoComments (4 handlers)
[GIN-debug] GET    /serve/video              --> WinterHomework/serve.GetVideoInformation (4 handlers)
[GIN-debug] GET    /serve/video/file/:bvCode/:fileName --> WinterHomework/serve.GetVideoFile (4 handlers)
[GIN-debug] Listening and serving HTTP on :8000
[GIN] 2021/02/07 - 11:45:09 | 404 |            0s |             ::1 | GET      "/static/BV1No4y1d7tA/FLAC.gif"
[GIN] 2021/02/07 - 11:45:32 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:33 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:34 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:34 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:35 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:35 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:36 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:37 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:37 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:38 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:38 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:39 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:40 | 404 |            0s |             ::1 | GET      "/favicon.ico"
[GIN] 2021/02/07 - 11:45:56 | 404 |      1.2053ms |             ::1 | GET      "/serve/video/file/BV1No4y1d7tA/FLAC.gif"
[GIN] 2021/02/07 - 11:46:22 | 200 |    851.4885ms |             ::1 | GET      "/serve/video/file/BV1No4y1d7tA/FLAC.gif"
