package process

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
)

func outputGroupMes(mes *message.Message) (err error) {
	// 显示即可
	// 1.反序列化mes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err=", err)
		return
	}

	// 显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
	fmt.Println()
	return
}
