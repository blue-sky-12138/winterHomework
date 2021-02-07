package utilities

type Author struct {
	Id int64				`json:"id"`
	Name string				`json:"name"`
	Vip int					`json:"vip"`
	Level int				`json:"level"`
}

type ReplyComment struct {
	Id int64				`json:"id"`
	Author					`json:"author"`
	Content string			`json:"content"`
	Date string				`json:"date"`
	Likes int64				`json:"likes"`
	ReplyAuthor Author		`json:"reply_author"`
}

type MetaComment struct {
	Id int64						`json:"id"`
	Author							`json:"author"`
	Content string					`json:"content"`
	Floor int						`json:"floor"`
	Date string						`json:"date"`
	Likes int64						`json:"likes"`
	ReplyComments []ReplyComment	`json:"comments_in_floor"`
}

type Common struct {
	Date string				`json:"date"`
	Likes int64				`json:"likes"`
	Coins int64				`json:"coins"`
	Collections int64		`json:"collections"`
	Shares int64			`json:"shares"`
	CommentNumbers int64	`json:"comment_number"`
}

type VideoInformation struct {
	Id int64				`json:"id"`
	BvCode string			`json:"bv_code"`
	VideoPath string		`json:"video_path"`
	CoverPath string		`json:"cover_path"`
	Title string 			`json:"title"`
	Brief string			`json:"brief"`
	Plays int64				`json:"plays"`
	Author					`json:"author"`
	Common					`json:"common"`
}

type GetVideoFile struct {
	BvCode string 		`uri:"bvCode"`
	FileName string 	`uri:"fileName"`
}