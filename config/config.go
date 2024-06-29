package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type Config struct {
	Ethereum       EthereumConfig
	ComputeService ComputeServiceConfig
}

type EthereumConfig struct {
	BSC_RPC_URL    string
	BSC_RPC_URL_WS string
	TestAddr1      string
	AdminAddr      string
	TestPrivate1   string
	AdminPrivate   string
	OracleAddress  string
}

type ComputeServiceConfig struct {
	URL string
}

var Cfg Config

func LoadConfig() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Cfg = Config{
		Ethereum: EthereumConfig{
			BSC_RPC_URL:    cfg.Section("ethereum").Key("bsc_rpc_url").String(),
			BSC_RPC_URL_WS: cfg.Section("ethereum").Key("bsc_rpc_url_ws").String(),

			TestAddr1: cfg.Section("ethereum").Key("test_addr_1").String(),
			AdminAddr: cfg.Section("ethereum").Key("admin_addr").String(),

			TestPrivate1:  cfg.Section("ethereum").Key("test_private_1").String(),
			AdminPrivate:  cfg.Section("ethereum").Key("admin_private").String(),
			OracleAddress: cfg.Section("ethereum").Key("oracle_address").String(),
		},
		ComputeService: ComputeServiceConfig{
			URL: cfg.Section("compute_service").Key("url").String(),
		},
	}
}
