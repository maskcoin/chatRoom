package myutils

import "net"

// ReadN comment
func ReadN(buf []byte, conn net.Conn, len int) (n int, err error) {
	var nleft = len
	for nleft != 0 {
		var retN int
		retN, err = conn.Read(buf[n:len])
		nleft -= retN
		n += retN
		if err != nil {
			return
		}
	}
	return
}
