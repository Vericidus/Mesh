package main

import (
	"mesh/core"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
)

func run() {

	go core.ReadCommands()
	go core.ProcessCommands()

	// TODO: Retry failures
	// _, _ = core.Connect(&c.Network)

	// Multi-tcp connection on the server side is beneficial for
	// tracking control and data independently.
	lns, _ := core.Listen(&cfg.Network)

	for {
		for _, ln := range lns {
			conn, err := ln.Accept()
			if err != nil {
				core.Logln("[E] Connection error.\n[EMSG]: ", err)
				continue
			}

			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) error {

	enc := core.NewEncoder(conn)
	dec := core.NewDecoder(conn)

	var ping core.Signal
	if err := dec.Decode(&ping); err != nil {
		core.Logln("[E] Decoding ping failed.\n[EMGS]: ", err)
	}

	// If reconnection then send a different signal type
	pong := core.Signal{
		Id:     uuid.New().String(),
		FromId: cfg.Name,
		Time:   time.Now(),
		Type:   core.PONG,
		Data:   nil,
	}

	if err := enc.Encode(pong); err != nil {
		core.Logln("[E] Error encoding.\n[EMSG]: ", err)
	}

	// REV: Any point in doing a conn healthCheck here too?

	// go readMessages()
	// go processMessages()
	sendSocketMessage()

	return nil
}

func sendSocketMessage() error {

	os.MkdirAll("tmp", 0755)
	os.Remove(core.SOCKET_ADDR)

	sock, err := net.Listen("unix", core.SOCKET_ADDR)
	if err != nil {
		core.Logln("[E] Failed to open socket.\n[EMSG]: ", err)
		return nil
	}

	conn, err := sock.Accept()
	if err != nil {
		core.Logln("[E] Failed to connect to socket.\n[EMSG]: ", err)
		return nil
	}

	// enc := core.NewEncoder(conn)
	dec := core.NewDecoder(conn)

	var ping core.Signal
	if err := dec.Decode(&ping); err != nil {
		core.Logln("[E] Decoding ping failed.\n[EMGS]: ", err)
	}

	core.Logln("Socket message: ", core.Prettify(ping))

	return nil
}
