package handlers

import (
	"context"
	"db/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type LearnerHandler struct {
	db *mongo.Database
}

func NewLearnerHandler(db *mongo.Database) *LearnerHandler {
	return &LearnerHandler{
		db,
	}
}

func (h *LearnerHandler) Find(uId string) (*[]models.MyCerts, error) {
	var (
		res []models.MyCerts
		cursor *mongo.Cursor
		err error
	)
	if cursor, err = h.db.Collection("my_certs").Find(
		context.TODO(),
		struct {
			Id string `bson:"id"`
		}{Id: uId},
		options.Find().SetSkip(0),
		options.Find().SetLimit(2)); err != nil {
		return nil, err
	}

	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err = cursor.All(context.TODO(), &res); err != nil {
		log.Fatal(err)
	}
	return &res, nil
}

func (h *LearnerHandler) Insert(certs *models.MyCerts) error {
	var (
		iResult    *mongo.InsertOneResult
		id         primitive.ObjectID
		err error
	)
	if iResult, err = h.db.Collection("my_certs").InsertOne(context.TODO(), certs); err != nil {
		fmt.Print(err)
		return err
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())
	return nil
}
