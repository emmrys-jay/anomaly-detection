package mongodb

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_CONTEXT_TIMEOUT = 90 * time.Second
var mongoClient *mongo.Client

func GetDB() *mongo.Client {
	return mongoClient
}

func ConnectDB() *mongo.Client {
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	uri := strings.TrimSpace(os.Getenv("MONGODB_URL"))

	connectOptions := options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig)

	ctx, cancel := context.WithTimeout(context.Background(), DB_CONTEXT_TIMEOUT)
	defer cancel()
	client, err := mongo.Connect(ctx, connectOptions)

	if err != nil {
		log.Fatalln("could not connect to database url, err - ", err)
	}
	log.Println("db connected at ConnectDB ✓")

	mongoClient = client

	return client
}
