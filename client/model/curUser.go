package model

import (
	"mygithub_code/chatRoom/common/message"
	"net"
)

// CurUser struct
type CurUser struct {
	message.User
	Conn net.Conn
}
