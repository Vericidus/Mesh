package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ComputeConfig struct {
	Id string

	Name        string
	Description []string

	ServerAddress string
}

type ServerConfig struct {
	Id string

	Name        string
	Description []string

	Port uint16
}

type MeshConfig interface {
	init()
}

func parseConfig(path string) (MeshConfig, error) {
	if !IS_SERVER {
		cfg, err := readConfig(path + ".mesh.config")
		if err != nil {
			return nil, err
		}
		return parseComputeCfg(&cfg)
	} else {
		cfg, err := readConfig(path + ".mesh.server.config")
		if err != nil {
			return nil, err
		}
		return parseServerCfg(&cfg)
	}
}

func readConfig(path string) ([]byte, error) {
	cfg, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[E] Failed to read config.")
		fmt.Println("[EMSG]:", err)
		return nil, err
	}
	return cfg, nil
}

func parseComputeCfg(cfg *[]byte) (*ComputeConfig, error) {
	var cCfg ComputeConfig
	if err := json.Unmarshal(*cfg, &cCfg); err != nil {
		return nil, err
	}
	return &cCfg, nil
}

func parseServerCfg(cfg *[]byte) (*ServerConfig, error) {
	var sCfg ServerConfig
	if err := json.Unmarshal(*cfg, &sCfg); err != nil {
		return nil, err
	}
	return &sCfg, nil
}
