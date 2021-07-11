package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server.com/config"
	"time"
)

// NewDatabase 返回一个新的数据库客户端连接
func NewDatabase(config *config.Database) (*mongo.Database, error) {
	var (
		client *mongo.Client
		err error
		db *mongo.Database
	)
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return nil, err
	}
	db = client.Database("my_db")
	if err := InitDatabase(db, config); err != nil {
		return nil, err
	}
	return db, nil
}

// InitDatabase 初始化了数据库
func InitDatabase(db *mongo.Database, config *config.Database) error {
	return nil
}
