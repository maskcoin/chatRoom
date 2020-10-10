package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/client/cutils"
	"mygithub_code/chatRoom/common/message"
	"net"
	"os"
)

// ShowMenu 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("------------恭喜xxx登录成功----------")
	fmt.Println("------------1、显示在线用户列表----------")
	fmt.Println("------------2、发送消息----------")
	fmt.Println("------------3、信息列表----------")
	fmt.Println("------------4、退出系统----------")
	fmt.Println("请选择（1-4）：")

	var key int
	smsProcess := &SmsProcess{}
	var content string
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说点什么：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的不对，请重新输入")
	}
}

// 和服务器保持通讯
func processServerMes(conn net.Conn) {
	// 创建一个transfer实例
	tf := &cutils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端正在等待读取服务器推送的消息")
		mes, err := tf.ReadPkg()

		if err != nil {
			fmt.Println("tf.ReadPkg() err = ", err)
			os.Exit(1)
		}

		// 如果读取到了mes，又是下一步处理逻辑
		// fmt.Println("mes = ", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: // 有人上线了
			//处理
			//1.取出NotifyUserStatusMes
			//2.把这个用户的信息，状态，保存到客户端的map[int]*User中
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes) err =", err)
				return
			}

			updateUsersStatus(&notifyUserStatusMes)

		case message.SmsMesType:
			// 处理服务器转发的消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了一个我没法识别的消息类型")
		}
	}
}
