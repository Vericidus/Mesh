package main

import (
	"mesh/core"
	"os"
)

func main() {
	core.Logln("Mesh Compute v", core.MESH_VERSION)
	core.Logln("---")

	if len(os.Args) > 2 {
		core.Logln("Too many arguments", )
	}
}
