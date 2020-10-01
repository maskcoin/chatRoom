package message

// 消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

// Message struct
type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"Data"` // 消息的类型
}

// LoginMes struct定义两个消息,后面需要再增加
type LoginMes struct {
	UserID   int    `json:"userID"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

// LoginResMes struct
type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码 500 表示该用户未注册 200 表示登录成功
	Error string `json:"error"` // 返回的错误信息
}

// RegisterMes struct
type RegisterMes struct {
	//...
}
