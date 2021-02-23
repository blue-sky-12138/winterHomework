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
	Likes int								`json:"likes"`
	ReplyAuthor CommentsAuthorInformation	`json:"reply_author"`	//被回复人的信息
}
//元评论结构构
type MetaComment struct {
	Id int64							`json:"id"`					//评论的id
	Author	CommentsAuthorInformation	`json:"author"`
	Content string						`json:"content"`
	Floor int							`json:"floor"`				//几楼
	Date string							`json:"date"`
	Likes int							`json:"likes"`
	Heat int							`json:"heat"`				//评论热度，数值为点赞与踩的数值和
	ReplyComments []ReplyComment		`json:"comments_in_floor"`	//该楼的楼中楼评论
}

//视频新增评论结构
type NewComment struct {
	VideoId int64					`json:"video_id"`			//视频ID
	DateTime string					`json:"date_time"`			//评论时间
	Content	string					`json:"content"`			//评论内容
	AuthorId int64					`json:"user_id"`			//评论者ID
	MetaFloor int					`json:"meta_floor"`			//元评论楼层
	ReplyCommentId int64			`json:"reply_comment_id"`	//被评论id
	ReplyAuthorId int64				`json:"reply_author_id"`	//被评论者id
}

//视频操作信息结构
type OperateVideoInformation struct {
	Id int64			//用于存储已登录的用户ID
	Like int			//点赞情况
	Coin int			//投币情况
	Collect int 		//收藏情况
	Share int			//分享情况
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
//视频参加的活动结构
type DetailedVideoActivity struct {
	Join int		`json:"join"`		//是否参加，默认为0(不参加)，1为参加
	Name string		`json:"name"`		//活动名字
	Url string		`json:"url"`		//活动URL
}
//视频的更多选项
type DetailedVideoMore struct {
	PersonalDeclaration int		`json:"p_declar"`	//自制声明，默认为0(不添加)，1为添加
	WaterMark int				`json:"water_mark"`	//水印类型
	BusincessDeclaration int	`json:"b_declar"`	//商业声明，默认为0(不含)，1为含
	SubtitleLanguage int		`json:"sub_lang"`	//字幕语言类型
	SubtitleOpen int			`json:"sub_open"`	//是否允许粉丝投稿字幕，默认为0(不允许)，1为允许
}
//视频的详细信息结构
type DetailedVideoInformation struct {
	BvCode string								`json:"bv_code"`		//bv号
	CoverPath string							`json:"cover_path"`		//封面文件地址
	Title string 								`json:"title"`			//标题
	Brief string								`json:"brief"`			//简介
	Plays int64									`json:"plays"`			//播放量
	Type int									`json:"type"`			//视频的类型(自制还是转载)
	Tags []string								`json:"tags"`			//视频标签
	Activity DetailedVideoActivity				`json:"activity"`		//视频参加的活动
	P int										`json:"p"`				//该视频的分集数
	Author	[]DetailedVideoAuthorInformation	`json:"author"`			//作者信息
	Common										`json:"common"`			//通用信息
	Operate OperateVideoInformation 			`json:"operate"`		//用户对视频的操作信息(如果处于登录状态)
}


//视频作者的简略信息结构
type BriefVideoAuthorInformation struct {
	Id int64				`json:"id"`			//用户id
	Name string				`json:"name"`		//用户名
}
//视频简略信息的结构
type BriefVideoInformation struct {
	BvCode string							`json:"bv_code"`		//bv号
	CoverPath string						`json:"cover_path"`		//封面文件地址
	Title string 							`json:"title"`			//标题
	Date string								`json:"date"`
	Plays int64								`json:"plays"`			//播放量
	Author BriefVideoAuthorInformation		`json:"author"`			//作者信息
}


//视频路径信息结构
type VideoPathInformation struct {
	P int 		`json:"p"`		//第几分p
	Path string	`json:"path"`	//视频地址
	Name string `json:"name"`	//视频名
}


//弹幕结构
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

//生成bv号
func NewBvCode() string {
	return "BVdfsaoffads"
}