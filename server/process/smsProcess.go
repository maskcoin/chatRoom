package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/server/commonutils"
	"net"
)

// SmsProcess struct
type SmsProcess struct {
}

// SendGroupMes 转发消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 先取出mes中的内容
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err =", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}

	// 遍历服务器端的onlineUsers map[int]*UserProcess
	// 将消息转发出去
	for id, up := range userMgr.onlineUsers {
		if id != smsMes.UserID {
			sp.sendMesToEachOnlineUser(data, up.Conn)
		}
	}

	return
}

func (sp *SmsProcess) sendMesToEachOnlineUser(data []byte, conn net.Conn) (err error) {
	// 创建一个transfer
	tf := &commonutils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err=", err)
		return
	}

	return
}
