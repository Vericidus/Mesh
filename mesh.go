package main

import (
	"fmt"
	"os"
)

const MESH_VERSION = "0.10"

// REV: Might be useful to be able to switch between worker and client with one const flip.
// Will simplify reverse engineering potential.
const IS_SERVER = true

func main() {
	fmt.Printf("Mesh v%s\n", MESH_VERSION)
	fmt.Println("---")

	fmt.Println("Server mode:", IS_SERVER)

	// TODO: change config path from flag.
	const configPath = "./files/"
	fmt.Println("Searching [ComputeConfig] at path:", configPath)

	cfg, err := parseConfig(configPath)
	if err != nil {
		fmt.Println("[FE] Failed to read config. Exiting.")
		os.Exit(1)
	}

	cfg.init()
}
