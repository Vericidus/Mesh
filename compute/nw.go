package main

import (
	"encoding/json"
	"mesh/core"
	"net"
	_ "net/http"
	_ "net/http/pprof"
	"time"
)

var ccfg *ComputeConfig

func worker() error {
	// go http.ListenAndServe("localhost:6060", nil)
	conn, err := net.Dial("tcp", ccfg.ServerAddress)
	if err != nil {
		core.Logf("[E] Connecting to server at:%s failed.\n", ccfg.ServerAddress)
		return err
	}

	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	start := time.Now()

	for range 100000 {

		signal := core.Signal{
			Id:     "Ping1",
			FromId: ccfg.Id,
			Time:   time.Now(),
			Type:   core.PING,
			Data:   nil,
		}

		enc.Encode(signal)
		var response core.Signal
		if err := dec.Decode(&response); err != nil {
			core.Logln("[E] Decoding server response failed.\n[EMSG]:", err)
			return err
		}

		core.Logln("[INFO] Message:", response)
	}

	elapsed := time.Since(start)
	core.Logln("Time Diff:", elapsed)

	return nil
}
