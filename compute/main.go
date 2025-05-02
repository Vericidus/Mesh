package main

import (
	"flag"
	"mesh/core"
)

var cfg ComputeConfig

func main() {
	core.Logln("Mesh Compute v", core.MESH_VERSION)
	core.Logln("---")

	// Flags
	cpath := flag.String("config", core.DEFAULT_COMPUTE_CONFIG_PATH, "Path to Compute Config")
	flag.Parse()

	// Config Parsing
	core.Logln("[INFO] Loading compute config from path: ", *cpath)
	core.ParseConfig(&cfg, *cpath)

	// Run server
	core.Logln("Initializing compute with:\n", core.PrettyPrint(cfg))
	run(cfg)
}

// go run . -config "../test/config/.mesh.config"
