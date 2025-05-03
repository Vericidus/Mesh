package main

import (
	"flag"
	"mesh/core"
)

var cfg ServerConfig

func main() {
	core.Logln("Mesh Server v", core.MESH_VERSION)
	core.Logln("---")

	// Flags
	cpath := flag.String("config", core.DEFAULT_SERVER_CONFIG_PATH, "Path to Server Config")
	flag.Parse()

	// Config Parsing
	core.Logln("[INFO] Loading server config from path: ", *cpath)
	core.ParseConfig(&cfg, *cpath)

	// Run server
	core.Logln("Initializing server with:\n", core.Prettify(cfg))
	run()
}

// go run . -config "../test/config/.mesh.server.config"
