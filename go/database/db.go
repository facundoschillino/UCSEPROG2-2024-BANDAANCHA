package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	Concect() error
	Disconnect() error
	GetClient() *mongo.Client
}
