package utilities

//用于登录验证密码是否正确
type LoginCheck struct {
	Id int64															//用户的Id，用于设置cookie
	Name string 	`json:"username" form:"username" xml:"username"`	//登录的用户名/邮箱/手机号
	Password string `json:"password" form:"password" xml:"password"`	//登录密码
	Md5salt  int64														//密码的md5盐
}

//注册的信息结构
//用户名、手机号、邮箱必需输入其中一个，密码必须输入
type RegisterInformation struct {
	Name string			`json:"username" form:"username" xml:"username"`	//注册的用户名
	Password string		`json:"password" form:"password" xml:"password"`	//注册的密码
	Telephone int64		`json:"telephone" form:"telephone" xml:"telephone"`	//注册的手机号
	Email string		`json:"email" form:"email" xml:"email"`				//注册的邮箱
}
