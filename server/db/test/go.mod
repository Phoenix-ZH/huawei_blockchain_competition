module test

go 1.13

replace server.com/api/common => ../../api/common

replace server.com/api/contract => ../../api/contract

replace server.com/api/registration => ../../api/registration

replace server.com/api/learner => ../../api/learner

replace server.com/api/organization => ../../api/organization

replace server.com/schema => ../../schema

replace git.huawei.com/poissonsearch/wienerchain/proto => ../../../proto/go

replace git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk => ../../../wienerchain-go-sdk

replace gmssl => ../../../thirdparty/GmSSL/gmssl

require (
	server.com/api/common v0.0.0-00010101000000-000000000000
	server.com/api/contract v0.0.0-00010101000000-000000000000
	server.com/api/learner v0.0.0-00010101000000-000000000000
	server.com/api/organization v0.0.0-00010101000000-000000000000
	server.com/api/registration v0.0.0-00010101000000-000000000000
	server.com/schema v0.0.0-00010101000000-000000000000
)
