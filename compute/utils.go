package main

import (
	"mesh/core"
	"time"
)

func healthChecks(ping core.Signal, pong core.Signal) {
	roundTripTime := time.Since(ping.Time)
	clientToServerTime := pong.Time.Sub(ping.Time)
	serverToClientTime := time.Since(pong.Time)

	core.LogTable(
		core.GenericCol[time.Duration]{Header: "RTT", Rows: []time.Duration{roundTripTime}},
		core.GenericCol[time.Duration]{Header: "C2S", Rows: []time.Duration{clientToServerTime}},
		core.GenericCol[time.Duration]{Header: "S2C", Rows: []time.Duration{serverToClientTime}},
	)
}
