package core

import (
	"time"
)

// REV: Optimize.
type Signal struct {
	// ComputeNameHash.SequenceNumber - Easily check for duplicated old
	// messages, know when delivery of a message in b/w failed (possible?)
	Id     string
	FromId string
	Time   time.Time
	Type   uint
	Data   any
}

const (
	PLACEHOLDER = 0
	PING
	PONG
	SIGNAL_DECODE_ERR
	CONTROL_CONN // Set as the control conn.
	DATA_CONN // Set as the data conn.
	PASSTHROUGH // For partially partitioned networks.
)

type SignalError struct {
	Message string
}

func CreateError(etype int, args ...string) any {
	var errorStruct any
	switch etype {
	case SIGNAL_DECODE_ERR:
		errorStruct = SignalError{
			Message: "Decoding message failure for signal with ID:" + args[0]}
	default:
		Logf("[E] Unknown error type: %d\n", etype)
	}

	return errorStruct
}
