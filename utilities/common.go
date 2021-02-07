package utilities

//用于与前端通信的响应统一结构
type Resp struct {
	Code int					`json:"code"`		//响应代号
	Message string				`json:"message"`	//响应信息
	Data map[string]interface{}	`json:"data"`		//返回数据
}