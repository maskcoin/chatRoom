package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/server/commonutils"
	"net"
)

// UserProcess comment
type UserProcess struct {
	Conn net.Conn
}

// ServerProcessLogin  编写一个函数serverProcessLogin函数，专门处理登录的请求
func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码
	//1. 先从mes中取出 mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) err =", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//在声明一个 LoginResMes
	var loginResMes message.LoginResMes

	//如果用户的ID为100，密码=123456，认为合法，否则不合法
	if loginMes.UserID == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500 //500表示用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用"
	}

	//3. 将loginResMes序列化
	data, err := json.Marshal(&loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(&loginResMes) err =", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(&resMes)
	if err != nil {
		fmt.Println("json.Marshal(&resMes) err =", err)
		return
	}

	//6. 发送data，我们将其封装到writePkg(conn, data)函数中
	transfer := &commonutils.Transfer{
		Conn: up.Conn,
	}

	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn, data) err = ", err)
	}

	return
}
