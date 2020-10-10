package process

import (
	"fmt"
	"mygithub_code/chatRoom/client/model"
	"mygithub_code/chatRoom/common/message"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// CurUser comment // 我们在用户登录成功后，完成的对CurUser的初始化
var CurUser model.CurUser

// 编写一个方法，处理返回的NotifyUserStatusMes
func updateUsersStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	// 适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok {
		user = &message.User{
			UserID:     notifyUserStatusMes.UserID,
			UserStatus: notifyUserStatusMes.Status,
		}
	} else {
		user.UserStatus = notifyUserStatusMes.Status
	}

	onlineUsers[notifyUserStatusMes.UserID] = user

	outputOnlineUser()
}

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	// 遍历onlineUsers
	fmt.Println("当前在线用户列表：")
	for id := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}
