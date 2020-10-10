package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/client/cutils"
	"mygithub_code/chatRoom/common/message"
	"os"
)

// SmsProcess struct
type SmsProcess struct {
}

// SendGroupMes 发送群聊的消息
func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	//2. 准备通过conn准备发送消息给服务器
	var mes message.Message
	mes.Type = message.SmsMesType

	//3. 创建一个LoginMes结构体
	smsMes := message.SmsMes{}
	smsMes.User = CurUser.User
	smsMes.Content = content

	//4.将smsMes 序列化
	data, err := json.Marshal(&smsMes)
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
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err =", err)
		return
	}

	return
}
