package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/server/commonutils"
	"mygithub_code/chatRoom/server/model"
	"net"
)

// UserProcess comment
type UserProcess struct {
	Conn net.Conn
	// 增加一个字段，表示该Conn是哪个用户的
	UserID int
}

// NotifyOtherOnlineUsers comment
func (up *UserProcess) NotifyOtherOnlineUsers() {
	// 遍历userMgr.onlineUsers，然后一个一个发送
	for id, up2 := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == up.UserID {
			continue
		}
		// 开始通知[单独写一个方法]
		up2.NotifyMeOnline(up.UserID)
	}
}

// NotifyMeOnline comment
func (up *UserProcess) NotifyMeOnline(userID int) (err error) {
	// 组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.Status = message.UserOnline

	// 将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) err =", err)
		return
	}

	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err =", err)
		return
	}

	tf := &commonutils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err=", err)
	}

	return
}

// ServerProcessRegister comment
func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1. 先从mes中取出 mes.Data，并直接反序列化成LoginMes
	var regMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &regMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &regMes) err =", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//再声明一个 LoginResMes
	var registerResMes message.RegisterResMes

	// 我们需要到redis数据库完成验证
	err = model.MyUserDao.Register(&regMes.User)
	if err != nil {
		if err == model.ErrorUserExists {
			registerResMes.Code = 505 //505表示用户已注册
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506 //未知错误
			registerResMes.Error = "服务器内部错误"
		}

	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	//3. 将registerResMes序列化
	data, err := json.Marshal(&registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(&registerResMes) err =", err)
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

	//再声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 我们需要到redis数据库完成验证
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNotExists {
			loginResMes.Code = 500 //500表示用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = 403 //403表示用户密码不正确
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505 //未知错误
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		// 这里，因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		// 将登录成功的用户的userID赋给up
		up.UserID = loginMes.UserID
		userMgr.AddOnlineUser(up)
		// 将当前在线用户的id 放入到loginResMes.UsersID
		// 遍历 userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}

		// 通知其他在线用户，我上线了
		up.NotifyOtherOnlineUsers()

		fmt.Println(user, "登录成功")
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
