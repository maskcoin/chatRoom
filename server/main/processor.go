package main

import (
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/server/commonutils"
	"mygithub_code/chatRoom/server/process"
	"net"
)

// Processor struct
type Processor struct {
	Conn net.Conn
}

// serverProcessMess 编写一个ServerProcessMes 函数
// 功能：根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (p *Processor) serverProcessMess(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		userProcessor := &process.UserProcess{
			Conn: p.Conn,
		}

		err = userProcessor.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		userProcessor := &process.UserProcess{
			Conn: p.Conn,
		}

		err = userProcessor.ServerProcessRegister(mes)

	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		err = smsProcess.SendGroupMes(mes)

	default:
		fmt.Println("消息类型不存在，无法处理")
	}

	return
}

func (p *Processor) process() error {
	// 循环读客户端发送的信息
	for {
		// 这里我们将读取数据包，直接封装成一个函数readPkg(conn)，返回message, err
		tf := &commonutils.Transfer{
			Conn: p.Conn,
		}

		msg, err := tf.ReadPkg()

		if err != nil {
			fmt.Println("readPkg err = ", err)
			return err
		}

		fmt.Println("msg = ", msg)

		err = p.serverProcessMess(&msg)
		if err != nil {
			fmt.Println("serverProcessMess(conn, &msg) err = ", err)
			return err
		}
	}
}
