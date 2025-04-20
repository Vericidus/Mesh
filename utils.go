package main

import (
	"encoding/json"
	"fmt"
)

func (c *ComputeConfig) start() error {
	v, _ := json.MarshalIndent(c, "", " ")
	fmt.Printf("Initializing with:\n%s\n", v)
	ccfg = c
	return worker()
}

func (c *ServerConfig) start() error {
	v, _ := json.MarshalIndent(c, "", " ")
	fmt.Printf("Initializing with:\n%s\n", v)
	scfg = c
	return server()
}

func (c *ComputeConfig) getID() string {
	return c.Id
}

func (c *ServerConfig) getID() string {
	return c.Id
}
