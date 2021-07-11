module organization

go 1.13

replace server.com/api/contract => ../contract

replace git.huawei.com/poissonsearch/wienerchain/proto => ../../../proto/go

replace git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk => ../../../wienerchain-go-sdk

replace gmssl => ../../../thirdparty/GmSSL/gmssl

replace server.com/schema => ../../schema

require (
	github.com/Andrew-M-C/go.jsonvalue v1.1.0
	server.com/api/contract v0.0.0-00010101000000-000000000000
	server.com/schema v0.0.0-00010101000000-000000000000
)
