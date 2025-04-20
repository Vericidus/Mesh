package main

import (
	"fmt"
	"os"
)

const MESH_VERSION = "0.10"

// Worth keeping for now to quickly iterate and spin up different nodes.
// Go doesn't seem to have strong dead code elim. Will separate for performance's sake.
var IS_SERVER = true

func main() {
	fmt.Println("Mesh v", MESH_VERSION)
	fmt.Println("---")

	// TODO: Improve
	if len(os.Args) == 2 {
		IS_SERVER = false
	}

	fmt.Println("[INFO] Server mode:", IS_SERVER)

	// TODO: change config path from flag.
	const configPath = "./files/"
	fmt.Println("[INFO] Searching [MeshConfig] at path:", configPath)

	cfg, err := parseConfig(configPath)
	if err != nil {
		fmt.Println("[FE] Failed to read config.\n[EMSG]:", err)
		panic("CONFIG")
	}

	if err := cfg.start(); err != nil {
		fmt.Printf("[FE] %s is exiting.\n", cfg.getID())
		fmt.Println("[EMSG]:", err)
		panic("EXECUTION")
	}
}
