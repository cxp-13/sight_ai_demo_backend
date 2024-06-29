package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway_service/config"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

// ASTNode 表示抽象语法树中的节点
type ASTNode struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
	Left  *ASTNode    `json:"left,omitempty"`
	Right *ASTNode    `json:"right,omitempty"`
}

// ComputeRequest 表示用户的计算请求数据结构
type ComputeRequest struct {
	Numbers []*big.Int `json:"numbers"`
	AST     ASTNode    `json:"ast"`
}

type ComputeResponse struct {
	Result int `json:"result"`
}

func CallLocalComputeService(req ComputeRequest) (int, error) {
	log.Printf("Calling local compute service with request: %v", req)
	var computeRes ComputeResponse

	// 构造本地计算服务的 URL
	localComputeURL := config.Cfg.ComputeService.URL

	// 将用户的 JSON 数据转换为字节流
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return -1, fmt.Errorf("Failed to marshal JSON: %v", err)
	}

	// 创建 HTTP POST 请求到本地计算服务
	resp, err := http.Post(localComputeURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return -1, fmt.Errorf("Failed to send request to local compute service: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("Failed to read response from local compute service: %v", err)
	}

	err = json.Unmarshal(respBody, &computeRes)
	if err != nil {
		return -1, fmt.Errorf("Failed to unmarshal response: %v", err)
	}
	log.Printf("Compute result: %d", computeRes)

	return computeRes.Result, nil
}
