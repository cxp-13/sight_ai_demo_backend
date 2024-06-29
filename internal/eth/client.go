package eth

import (
	"context"
	"gateway_service/config"
	"gateway_service/internal/eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	OracleFiltererIns *contracts.OracleFilterer
	OracleContractIns *contracts.Oracle

	TransactOpts1     *bind.TransactOpts
	AdminTransactOpts *bind.TransactOpts
)

func Init() {
	client, err := ethclient.Dial(config.Cfg.Ethereum.BSC_RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	clientWS, err := ethclient.Dial(config.Cfg.Ethereum.BSC_RPC_URL_WS)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum clientWS: %v", err)
	}

	oracleFilterer, err := contracts.NewOracleFilterer(common.HexToAddress(config.Cfg.Ethereum.OracleAddress), clientWS)
	if err != nil {
		log.Fatalf("Failed to create OracleFilterer eth: %v", err)
	}
	OracleFiltererIns = oracleFilterer

	oracleContract, err := contracts.NewOracle(common.HexToAddress(config.Cfg.Ethereum.OracleAddress), client)
	if err != nil {
		log.Fatalf("Failed to create Oracle eth: %v", err)
	}
	OracleContractIns = oracleContract

	privateKey1, _ := crypto.HexToECDSA(config.Cfg.Ethereum.TestPrivate1)
	adminPrivateKey, _ := crypto.HexToECDSA(config.Cfg.Ethereum.AdminPrivate)
	chainId, _ := client.ChainID(context.Background())
	TransactOpts1, _ = bind.NewKeyedTransactorWithChainID(privateKey1, chainId)
	AdminTransactOpts, _ = bind.NewKeyedTransactorWithChainID(adminPrivateKey, chainId)
}
