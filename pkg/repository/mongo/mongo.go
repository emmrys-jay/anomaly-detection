package mongodb

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"strings"
	"time"

	"github.com/emmrys-jay/anomaly-detection-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_CONTEXT_TIMEOUT = 90 * time.Second
var mongoClient *mongo.Client
var newMongoClient *mongo.Client

func GetDB() *mongo.Client {
	return mongoClient
}

func ConnectDB() *mongo.Client {
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	uri := strings.TrimSpace(os.Getenv("MONGODB_URL"))

	if uri == "" {
		log.Fatalln("No database URL was specified")
	}

	connectOptions := options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig)
	// Create a new client and connect to the server


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

func ConnectNewDB() *mongo.Client {
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	uri := strings.TrimSpace(os.Getenv("NEW_URL"))

	if uri == "" {
		log.Fatalln("No database URL was specified")
	}

	connectOptions := options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig)
	// Create a new client and connect to the server


	ctx, cancel := context.WithTimeout(context.Background(), DB_CONTEXT_TIMEOUT)
	defer cancel()
	client, err := mongo.Connect(ctx, connectOptions)

	if err != nil {
		log.Fatalln("could not connect to database url, err - ", err)
	}
	log.Println("db connected at ConnectDB ✓")

	newMongoClient = client

	return client
}

func MoveData() error {
	collOld := mongoClient.Database(model.DatabaseName).Collection(model.CollectionName)
	collNew := newMongoClient.Database(model.DatabaseName).Collection(model.CollectionName)

	var values = []model.SensorsData{}
	cur, err := collOld.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	err = cur.All(context.Background(), &values)
	if err != nil {
		return err
	}

	inter := make([]any, 0, len(values))
	for _, v := range values {
		inter = append(inter, v)
	}

	_, err = collNew.InsertMany(context.Background(), inter)
	if err != nil {
		return err
	}
	return nil
}


func CreateSensorDataEntry(entry []model.SensorsData) error {
	coll := newMongoClient.Database(model.DatabaseName).Collection(model.CollectionName)
	
	now := time.Now()
	var data []any
	for _, v := range entry {
		v.CreatedAt = now
		data = append(data, v)
	}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		return err
	}

	log.Printf("Successfully added %v entry(ies) to the collection\n", len(entry))
	return err
}

func LabelData(filters []model.LabelFilter) error {
	coll := newMongoClient.Database(model.DatabaseName).Collection(model.CollectionName)

	writeModels := make([]mongo.WriteModel, 0, len(filters))
	for _, label := range filters {
		model := mongo.NewUpdateManyModel().
			SetFilter(bson.D{
				{Key: "Time", Value: bson.D{{Key: "$gte", Value: label.StartTime}}},
				{Key: "Time", Value: bson.D{{Key: "$lte", Value: label.EndTime}}},
				}).
			SetUpdate(bson.D{{Key: "$set", Value: bson.D{
				{Key: "Anomaly", Value: label.Anomaly},
			}}})

		writeModels = append(writeModels, model)
	}

	bulkOption := options.BulkWrite().SetOrdered(false)
	_, err := coll.BulkWrite(context.Background(), writeModels, bulkOption)
	if err != nil {
		return err
	}

	return nil
}

func DeleteInvalidData() error {
	coll := newMongoClient.Database(model.DatabaseName).Collection(model.CollectionName)

	filter := bson.D{{
		Key: "Latitude",
		Value: 0,
	}}

	_, err := coll.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func LabelNoneData() error {
	coll := newMongoClient.Database(model.DatabaseName).Collection(model.CollectionName)

	model := mongo.NewUpdateManyModel().
		SetFilter(bson.D{{Key: "Anomaly", Value: bson.D{{Key: "$exists", Value: false}}}}).
		SetUpdate(bson.D{{Key: "$set", Value: bson.D{{Key: "Anomaly", Value: "None"}}}})

	
	writeModels := []mongo.WriteModel{model}
	
	bulkOption := options.BulkWrite().SetOrdered(false)
	_, err := coll.BulkWrite(context.Background(), writeModels, bulkOption)
	if err != nil {
		return err
	}

	return nil
}