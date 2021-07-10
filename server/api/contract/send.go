package contract

import (
	"errors"
	"fmt"
	"git.huawei.com/poissonsearch/wienerchain/proto/common"
	"git.huawei.com/poissonsearch/wienerchain/proto/nodeservice"
	"git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/client"
	"git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/node"
	t "git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/utils"
	"github.com/golang/protobuf/proto"
	"strings"
	"time"
)

const WaitTime = 60

func Send(args []string, function string, nodes string) ([]byte, error) {
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
		fmt.Printf("contract raw messessage build invoke message error： %v\n", err)
		return nil, err
	}

	nodeNames := strings.Split(nodes, ",")
	nodeMap := gatewayClient.Nodes
	var invokeResponses []*common.RawMessage
	for _, nodeName := range nodeNames {
		wNode, ok := nodeMap[nodeName]
		if !ok {
			fmt.Printf("node not exist: %v\n", nodeName)
			return nil, errors.New(fmt.Sprintf("node not exist: %v\n", nodeName))
		}
		var invokeResponse *common.RawMessage
		invokeResponse, err = wNode.ContractAction.Invoke(rawMsg)
		if err != nil {
			fmt.Printf("invoke err: %v\n", err)
			return nil, err
		}
		invokeResponses = append(invokeResponses, invokeResponse)
	}

	transactionRawMsg, err := gatewayClient.ContractRawMessage.BuildTransactionMessage(invokeResponses)
	if err != nil {
		fmt.Printf("build transaction message error: %v\n", err)
		return nil, err
	}

	return processResult(transactionRawMsg, nodeMap, txID, nodeNames[0])
}

func processResult(transactionRawMsg *common.RawMessage, nodeMap map[string]*node.WNode, txID string, listenNodename string) ([]byte, error) {
	wNode, ok := nodeMap[listenNodename]
	if !ok {
		fmt.Printf("node not exist: %v\n", listenNodename)
		return nil, errors.New(fmt.Sprintf("node not exist: %v\n", listenNodename))
	}

	event, err := wNode.EventAction.Listen(ChainName)
	if err != nil {
		fmt.Printf("event action listen error: %v\n", err)
		return nil, err
	}

	in := make(chan string)
	go listen(event, txID, in)
	consenterNodeNames := strings.Split(ConsensusPeers, ",")
	var transactionResponses []*common.Response
	for _, nodeName := range consenterNodeNames {
		n, ok := nodeMap[nodeName]
		if !ok {
			fmt.Printf("node not exist: %v\n", nodeName)
			return nil, errors.New(fmt.Sprintf("node not exist: %v\n", nodeName))
		}
		transactionResponse, err := n.ContractAction.Transaction(transactionRawMsg)
		if err != nil {
			fmt.Printf("invoke error: %v\n", err)
			return nil, err
		}
		txResponse := &common.Response{}
		if err := proto.Unmarshal(transactionResponse.Payload, txResponse); err != nil {
			fmt.Printf("unmarshal transaction response error: %v\n", err)
			return nil, err
		}
		transactionResponses = append(transactionResponses, txResponse)
	}

	select {
	case recv := <-in:
		for i, response := range transactionResponses {
			if response.Status == common.Status_SUCCESS {
				return parseResponse(response, consenterNodeNames[i], recv)
			} else {
				return parseResponse(response, consenterNodeNames[i], "")
			}
		}
		return nil, err
	case <-time.After(WaitTime * time.Second):
		fmt.Printf("Invoke time out\n")
		return nil, err
	}
}

func listen(event nodeservice.EventService_RegisterBlockEventClient, txID string, in chan string) {
	for {
		responseMsg, err := event.Recv()
		if err != nil {
			fmt.Printf("event receive response message error: %v", err)
			return
		}
		res := &common.BlockResult{}
		err = proto.Unmarshal(responseMsg.Payload, res)
		if err != nil {
			fmt.Printf("unmarshal block result error: %v", err)
			return
		}

		for _, txRes := range res.TxResults {
			if txRes.TxId == txID {
				in <- txRes.Status.String()
				return
			}
		}
	}
}
