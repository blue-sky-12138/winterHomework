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


//视频作者的详细信息结构
type DetailedVideoAuthorInformation struct {
	Id int64				`json:"id"`			//用户id
	Name string				`json:"name"`		//用户名
	Signature string		`json:"signature"`	//个性签名
	Vip int					`json:"vip"`		//是否是大大大会员
	Level int				`json:"level"`		//几级号
	HeadPath string			`json:"head_path"`	//头像地址
	Position string			`json:"position"`	//在作品中的职位，用于联合投稿的情况
}
//视频的详细信息结构
type DetailedVideoInformation struct {
	BvCode string								`json:"bv_code"`		//bv号
	CoverPath string							`json:"cover_path"`		//封面文件地址
	Title string 								`json:"title"`			//标题
	Brief string								`json:"brief"`			//简介
	Plays int64									`json:"plays"`			//播放量
	P int										`json:"p"`				//该视频的分集数
	Author	[]DetailedVideoAuthorInformation	`json:"author"`			//作者信息
	Common										`json:"common"`			//通用信息
}


//视频作者的简略信息结构
type BriefVideoAuthorInformation struct {
	Id int64				`json:"id"`			//用户id
	Name string				`json:"name"`		//用户名
}
//视频简略信息的结构体
type BriefVideoInformation struct {
	BvCode string							`json:"bv_code"`		//bv号
	CoverPath string						`json:"cover_path"`		//封面文件地址
	Title string 							`json:"title"`			//标题
	Date string								`json:"date"`
	Plays int64								`json:"plays"`			//播放量
	Author BriefVideoAuthorInformation		`json:"author"`			//作者信息
}


//视频路径信息结构体
type VideoPathInformation struct {
	P int 		`json:"p"`		//第几分p
	Path string	`json:"path"`	//视频地址
	Name string `json:"name"`	//视频名
}


//弹幕结构体
type VideoBarrage struct {
	Id int64				`json:"id"`			//弹幕id
	DateTime string 		`json:"date_time"`	//弹幕发表日期
	VideoTime string 		`json:"video_time"`	//弹幕在视频中出现的时间点
	UsersId int64			`json:"users_id"`
	Content string 			`json:"content"`	//弹幕内容
	Type int 				`json:"type"`		//弹幕类型
	Size int				`json:"size"`		//弹幕字体大小
	Pattern int				`json:"pattern"`	//弹幕飘出表现形式
	Color int				`json:"color"`		//弹幕颜色
}