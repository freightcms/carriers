package mongodb

import (
	"context"
	"errors"
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
func (db *carrierDb) GetCarrier(id string) (*models.FreightCarrier, error) {
	carrier := new(models.FreightCarrier)
	filter := bson.M{"_id": id}
	err := db.Collection().FindOne(context.TODO(), filter).Decode(&carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

// GetCarriers returns all carriers.
func (db *carrierDb) GetCarriers() ([]*models.FreightCarrier, error) {
	var carriers []*models.FreightCarrier
	cursor, err := db.Collection().Find(context.TODO(), bson.M{})
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
func (db *carrierDb) CreateCarrier(carrier *models.CreateFreightCarrier) (*models.FreightCarrier, error) {
	query := bson.M{
		"$or": []bson.M{
			{"name": carrier.Name},
			{"dba": carrier.DBA},
		},
	}
	count, err := db.Collection().CountDocuments(context.TODO(), query)
	if err != nil {
		return nil, errors.New("Carrier already exists")
	}
	if count > 0 {
		return nil, errors.New("Carrier already exists")
	}

	result, err := db.Collection().InsertOne(context.TODO(), &carrier)
	if err != nil {
		return nil, err
	}
	var carrierResult models.FreightCarrier
	filter := bson.M{"_id": result.InsertedID}
	err = db.Collection().FindOne(context.Background(), filter).Decode(&carrierResult)
	if err != nil {
		return nil, err
	}
	return &carrierResult, nil
}

// UpdateCarrier updates a carrier.
func (db *carrierDb) UpdateCarrier(carrier *models.FreightCarrier) (*models.FreightCarrier, error) {
	carrier.UpdatedAtUTC = time.Now().Format(time.RFC3339)
	objectID, err := primitive.ObjectIDFromHex(carrier.ID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": carrier}
	_, err = db.Collection().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

// DeleteCarrier deletes a carrier.
func (db *carrierDb) DeleteCarrier(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	_, err = db.Collection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (db *carrierDb) Collection() *mongo.Collection {
	return db.client.Database("carriers").Collection("carriers")
}
