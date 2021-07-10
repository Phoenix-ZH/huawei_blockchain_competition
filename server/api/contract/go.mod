module contract

go 1.13

replace git.huawei.com/poissonsearch/wienerchain/proto => ../../../proto/go

replace git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk => ../../../wienerchain-go-sdk

replace gmssl => ../../../thirdparty/GmSSL/gmssl

replace schema => ../../schema

require (
	git.huawei.com/poissonsearch/wienerchain/proto v0.0.0
	git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	gmssl v0.0.0
	schema v0.0.0-00010101000000-000000000000
)
