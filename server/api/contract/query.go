package contract

import (
	"errors"
	"fmt"
	"git.huawei.com/poissonsearch/wienerchain/proto/common"
	"git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/client"
	t "git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/utils"
	"github.com/golang/protobuf/proto"
)

func Query(args []string, function string, nodeName string) ([]byte, error) {
	gatewayClient, err := client.NewGatewayClient(ConfigPath)
	if err != nil {
		fmt.Printf("new gateway client error: %v\n", err)
		return nil, err
	}

	txID, err := t.GenerateTxID()
	if err != nil {
		fmt.Printf("generate tx id error: %v\n", err)
		return nil, err
	}
	rawMsg, err := gatewayClient.ContractRawMessage.BuildInvokeMessage(ChainName, txID, ContractName, function, args)
	if err != nil {
		fmt.Printf("contract raw message build invoke message error: %v\n", err)
		return nil, err
	}

	nodeMap := gatewayClient.Nodes
	node, ok := nodeMap[nodeName]
	if !ok {
		fmt.Printf("node not exist: %v\n", nodeName)
		return nil, errors.New(fmt.Sprintf("node not exist: %v\n", nodeName))
	}
	invokeResponse, err := node.ContractAction.Invoke(rawMsg)
	if err != nil {
		fmt.Printf("invoke error: %v\n", err)
		return nil, err
	}
	return processQueryResult(invokeResponse, nodeName)
}

func processQueryResult(invokeResponse *common.RawMessage, nodeName string) ([]byte, error) {
	response := &common.Response{}
	if err := proto.Unmarshal(invokeResponse.Payload, response); err != nil {
		fmt.Printf("unmarshal invoke response error: %v\n", err)
	}
	if response.Status == common.Status_SUCCESS {
		tx := &common.Transaction{}
		fmt.Println(tx)
		if err := proto.Unmarshal(response.Payload, tx); err != nil {
			fmt.Printf("unmarshal transaction error: %v\n", err)
			return nil, err
		}
		if tx.Payload ==nil {
			fmt.Println("tx.Payload nil")
			return nil, errors.New("tx.Payload nil")
		}
		txPayload := &common.TxPayload{}
		if err := proto.Unmarshal(tx.Payload, txPayload); err != nil {
			fmt.Printf("unmarshal tx payload error: %v\n", err)
			return nil, err
		}
		if txPayload.Data == nil {
			fmt.Println("txPayload.Data nil")
			return nil, errors.New("txPayload.Data nil")
		}
		txData := &common.CommonTxData{}
		if err:= proto.Unmarshal(txPayload.Data, txData); err != nil {
			fmt.Printf("unmarshal common tx data error: %v\n", err)
			return nil, err
		}
		if txData.Response == nil{
			fmt.Println("txData nil")
			return nil, errors.New("txData nil")
		}
		return parseResponse(response, nodeName, string(txData.Response.Payload))
	} else {
		return parseResponse(response, nodeName, "")
	}
}
