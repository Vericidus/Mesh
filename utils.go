package main

import (
	"encoding/json"
	"fmt"
)

func (c *ComputeConfig) init() {
	v, _ := json.MarshalIndent(c, "", " ")
	fmt.Printf("Initializing with: %s\n", v)

	worker(c)
}

func (c *ServerConfig) init() {
	v, _ := json.MarshalIndent(c, "", " ")
	fmt.Printf("Initializing with: %s\n", v)

	server(c)
}
