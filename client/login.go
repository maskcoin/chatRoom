package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"net"
	"os"
)

//写一个函数，完成登录
func login(userID int, userPwd string) (err error) {
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
	//7.1 先把data的长度，发送给服务器
	//先获取到 data的长度->转成一个表示长度的[]byte
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

	//这里还需要处理服务器返回的消息
	mes, err = readPkg(conn)
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
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}