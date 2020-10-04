package main

import (
	"fmt"
	"net"
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
