package utilities

//获取视频文件及封面
type GetVideoFile struct {
	BvCode string 		`uri:"bvCode"`		//视频bv号
	FileName string 	`uri:"fileName"`	//文件名
}

//获取用户头像
type GetUserHead struct {
	Id string			`uri:"id"`			//用户id
	FileName string		`uri:"fileName"`	//文件名
}