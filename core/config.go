package core

import (
	"encoding/json"
	"os"
)

type BaseConfig struct {
	General
	Network Network
	Storage Storage
}

type General struct {
	Id          string
	Name        string
	Description []string
}

type Port uint16
type Address string

type Network struct {
	Expose  []Port
	Connect []Address
}

type Storage struct {
	Path string
}

// Cfg must be a pointer to a struct.
func ParseConfig(cfg any, path string) {
	cfgData, err := os.ReadFile(path)
	if err != nil {
		Logln("[E] Failed to read config.\n[EMSG]:", err)
		panic("CONFIG")
	}

	if err := json.Unmarshal(cfgData, cfg); err != nil {
		Logln("[E] Failed to decode config.\n[EMSG]:", err)
		panic("CONFIG")
	}
}
