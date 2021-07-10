package main

import (
	"encoding/json"
	"fmt"
	"git.huawei.com/poissonsearch/wienerchain/contract/sdk"
	"git.huawei.com/poissonsearch/wienerchain/contract/sdk/smstub"
	"git.huawei.com/poissonsearch/wienerchain/proto/common"
	"strconv"
)

type Cert struct {
	Owner string
	Content string
	Issuer string
	Item string
	Point float64
	Date string
	Signature string
}

type Item struct {
	Id string
	Name string
	Description string
	Issuer string
	Point float64
}

type Person struct {
	Id string
	Name string
	Password string
	//PrivateKey string
	PublicKey string
}

type Issuer struct {
	Id string
	Name string
	Password string
	//PrivateKey string
	PublicKey string
}

type Chaincode struct {
}

//var log = logger.GetDefaultLogger()

func (cc Chaincode) Init(stub sdk.ContractStub) common.InvocationResponse {
	fmt.Println("Enter chaincode init function")
	return sdk.Success([]byte("Chaincode init success!"))
}

func (cc Chaincode) Invoke(stub sdk.ContractStub) common.InvocationResponse {
	funcName := stub.FuncName()
	args := stub.Parameters()

	switch funcName {
	case "addCert":
		return addCert(stub, args)
	case "queryCert":
		return queryCert(stub, args)
	case "queryCerts":
		return queryCerts(stub, args)
	case "addItem":
		return addItem(stub, args)
	case "queryItem":
		return queryItem(stub, args)
	case "queryItems":
		return queryItems(stub, args)
	case "addPerson":
		return addPerson(stub, args)
	case "queryPerson":
		return queryPerson(stub, args)
	case "addIssuer":
		return addIssuer(stub, args)
	case "queryIssuer":
		return queryIssuer(stub, args)
	case "getPublicKey":
		return getPublicKey(stub, args)
	}
	str := fmt.Sprintf("Func name is not correct, the function name is %s ", funcName)
	return sdk.Error(str)
}

func addCert(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 8 {
		return sdk.Error(fmt.Sprintf("Incorrect num of addCert args: Expect: 7, Actual: %d.", len(args)))
	}
	point, err := strconv.ParseFloat(string(args[4]), 64)
	if err != nil {
		return sdk.Error("Point is not a number.")
	}
	cert := Cert{
		Owner: string(args[0]),
		Content: string(args[1]),
		Issuer: string(args[2]),
		Item: string(args[3]),
		Point: point,
		Date: string(args[5]),
		Signature: string(args[6]),
	}
	key := string(args[7])
	certBytes, _ := json.Marshal(cert)
	if certBytes == nil {
		return sdk.Error("Encode Cert to byte error!")
	}

	err = stub.PutKV(key, certBytes)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success([]byte(key))
}

func queryCert(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of queryCert args: Expect: 1, Actual: %d.", len(args)))
	}
	certHash := string(args[0])
	certBytes, err := stub.GetKV(certHash)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(certBytes)
}

/*
 * args: list of hash
 */
func queryCerts(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) < 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of queryCerts args: Expect: at least 1, Actual: %d.", len(args)))
	}
	var certsMap []Cert
	for _, arg := range args {
		certHash := string(arg)
		certBytes, err := stub.GetKV(certHash)
		if err != nil {
			return sdk.Error(err.Error())
		}
		var cert Cert
		err = json.Unmarshal(certBytes, &cert)
		if err != nil {
			return sdk.Error("unmarshal json error")
		}
		//certsMap[certHash] = cert
		certsMap = append(certsMap, cert)
	}
	certsMapBytes, err := json.Marshal(certsMap)
	if certsMapBytes == nil || err != nil {
		return sdk.Error("marshal json error")
	}
	return sdk.Success(certsMapBytes)
}

func addItem(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 5 {
		return sdk.Error(fmt.Sprintf("Incorrect num of addItem args: Expect: 5, Actual: %d.", len(args)))
	}
	point, err := strconv.ParseFloat(string(args[4]), 64)
	if err != nil {
		return sdk.Error("Point is not a number")
	}
	item := Item{
		Id: string(args[0]),
		Name: string(args[1]),
		Description: string(args[2]),
		Issuer: string(args[3]),
		Point: point,
	}
	key := string(args[0])
	itemBytes, err := json.Marshal(item)
	if itemBytes == nil || err != nil {
		return sdk.Error("Encode item to byte error")
	}

	err = stub.PutKV(key, itemBytes)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success([]byte(key))
}

func queryItem(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of queryItem args: Expect: 1, Actual: %d.", len(args)))
	}
	itemHash := string(args[0])
	itemBytes, err := stub.GetKV(itemHash)
	if err != nil {
		sdk.Error(err.Error())
	}
	return sdk.Success(itemBytes)
}
/*
 * args: list of hash
 */
func queryItems(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) < 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of queryItems args: Expect: at least 1, Actual: %d.", len(args)))
	}
	var itemMap []Item
	for _, arg := range args {
		itemHash := string(arg)
		itemBytes, err := stub.GetKV(itemHash)
		if err != nil {
			return sdk.Error(err.Error())
		}
		var item Item
		err = json.Unmarshal(itemBytes, &item)
		if err != nil {
			return sdk.Error("unmarshal json error")
		}
		itemMap = append(itemMap, item)
	}
	itemMapBytes, err := json.Marshal(itemMap)
	if itemMapBytes == nil || err != nil {
		return sdk.Error("marshal json error")
	}
	return sdk.Success(itemMapBytes)
}

func addPerson(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 4 {
		return sdk.Error(fmt.Sprintf("Incorrect num of addPerson args: Expect: 4, Actual: %d.", len(args)))
	}
	person := Person{
		Id: string(args[0]),
		Name: string(args[1]),
		Password: string(args[2]),
		PublicKey: string(args[3]),
	}
	key := string(args[0])
	personBytes, err := json.Marshal(person)
	if personBytes == nil || err != nil {
		return sdk.Error("Encode person to byte error")
	}

	err = stub.PutKV(key, personBytes)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success([]byte(key))
}

func queryPerson(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of queryPerson args: Expect: 1, Actual: %d.", len(args)))
	}
	personHash := string(args[0])
	personBytes, err := stub.GetKV(personHash)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(personBytes)
}

func addIssuer(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 4 {
		return sdk.Error(fmt.Sprintf("Incorrect num of addIssuer args: Expect: 4, Actual: %d.", len(args)))
	}
	issuer := Issuer{
		Id: string(args[0]),
		Name: string(args[1]),
		Password: string(args[2]),
		PublicKey: string(args[3]),
	}
	key := string(args[0])
	issuerBytes, err := json.Marshal(issuer)
	if issuerBytes == nil || err != nil {
		return sdk.Error("Encode issuer to byte error")
	}

	err = stub.PutKV(key, issuerBytes)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success([]byte(key))
}

func queryIssuer(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of queryIssuer args: Expect: 1, Actual: %d.", len(args)))
	}
	issuerHash := string(args[0])
	issuerBytes, err := stub.GetKV(issuerHash)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(issuerBytes)
}

func getPublicKey(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) != 1 {
		return sdk.Error(fmt.Sprintf("Incorrect num of getPublicKey args: Expect: 1, Actual: %d.", len(args)))
	}
	key := string(args[0])
	keyBytes, err := stub.GetKV(key)
	if err != nil {
		return sdk.Error(err.Error())
	}
	var obj Person
	err = json.Unmarshal(keyBytes, &obj)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success([]byte(obj.PublicKey))
}

func main() {
	smstub.Start(&Chaincode{})
}