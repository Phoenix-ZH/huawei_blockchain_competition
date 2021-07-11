package models

type MyCerts struct {
	Id string `bson:"id"` // 用户id
	Hash []string `bson:"hash"`   // 证书hash列表
}

type MyItems struct {
	Id string `bson:"id"` // 用户id
	ItemIds []string `bson:"item_ids"` // 课程项目id列表
}

type CombineCerts struct {
	Key string `bson:"key"`  // 组合证书hash
	Hash []string `bson:"hash"`  // 对应的子证书hash
}