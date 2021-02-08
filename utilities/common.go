package utilities

//用于与前端通信的响应统一结构
type Resp struct {
	Code int					`json:"code"`		//响应代号
	Message string				`json:"message"`	//响应信息
	Data map[string]interface{}	`json:"data"`		//返回数据
}

//通用信息
type Common struct {
	Date string				`json:"date"`			//日期
	Likes int64				`json:"likes"`			//点赞数
	Coins int64				`json:"coins"`			//投币数
	Collections int64		`json:"collections"`	//收藏数
	Shares int64			`json:"shares"`			//分享数
	CommentNumbers int64	`json:"comment_number"`	//评论总数
}