package mongodb

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type carrierDb struct {
	client *mongo.Client
}

// NewCarrierDb creates a new carrier database instance.
func NewCarrierDb(uri string) (db.CarrierDb, error) {
	ctx, _ := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &carrierDb{
		client: client,
	}, nil
}

func (db *carrierDb) Close() error {
	return db.client.Disconnect(context.Background())
}

// GetCarrier returns a carrier by id.
func (db *carrierDb) GetCarrier(id string) (*models.FreightCarrier, error) {
	carrier := &models.FreightCarrier{}
	filter := bson.M{"_id": id}
	err := db.client.Database("carriers").Collection("carriers").FindOne(context.Background(), filter).Decode(carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

// GetCarriers returns all carriers.
func (db *carrierDb) GetCarriers() ([]*models.FreightCarrier, error) {
	var carriers []*models.FreightCarrier
	cursor, err := db.client.Database("carriers").Collection("carriers").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var carrier models.FreightCarrier
		if err := cursor.Decode(&carrier); err != nil {
			return nil, err
		}
		carriers = append(carriers, &carrier)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return carriers, nil
}

// CreateCarrier creates a new carrier.

func (db *carrierDb) CreateCarrier(carrier *models.FreightCarrier) (*models.FreightCarrier, error) {
	carrier.CreatedAtUTC = time.Now().UTC().Format(time.RFC3339)
	carrier.UpdatedAtUTC = time.Now().UTC().Format(time.RFC3339)
	carrier.ID = uuid.New().String()
	result, err := db.client.Database("carriers").Collection("carriers").InsertOne(context.TODO(), &carrier)
	if err != nil {
		return nil, err
	}
	idBytes, err := result.InsertedID.(primitive.ObjectID).MarshalText() // get jus the text out of the ObjectID("...")
	if err != nil {
		return nil, err
	}
	carrier.ID = string(idBytes)
	return carrier, nil
}

// UpdateCarrier updates a carrier.
func (db *carrierDb) UpdateCarrier(carrier *models.FreightCarrier) (*models.FreightCarrier, error) {
	carrier.UpdatedAtUTC = time.Now().Format(time.RFC3339)
	filter := bson.M{"_id": carrier.ID}
	update := bson.M{"$set": carrier}
	_, err := db.client.Database("carriers").Collection("carriers").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

// DeleteCarrier deletes a carrier.
func (db *carrierDb) DeleteCarrier(id string) error {
	filter := bson.M{"_id": id}
	_, err := db.client.Database("carriers").Collection("carriers").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
