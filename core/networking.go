package core

import (
	"net"
	"strconv"
)

type MeshLink struct {
	Control net.Conn
	Data    net.Conn
}

func Listen(c *Network) ([]net.Listener, []Port) {
	exposedPorts := c.Expose

	attempts := make([]string, 0)
	active := make([]net.Listener, 0)
	failed := make([]Port, 0)
	errors := make([]error, 0)

	Logln("[INFO] Attempting to listen on exposed ports: ", exposedPorts)
	for _, v := range exposedPorts {
		listener, err := v.Listen()
		errors = append(errors, err)

		if err == nil {
			attempts = append(attempts, listener.Addr().String())
			active = append(active, listener)
		} else {
			attempts = append(attempts, "<nil>")
			failed = append(failed, v)
		}
	}

	LogTable(
		GenericCol[Port]{"Port", exposedPorts},
		GenericCol[string]{"Listener", attempts},
		GenericCol[error]{"Error", errors})

	return active, failed
}

func Connect(c *Network) ([]net.Conn, []Address) {
	addrs := c.Connect

	attempts := make([]string, 0)
	active := make([]net.Conn, 0)
	failed := make([]Address, 0)
	errors := make([]error, 0)

	Logln("[INFO] Attempting to connect to addresses: ", addrs)
	for _, v := range addrs {
		conn, err := v.Connect()
		errors = append(errors, err)

		if err == nil {
			attempts = append(attempts, conn.LocalAddr().String())
			active = append(active, conn)
		} else {
			attempts = append(attempts, "<nil>")
			failed = append(failed, v)
		}
	}

	LogTable(
		GenericCol[Address]{"Address", addrs},
		GenericCol[string]{"Conn", attempts},
		GenericCol[error]{"Error", errors})

	return active, failed
}

func (p Port) Listen() (net.Listener, error) {
	port := ":" + strconv.Itoa((int)(p))
	return net.Listen("tcp", port)
}

func (a Address) Connect() (net.Conn, error) {
	return net.Dial("tcp", string(a))
}

/*
	Goal:
	Send files. Maybe dirs.

	P0: Some nice API to connect to other computes/server
	P1: Separate control and data planes.

	P2: Optimize with multiple links ?
*/

const SOCKET_ADDR = "/tmp/socketaaaaaa"
