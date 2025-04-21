package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var ccfg *ComputeConfig

func worker() error {
	conn, err := net.Dial("tcp", ccfg.ServerAddress)
	if err != nil {
		fmt.Printf("[E] Connecting to server at:%s failed.\n", ccfg.ServerAddress)
		return err
	}

	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	signal := Signal{
		"Ping1",
		ccfg.Id,
		time.Now(),
		PING,
		nil,
	}

	for {
		enc.Encode(signal)
		var response Signal
		if err := dec.Decode(&response); err != nil {
			fmt.Println("[E] Decoding server response failed.\n[EMSG]:", err)
			return err
		}

		fmt.Println("[INFO] Message:", response)
	}
}
