package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"
	"mygithub_code/chatRoom/myutils"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 4096)
	_, err = myutils.ReadN(buf[:4], conn, 4)
	if err != nil {
		return
	}

	// 根据buf[:4]，转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen来读取消息内容
	_, err = myutils.ReadN(buf[:pkgLen], conn, int(pkgLen))
	if err != nil {
		return
	}

	//把buf[:pkgLen]反序列化成Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个len给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var lenBytes [4]byte
	binary.BigEndian.PutUint32(lenBytes[:], pkgLen)

	//发送长度
	n, err := conn.Write(lenBytes[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(lenBytes[:]) err = ", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	return
}
