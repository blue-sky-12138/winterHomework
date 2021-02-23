服务器地址：121.196.155.183:8000

本项目架构
MSC架构(应该)

Utilities包
里面主要包含了各种自定义的结构体，其中md5Cryptography包封装了md5快速加密，standatdCheck包封装了对部分数据的合法性判断，logError包封装了简单的日志输出

middleware包
用于存储中间件，包括cors跨域中间件

static文件夹
里面用于存储所所有的静态文件，包括用户头像，视频文件，视频封面等
