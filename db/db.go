package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/qiniu/qmgo"
)

func New() *qmgo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: os.Getenv("DB_URI")})
	if err != nil {
		fmt.Println("error connecting client", err)
	}
	db := client.Database("supplychain")
	return db
}
