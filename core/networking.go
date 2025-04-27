package core

import (
	"net"
	"strconv"
)

// func Listen(portUint uint16) (net.Listener, error) {
// 	port := ":" + strconv.Itoa((int)(portUint))
// 	return net.Listen("tcp", port)
// }

// func CloseListener(ln net.Listener) error {
// 	err := ln.Close()
// 	if err != nil {
// 		Logln("[E] Closing connection failed.\n[EMSG]: ", err)
// 	}
// 	return err
// }

func Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		Logf("[E] Failed to connect to server at:%s failed.\n", addr)
		return nil, err
	}

	return conn, err
}

type ALink interface {
	Listen()
	Close()
	Transfer()
}



func (p Port) Listen() (net.Listener, error) {
	port := ":" + strconv.Itoa((int)(p))
	return net.Listen("tcp", port)
}
// func (ln net.Listener) Close() error {
// 	err := ln.Close()
// 	if err != nil {
// 		Logln("[E] Closing connection failed.\n[EMSG]: ", err)
// 	}
// 	return err
// }

func (a Address) Connect() (net.Conn, error) {
	return nil, nil
}
