package mongodb

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/emmrys-jay/anomaly-detection-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_CONTEXT_TIMEOUT = 90 * time.Second
var mongoClient *mongo.Client

func GetDB() *mongo.Client {
	return mongoClient
}

func ConnectDB() *mongo.Client {
	// tlsConfig := &tls.Config{}
	// tlsConfig.InsecureSkipVerify = true
	uri := strings.TrimSpace(os.Getenv("MONGODB_URL"))

	if uri == "" {
		log.Fatalln("No database URL was specified")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// connectOptions := options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig)

	ctx, cancel := context.WithTimeout(context.Background(), DB_CONTEXT_TIMEOUT)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatalln("could not connect to database url, err - ", err)
	}
	log.Println("db connected at ConnectDB ✓")

	mongoClient = client

	return client
}


func CreateSensorDataEntry(entry []model.SensorsData) error {
	coll := mongoClient.Database(model.DatabaseName).Collection(model.CollectionName)

	var data []any
	for _, v := range entry {
		data = append(data, v)
	}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		return err
	}

	log.Printf("Successfully added %v entry(ies) to the collection\n", len(entry))
	return err
}
