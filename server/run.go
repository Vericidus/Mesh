package main

import "mesh/core"

var Config ServerConfig

func run(c ServerConfig) {
	Config = c

	listener, err := (Config.Network.Expose).Listen()
	if err != nil {
		core.Logf("[FE] Cannot listen on port %d.\n", Config.Network.Expose)
		panic("NETWORKING")
	}
	

	core.AcceptLinks(listener)
}
