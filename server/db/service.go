package db

import (
	"db/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Account  *handlers.LearnerHandler
	Category *handlers.OrgHandler
}
func NewService(db *mongo.Database) Service {
	return Service{
		Account:  handlers.NewLearnerHandler(db),
		Category: handlers.NewOrgHandler(db),
	}
}

