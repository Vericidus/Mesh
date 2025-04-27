package core

import (
	"time"
)

// REV: Optimize.
type Signal struct {
	Id     string
	FromId string
	Time   time.Time
	Type   uint
	Data   any
}

const (
	PING = 1
	SIGNAL_DECODE_ERR
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
