package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/myutils"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 4096)
	_, err = myutils.ReadN(buf[:4], conn, 4)
	if err != nil {
		return
	}

	// 根据buf[:4]，转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen来读取消息内容
	_, err = myutils.ReadN(buf[:pkgLen], conn, int(pkgLen))
	if err != nil {
		return
	}

	//把buf[:pkgLen]反序列化成Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个len给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var lenBytes [4]byte
	binary.BigEndian.PutUint32(lenBytes[:], pkgLen)

	//发送长度
	n, err := conn.Write(lenBytes[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(lenBytes[:]) err = ", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	return
}

// 编写一个函数serverProcessLogin函数，专门处理登录的请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("writePkg(conn, data) err = ", err)
	}

	return
}

// serverProcessMess 编写一个ServerProcessMes 函数
// 功能：根据客户端发送的消息种类不同，决定调用哪个函数来处理
func serverProcessMess(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")
	}

	return
}

// 处理和客户端的通信
func process(conn net.Conn) {
	defer conn.Close()

	// 循环读客户端发送的信息
	for {
		// 这里我们将读取数据包，直接封装成一个函数readPkg(conn)，返回message, err
		msg, err := readPkg(conn)

		if err != nil {
			fmt.Println("readPkg err = ", err)
			return
		}

		fmt.Println("msg = ", msg)
		err = serverProcessMess(conn, &msg)
		if err != nil {
			fmt.Println("serverProcessMess(conn, &msg) err = ", err)
			return
		}
	}
}

func main() {
	// 提示信息
	fmt.Println("服务器在8889端口监听...")

	ln, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	// 一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ln.Accept() err = ", err)
			continue
		}
		// 一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
