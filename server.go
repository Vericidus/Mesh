package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

var scfg *ServerConfig

func server() error {
	ln, err := listen(scfg.Port)
	if err != nil {
		return err
	}
	defer close(ln)

	fmt.Println("[INFO] Server started successfully.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[E] Connection error.\n[EMSG]:", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	req := json.NewDecoder(conn)
	res := json.NewEncoder(conn)

	for {
		var sgl Signal
		if err := req.Decode(&sgl); err != nil {
			res.Encode(Signal{
				sgl.Id,
				scfg.Id,
				time.Now(),
				SIGNAL_ERR,
				createError(SIGNAL_ERR, sgl.Id),
			})

			continue
		}

		// REV: go handleSignal? What patterns would make it worth it?
		handleSignal(sgl, res)
	}
}

func handleSignal(sgl Signal, res *json.Encoder) {
	fmt.Println("[INFO] Message:", sgl)
	switch sgl.Type {
	case PING:
		err := res.Encode(Signal{
			"2",
			"Server",
			time.Now(),
			PING,
			map[string]any{"pong": "pong"},
		})

		if err != nil {
			fmt.Println("[E] Encoding response signal failed.\n[EMSG]:", err)
		}
	}
}

func listen(cfgPort uint16) (net.Listener, error) {
	port := ":" + strconv.Itoa((int)(cfgPort))
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("[E] Error listening to %d", cfgPort)
		return nil, err
	}

	return ln, nil
}

func close(ln net.Listener) {
	if err := ln.Close(); err != nil {
		fmt.Println("[E] Closing connection failed.", err)
	}
}
