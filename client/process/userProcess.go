package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/client/cutils"
	"mygithub_code/chatRoom/common/message"
	"net"
	"os"
)

// UserProcess struct
type UserProcess struct {
	//字段...
}

// Login 写一个函数，完成登录
func (up *UserProcess) Login(userID int, userPwd string) (err error) {
	//下一步就要开始定协议
	//1. 连接到服务器
	conn, err := net.Dial("tcp", ":8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		os.Exit(1)
	}

	//延时关闭
	defer conn.Close()

	//2. 准备通过conn准备发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3. 创建一个LoginMes结构体
	loginMes := message.LoginMes{
		UserID:  userID,
		UserPwd: userPwd,
	}

	//4.将loginMes 序列化
	data, err := json.Marshal(&loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		os.Exit(1)
	}

	//5. 把data赋给了mes.Data
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(&mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		os.Exit(1)
	}

	//7. 到这个时候，data就是我们要发送的消息
	tf := &cutils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err =", err)
		return
	}

	// 这里还需要处理服务器返回的消息

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err =", err)
		return
	}

	var loginResMes message.LoginResMes
	//将mes的Data部分，反序列化成LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginResMes) err= ", err)
		return
	}

	if loginResMes.Code == 200 {
		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserID = userID
		CurUser.UserStatus = message.UserOnline
		// 这里我们还需要启动一个协程，该协程保持和服务器端的通讯。如果服务器有数据推送过来，则接受并显示在客户端的终端
		// 可以显示当前在线用户列表，遍历loginResMes.UsersID
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersID {
			fmt.Println("用户id:\t", v)
			// 完成客户端的onlineUsers的初始化
			user := &message.User{
				UserID:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}

		fmt.Println()
		fmt.Println()

		go processServerMes(conn)

		//1. 显示登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

// Register comment
func (up *UserProcess) Register(userID int, userPwd, userName string) (err error) {
	//1. 连接到服务器
	conn, err := net.Dial("tcp", ":8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		os.Exit(1)
	}

	//延时关闭
	defer conn.Close()

	//2. 准备通过conn准备发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3. 创建一个RegisterMes结构体
	user := message.User{
		UserID:   userID,
		UserPwd:  userPwd,
		UserName: userName,
	}

	regMes := message.RegisterMes{
		User: user,
	}

	//4.将regMes 序列化
	data, err := json.Marshal(&regMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		os.Exit(1)
	}

	//5. 把data赋给了mes.Data
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(&mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		os.Exit(1)
	}

	//7. 到这个时候，data就是我们要发送的消息
	tf := &cutils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)

	if err != nil {
		fmt.Println("tf.WritePkg(data) err =", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err =", err)
		return
	}

	var registerResMes message.RegisterResMes
	//将mes的Data部分，反序列化成RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &registerResMes) err= ", err)
		return
	}

	if registerResMes.Code == 200 {
		fmt.Println("注册成功，重新登录一把")
	} else {
		fmt.Println(registerResMes.Error)
	}

	return
}
