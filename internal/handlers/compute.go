package handlers

import (
	"encoding/json"
	"gateway_service/internal/eth"
	"gateway_service/internal/services"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
)

func Compute(c *gin.Context) {
	var computeReq services.ComputeRequest

	if err := c.ShouldBindJSON(&computeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	ASTBytes, _ := json.Marshal(computeReq.AST)
	log.Printf("req: %v\n", computeReq)
	go func() {
		eth.TransactOpts1.Value = new(big.Int).Mul(big.NewInt(10000000), big.NewInt(params.GWei)) // 0.01 eth
		_, err := eth.OracleContractIns.Compute(eth.TransactOpts1, computeReq.Numbers, ASTBytes)
		if err != nil {
			log.Print("Failed to call OracleContract.Compute: ", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"result": "Calculation has been submitted to the Oracle contract. Please wait for the result."})
}
