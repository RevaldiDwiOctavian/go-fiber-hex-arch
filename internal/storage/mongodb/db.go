package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	db         *mongo.Database
	collection string
}

func New(db *mongo.Database, collection string) (*DB, error) {
	return &DB{
		db:         db,
		collection: collection,
	}, nil
}
