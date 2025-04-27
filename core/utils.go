package core

import (
	"encoding/json"
	"fmt"
)

const MESH_VERSION = "0.10"

const DEFAULT_SERVER_CONFIG_PATH = ".mesh.server.config"
const DEFAULT_COMPUTE_CONFIG_PATH = "./mesh.compute.config"

func Logln(v ...any) {
	fmt.Println(fmt.Sprint(v...))
}

func Logf(s string, a ...any) {
	fmt.Printf(s, a...)
}

// Pretty print structs
type Struct any

func PrettyPrint(a Struct) string {
	b, _ := json.MarshalIndent(a, "", "  ")
	return string(b)
}
