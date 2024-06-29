package services

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway_service/internal/eth"
	"gateway_service/internal/eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func ListenOnChain() {

	computeReqs := make(chan *contracts.OracleComputationRequested)
	computeCompleteReqs := make(chan *contracts.OracleComputationCompleted)

	computeReqsEven, err := eth.OracleFiltererIns.WatchComputationRequested(&bind.WatchOpts{Context: context.Background()}, computeReqs, nil, nil)
	if err != nil {
		log.Printf("WatchComputationRequested error: %v", err)
		return
	}
	computeCompleteReqsEven, err := eth.OracleFiltererIns.WatchComputationCompleted(&bind.WatchOpts{Context: context.Background()}, computeCompleteReqs, nil)
	if err != nil {
		log.Printf("WatchComputationCompleted error: %v", err)
		return
	}

	defer computeReqsEven.Unsubscribe()
	defer computeCompleteReqsEven.Unsubscribe()

	for {
		select {
		case err := <-computeReqsEven.Err():
			log.Printf("Subscription error: %v", err)
			return
		case err := <-computeCompleteReqsEven.Err():
			log.Printf("Subscription error: %v", err)
			return
		case event := <-computeReqs:
			fmt.Printf("ComputationRequested Event:\n")
			fmt.Printf("  Request ID: %x\n", event.RequestId)
			fmt.Printf("  User: %s\n", event.User.Hex())
			handleComputationRequestEvent(event.RequestId)
		case event := <-computeCompleteReqs:
			fmt.Printf("ComputationCompleted Event:\n")
			fmt.Printf("  Request ID: %x\n", event.RequestId)
			fmt.Printf("  Result: %d\n", event.Result)
		}
	}
}

func handleComputationRequestEvent(reqId [32]byte) {
	var ast ASTNode
	numbers, logicBytes, _, err := eth.OracleContractIns.GetComputationData(&bind.CallOpts{Pending: true}, reqId)
	err = json.Unmarshal(logicBytes, &ast)
	if err != nil {
		log.Fatalf("Failed to unmarshal AST: %v", err)
		return
	}
	result, err := CallLocalComputeService(ComputeRequest{
		Numbers: numbers,
		AST:     ast,
	})

	_, err = eth.OracleContractIns.Callback(eth.AdminTransactOpts, reqId, big.NewInt(int64(result)))
	if err != nil {
		log.Fatalf("Failed to call callback: %v", err)
		return
	}
}
