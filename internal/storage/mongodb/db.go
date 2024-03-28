package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/adapter/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func NewDB(cfg *config.DB) (*mongo.Database, error) {
	uri := fmt.Sprintf("%s://%s:%s/", cfg.Connection, cfg.Host, cfg.Port)

	dbOptions := options.Client()
	dbOptions.ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, dbOptions)
	if err != nil {
		slog.Error("Failed to create new database connection", "error", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("Failed to ping the database", "error", err)
		return nil, err
	}
	slog.Info("Successfully connected to the database", "db", cfg.Name)
	return client.Database(cfg.Name), nil
}
