package main

import (
	"mesh/core"
)

var Config ServerConfig

func run(c ServerConfig) {
	Config = c

	// TODO: Retry failures
	// _, _ = core.Connect(&c.Network)

	lns, _ := core.Listen(&c.Network)
	_, _ = core.Accept(lns)

}
