package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
	}

	App struct {
		Name string
		Env  string
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		app,
		db,
		http,
	}, nil

}

func NewDB(cfg *DB) (*mongo.Database, error) {
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
