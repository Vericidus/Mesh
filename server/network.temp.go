package main

// import (
// 	"encoding/json"
// 	"mesh/core"
// 	"net"
// 	"strconv"
// 	"time"
// )

// var scfg *ServerConfig

// func server() error {
// 	ln, err := listen(scfg.Port)
// 	if err != nil {
// 		return err
// 	}
// 	defer close(ln)

// 	core.Logln("[INFO] Server started successfully.")

// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			core.Logln("[E] Connection error.\n[EMSG]:", err)
// 			continue
// 		}

// 		go handleConn(conn)
// 	}
// }

// func handleConn(conn net.Conn) {
// 	defer conn.Close()
// 	req := json.NewDecoder(conn)
// 	res := json.NewEncoder(conn)

// 	for {
// 		var sgl core.Signal
// 		if err := req.Decode(&sgl); err != nil {
// 			res.Encode(core.Signal{
// 				Id:     sgl.Id,
// 				FromId: scfg.Id,
// 				Time:   time.Now(),
// 				Type:   core.SIGNAL_DECODE_ERR,
// 				Data:   core.CreateError(core.SIGNAL_DECODE_ERR, sgl.Id),
// 			})

// 			continue
// 		}

// 		// REV: go handleSignal? What patterns would make it worth it?
// 		handleSignal(sgl, res)
// 	}
// }

// func handleSignal(sgl core.Signal, res *json.Encoder) {
// 	core.Logln("[INFO] Message:", sgl)
// 	switch sgl.Type {
// 	case core.PING:
// 		err := res.Encode(core.Signal{
// 			Id:     "2",
// 			FromId: "Server",
// 			Time:   time.Now(),
// 			Type:   core.PING,
// 			Data:   string("PONG"),
// 		})

// 		if err != nil {
// 			core.Logln("[E] Encoding response signal failed.\n[EMSG]:", err)
// 		}
// 	}
// }

// func listen(cfgPort uint16) (net.Listener, error) {
// 	port := ":" + strconv.Itoa((int)(cfgPort))
// 	ln, err := net.Listen("tcp", port)
// 	if err != nil {
// 		core.Logf("[E] Error listening to %d", cfgPort)
// 		return nil, err
// 	}

// 	return ln, nil
// }

// func close(ln net.Listener) {
// 	if err := ln.Close(); err != nil {
// 		core.Logln("[E] Closing connection failed.", err)
// 	}
// }
