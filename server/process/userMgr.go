package process

import "fmt"

// 因为Usermgr实例在服务器端有且只有一个，我们将其定义为全局变量
var userMgr *UserMgr

// 完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// UserMgr struct
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// AddOnlineUser 完成对onlineUsers添加
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserID] = up
}

// DeleteOnlineUser delete
func (um *UserMgr) DeleteOnlineUser(userID int) {
	delete(um.onlineUsers, userID)
}

// GetAllOnlineUsers select
func (um *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return um.onlineUsers
}

// GetOnlineUserByID comment
func (um *UserMgr) GetOnlineUserByID(userID int) (up *UserProcess, err error) {
	// 如何从map中取出一个值，带检测的方式
	up, ok := um.onlineUsers[userID]
	if !ok { // 说明你要查找的这个用户，当前不在线
		err = fmt.Errorf("用户%d 不存在", userID)
	}
	return
}
