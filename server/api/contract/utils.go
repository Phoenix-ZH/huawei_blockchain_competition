package contract

import (
	"encoding/json"
	"fmt"
	"git.huawei.com/poissonsearch/wienerchain/proto/common"
	"schema"
)

func parseResponse(response *common.Response, node string, info string) ([]byte, error) {
	if response.Status == common.Status_SUCCESS {
		res := schema.Res{
			Info: info,
			Status: response.Status.String(),
			Node: node,
		}
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("marshal json error: %v\n", err)
			return nil, err
		}
		return resBytes, nil
	} else {
		res := schema.Res{
			Info: response.StatusInfo,
			Status: response.Status.String(),
			Node: node,
		}
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("marshal json error: %v\n", err)
			return nil, err
		}
		return resBytes, nil
	}
}
