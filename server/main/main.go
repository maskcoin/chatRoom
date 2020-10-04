package main

import (
	"fmt"
	"mygithub_code/chatRoom/server/model"
	"net"
	"time"
)

// 处理和客户端的通信
func allProcess(conn net.Conn) {
	defer conn.Close()

	p := &Processor{
		Conn: conn,
	}

	err := p.process()

	if err != nil {
		fmt.Println("p.process() err =", err)
		return
	}
}

// 这里我们编写一个函数，完成对UserDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//当服务器启动时，我们就去初始化redis连接池
	initPool(":6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {
	// 提示信息
	fmt.Println("服务器[新的结构]在8889端口监听...")

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
		go allProcess(conn)
	}
}
