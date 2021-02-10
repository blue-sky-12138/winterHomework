package utilities

//评论用户的基本信息结构
type CommentsAuthorInformation struct {
	Id int64				`json:"id"`		//用户id
	Name string				`json:"name"`	//用户名
	Vip int					`json:"vip"`	//是否是大大大会员
	Level int				`json:"level"`	//几级号
}

//楼中楼评论结构
type ReplyComment struct {
	Id int64								`json:"id"`				//评论的id
	Author	CommentsAuthorInformation		`json:"author"`			//评论作者信息
	Content string							`json:"content"`		//评论内容
	Date string								`json:"date"`
	Likes int64								`json:"likes"`
	ReplyAuthor CommentsAuthorInformation	`json:"reply_author"`	//被回复人的信息
}

//元评论结构构
type MetaComment struct {
	Id int64							`json:"id"`					//评论的id
	Author	CommentsAuthorInformation	`json:"author"`
	Content string						`json:"content"`
	Floor int							`json:"floor"`				//几楼
	Date string							`json:"date"`
	Likes int64							`json:"likes"`
	ReplyComments []ReplyComment		`json:"comments_in_floor"`	//该楼的楼中楼评论
}

//视频作者信息结构
type VideoAuthorInformation struct {
	Id int64				`json:"id"`			//用户id
	Name string				`json:"name"`		//用户名
	Signature string		`json:"signature"`	//个性签名
	Vip int					`json:"vip"`		//是否是大大大会员
	Level int				`json:"level"`		//几级号
	HeadPath string			`json:"head_path"`	//头像地址
	Position string			`json:"position"`	//在作品中的职位，用于联合投稿的情况
}

//视频信息结构
type VideoInformation struct {
	Id int64							`json:"id"`				//视频id
	BvCode string						`json:"bv_code"`		//bv号
	VideoPath string					`json:"video_path"`		//视频文件地址
	CoverPath string					`json:"cover_path"`		//封面文件地址
	Title string 						`json:"title"`			//标题
	Brief string						`json:"brief"`			//简介
	Plays int64							`json:"plays"`			//播放量
	Author	[]VideoAuthorInformation	`json:"author"`			//作者信息
	Common								`json:"common"`			//通用信息
}
