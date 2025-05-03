package main

import (
	_ "encoding/json"
	"mesh/core"
	"net"
	"time"

	"github.com/google/uuid"
)

func run() {

	// TODO: Retry failures & accept conns from computes.
	// _, _ := core.Listen(&c.Network)
	conns, _ := core.Connect(&cfg.Network)

	for _, conn := range conns {
		go handleConn(conn)
	}

	// TODO: retry failed conns.

	select {}
}

func handleConn(conn net.Conn) error {
	enc := core.NewEncoder(conn)
	dec := core.NewDecoder(conn)

	// If reconnection then send a different signal type
	ping := core.Signal{
		Id:     uuid.New().String(),
		FromId: cfg.Name,
		Time:   time.Now(),
		Type:   core.PING,
		Data:   nil,
	}

	if err := enc.Encode(ping); err != nil {
		core.Logln("[E] Encoding ping failed.\n[EMGS]: ", err)
	}

	var pong core.Signal
	if err := dec.Decode(&pong); err != nil {
		core.Logln("[E] Decoding pong failed.\n[EMGS]: ", err)
	}

	// TODO: Make this periodic and affect routing decisions.
	healthChecks(ping, pong)

	go readMessages()
	go processMessages()

	return nil
}

func readMessages() error {
	conn, err := net.Dial("unix", core.SOCKET_ADDR)

	if err != nil {
		core.Logln("[E] Failed to open socket.\n[EMSG]: ", err)
		return err
	}

	// enc := core.NewEncoder(conn)
	enc := core.NewEncoder(conn)

	pong := core.Signal{
		Id:     uuid.New().String(),
		FromId: cfg.Name,
		Time:   time.Now(),
		Type:   core.PONG,
		Data:   nil,
	}
	if err := enc.Encode(&pong); err != nil {
		core.Logln("[E] Decoding ping failed.\n[EMGS]: ", err)
	}

	core.Logln("Socket message: ", core.Prettify(pong))

	return nil
}

func processMessages() {}
