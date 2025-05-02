package core

import (
	"encoding/json"
	"net"
)

func NewEncoder(conn net.Conn) *json.Encoder {
	return json.NewEncoder(conn)
}

func NewDecoder(conn net.Conn) *json.Decoder {
	return json.NewDecoder(conn)
}
