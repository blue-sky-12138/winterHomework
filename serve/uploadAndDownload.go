package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

//获取视频文件及封面
func GetVideoFile(context *gin.Context) {
	var (
		resp utilities.Resp
		tem utilities.GetVideoFile
	)
	context.ShouldBindUri(&tem)
	if tem.BvCode == "" && tem.FileName ==  ""{
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
	var (
		resp utilities.Resp
		tem utilities.GetUserHead
	)
	context.ShouldBindUri(&tem)
	if tem.Id == "" && tem.FileName ==  ""{
		resp.Code = 401
		resp.Message = "路径为空！"
		context.JSON(200,resp)
	}else{
		headPath := path.Join("./static","users",tem.Id,tem.FileName)
		context.File(headPath)
	}
}

//更新用户头像
func UpdateUserHead(context *gin.Context) {
	var resp utilities.Resp
	
	userId, _ := strconv.ParseInt(context.PostForm("user_id"),10,64)
	if userId == 0 {
		resp.Code = 11001
		resp.Message = "用户ID不合法"
		context.JSON(200,resp)
	}

	file, header, err := context.Request.FormFile("img")
	if err != nil {
		utilities.LogError("GetFormFile(head) Error",err)

		resp.Code = 11002
		resp.Message = "获取文件失败"
		context.JSON(200,resp)
		return
	}

	//获取文件后缀名
	var (
		fileNameLen = len(header.Filename)				//文件名总长
		suffix string
	)
	for i := fileNameLen; ; i-- {
		if header.Filename[i-1:i] == "." {
			suffix = header.Filename[i-1:fileNameLen]
			break
		}
	}

	//路径添加，以系统Unix时间作为文件名
	newPath := fmt.Sprintf("%d/%d",userId,time.Now().Unix()) + suffix
	headPath :=utilities.HeadPath + newPath

	//保存文件
	out, err := os.Create(headPath)
	if err != nil {
		utilities.LogError("CreateFile(head) Error",err)

		resp.Code = 11003
		resp.Message = "未知错误11003"
		context.JSON(200,resp)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		utilities.LogError("CopyFile(head) Error",err)

		resp.Code = 11003
		resp.Message = "未知错误11003"
		context.JSON(200,resp)
		return
	}

	//删除用户原头像
	previousPath, err := model.GetUserHeadPath(userId)
	if err != nil {				//获取用户原头像路径失败
		resp.Code = 11004
		resp.Message = "未知错误11004"
		context.JSON(200,resp)
		return
	}
	err = os.Remove(utilities.HeadPath + previousPath)
	if err != nil {				//删除用户原头像路失败
		resp.Code = 11005
		resp.Message = "未知错误11005"
		context.JSON(200,resp)
		return
	}

	err = model.ChangeUserHead(userId, newPath)
	if err != nil {				//更新路径失败
		resp.Code = 11006
		resp.Message = "未知错误11006"
		context.JSON(200,resp)
		return
	}
	
	resp.Code = 1100
	resp.Message = "更新成功"
	context.JSON(200,resp)
}

//Attention!未完工状态
//上传单个视频
func UploadVideoOne(context *gin.Context) {
	var (
		resp utilities.Resp
		data utilities.NewVideo
	)

	fileHeader, err := context.FormFile("file")
	if err != nil {
		utilities.LogError("GetVideoFile Error",err)
		resp.Code = 13002
		resp.Message = "未知错误13002"
		context.JSON(200,resp)
		return
	}

	//获取文件后缀名
	var (
		fileNameLen = len(fileHeader.Filename)				//文件名总长
		suffix string
	)
	for i := fileNameLen; ; i-- {
		if fileHeader.Filename[i-1:i] == "." {
			suffix = fileHeader.Filename[i-1:fileNameLen]
			break
		}
	}

	data.BvCode = utilities.NewBvCode()
	videoPath := utilities.VideoPath + fileHeader.Filename + suffix

	err = context.SaveUploadedFile(fileHeader,videoPath)
	if err != nil {
		utilities.LogError("SaveVideoFileError",err)
		resp.Code = 13003
		resp.Message = "未知错误13003"
		context.JSON(200,resp)
		return
	}

	err = model.AddNewVideoOne(&data)
	if err != nil {
		utilities.LogError("AddNewVideoOneInformation Error",err)
		resp.Code = 13004
		resp.Message = "未知错误13004"
		context.JSON(200,resp)
		return
	}

	resp.Code = 1300
	resp.Message = "投稿成功"
	context.JSON(200,resp)
}

//Attention!未完工状态
//上传多个视频
func UploadVideoMore(context *gin.Context) {
	//var resp utilities.Resp

}

//Attention!未完工状态
//仅在本包使用。
//保存文件的快捷方式。
//func saveFile(context *gin.Context) error {
//
//}