package mongodb

import (
	"context"
	"fmt"
	"time"

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
func (db *carrierDb) GetCarrier(ctx context.Context, id string) (*models.FreightCarrier, error) {
	carrier := new(models.FreightCarrier)
	filter := bson.M{"_id": id}
	err := db.Collection().FindOne(ctx, filter).Decode(&carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

// GetCarriers returns all carriers.
func (db *carrierDb) GetCarriers(ctx context.Context) ([]*models.FreightCarrier, error) {
	var carriers []*models.FreightCarrier
	cursor, err := db.Collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &carriers); err != nil {
		return nil, err
	}
	return carriers, nil
}

// CreateCarrier creates a new carrier.
func (db *carrierDb) CreateCarrier(ctx context.Context, carrier *models.CreateFreightCarrier) (*models.FreightCarrier, error) {
	count, err := db.Collection().CountDocuments(ctx, bson.M{"name": carrier.Name})
	if err != nil {
		return nil, fmt.Errorf("Error checking if carrier exists: %s", err.Error())
	}
	if count > 0 {
		return nil, fmt.Errorf("Carrier with name %s already exists", carrier.Name)
	}
	count, err = db.Collection().CountDocuments(ctx, bson.M{"dba": carrier.DBA})
	if err != nil {
		return nil, fmt.Errorf("Error checking if carrier exists: %s", err.Error())
	}
	if count > 0 {
		return nil, fmt.Errorf("Carrier with dba %s already exists", carrier.DBA)
	}

	result, err := db.Collection().InsertOne(ctx, &carrier)
	if err != nil {
		return nil, err
	}
	var carrierResult models.FreightCarrier
	filter := bson.M{"_id": result.InsertedID}
	err = db.Collection().FindOne(ctx, filter).Decode(&carrierResult)
	if err != nil {
		return nil, err
	}
	return &carrierResult, nil
}

// UpdateCarrier updates a carrier.
func (db *carrierDb) UpdateCarrier(ctx context.Context, carrier *models.FreightCarrier) (*models.FreightCarrier, error) {
	carrier.UpdatedAtUTC = time.Now().Format(time.RFC3339)
	objectID, err := primitive.ObjectIDFromHex(carrier.ID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": carrier}
	_, err = db.Collection().UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

// DeleteCarrier deletes a carrier.
func (db *carrierDb) DeleteCarrier(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	_, err = db.Collection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (db *carrierDb) Collection() *mongo.Collection {
	return db.client.Database("carriers").Collection("carriers")
}
