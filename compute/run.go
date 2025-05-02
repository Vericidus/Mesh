package main

import (
	_ "encoding/json"
	"mesh/core"
	"net"
	"time"

	"github.com/google/uuid"
)

var config ComputeConfig

func run(c ComputeConfig) {
	config = c

	// TODO: Retry failures & accept conns from computes.
	// _, _ := core.Listen(&c.Network)
	conns, _ := core.Connect(&c.Network)

	for _, conn := range conns {
		go handleConn(conn)
	}

	// TODO: retry failed conns.

	// encs := core.NewEncoders(conns)
	// decs := core.NewDecoders(conns)

	// // Send / receive PING and measure latency, RTT.
	// // Tabulate this for better routing later on.
	// // TODO: Do this periodically in a go routine.
	// core.SmartHealthCheck(encs, decs)
	
	// // Receive a message
	// // Start a shell
	// // Receive a msg from a program IPC socket
	// // Process and send to server
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

	enc.Encode(ping)
	var pong core.Signal
	if err := dec.Decode(&pong); err != nil {
		core.Logln("[E] Decoding pong failed.\n[EMGS]: ", err)
	}

	healthChecks(ping, pong)

	go readMessages()
	// go processMessages()

	return nil
}

func readMessages() {
	// net.Listen("unix")
}

// Have a UDP socket open too for extremely high throughput logs.
