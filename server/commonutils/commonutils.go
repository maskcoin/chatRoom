package commonutils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/myutils"
	"net"
)

// Transfer struct
//这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [4096]byte //缓存
}

// ReadPkg 用来读取
func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = myutils.ReadN(transfer.Buf[:4], transfer.Conn, 4)
	if err != nil {
		return
	}

	// 根据buf[:4]，转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])

	// 根据pkgLen来读取消息内容
	_, err = myutils.ReadN(transfer.Buf[:pkgLen], transfer.Conn, int(pkgLen))
	if err != nil {
		return
	}

	//把buf[:pkgLen]反序列化成Message
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

// WritePkg 用来写入
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个len给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var lenBytes [4]byte
	binary.BigEndian.PutUint32(lenBytes[:], pkgLen)

	//发送长度
	n, err := transfer.Conn.Write(lenBytes[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(lenBytes[:]) err = ", err)
		return
	}

	//发送消息本身
	_, err = transfer.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	return
}
