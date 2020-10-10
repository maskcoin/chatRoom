package message

// 消息类型
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
	// ...
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
	Code    int    `json:"code"`  // 返回状态码 500 表示该用户未注册 200 表示登录成功
	UsersID []int  `json:"users"` // 增加字段，保存用户id的切片
	Error   string `json:"error"` // 返回的错误信息
}

// RegisterMes struct
type RegisterMes struct {
	User User `json:"user"`
}

// RegisterResMes struct
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示该用户已经被注册 200 表示注册成功
	Error string `json:"error"` // 返回的错误信息
}

// NotifyUserStatusMes struct
type NotifyUserStatusMes struct {
	UserID int `json:"userID"`
	Status int `json:"status"`
}

// SmsMes struct 增加一个SmsMes // 发送的
type SmsMes struct {
	User           // User   `json:"user"`
	Content string `json:"content"`
}
