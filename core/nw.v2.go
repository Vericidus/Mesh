package core

import "net"

type MeshLink struct {
	Control net.Conn
	Data net.Conn
}

/* 
	Goal: 
	Send files. Maybe dirs.

	P0: Some nice API to connect to other computes/server

	P2: Optimize with multiple links ?
*/