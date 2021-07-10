module registration

go 1.13

replace contract => ../contract

replace git.huawei.com/poissonsearch/wienerchain/proto => ../../../proto/go

replace git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk => ../../../wienerchain-go-sdk

replace gmssl => ../../../thirdparty/GmSSL/gmssl

replace schema => ../../schema

require (
	contract v0.0.0-00010101000000-000000000000
	github.com/Andrew-M-C/go.jsonvalue v1.1.0
	schema v0.0.0-00010101000000-000000000000
)
