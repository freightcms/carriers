package mongodb

import (
	"context"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type carrierDb struct {
	databaseName string
	collection   string
	session      mongo.Session
}

func CreateCarrierDb(session mongo.Session) db.CarrierDb {
	return &carrierDb{
		session:      session,
		databaseName: "freightcms",
		collection:   "carriers",
	}
}

func (c *carrierDb) GetCarriers(ctx context.Context) ([]models.FreightCarrierModel, error) {
	cursor, err := c.session.Client().Database(c.databaseName).Collection(c.collection).Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.FreightCarrierModel
	for cursor.Next(ctx) {
		var carrier models.FreightCarrierModel
		if err = cursor.Decode(&carrier); err != nil {
			return results, err
		}
		results = append(results, carrier)
	}
	return results, cursor.Err()
}

func (c *carrierDb) GetCarrier(ctx context.Context, id string) (*models.FreightCarrierModel, error) {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": idPrimitive,
	}
	carrier := c.session.Client().Database(c.databaseName).Collection(c.collection).FindOne(ctx, &filter)
	var result models.FreightCarrierModel
	if err := carrier.Decode(&models.FreightCarrierModel{}); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *carrierDb) CreateCarrier(ctx context.Context, carrier *models.FreightCarrierModel) error {
	if err := c.session.StartTransaction(); err != nil {
		return err
	}
	if _, err := c.session.Client().Database(c.databaseName).Collection(c.collection).InsertOne(ctx, carrier); err != nil {
		return err
	}
	if err := c.session.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}

func (c *carrierDb) UpdateCarrier(ctx context.Context, id string, carrier *models.FreightCarrierModel) error {
	if err := c.session.StartTransaction(); err != nil {
		return err
	}
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if _, err = c.session.Client().Database(c.databaseName).Collection(c.collection).UpdateByID(ctx, idPrimitive, carrier); err != nil {
		return err
	}
	if err = c.session.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}

func (c *carrierDb) DeleteCarrier(ctx context.Context, id string) error {
	if err := c.session.StartTransaction(); err != nil {
		return err
	}
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if _, err = c.session.Client().Database(c.databaseName).Collection(c.collection).DeleteOne(ctx, bson.M{"_id": idPrimitive}); err != nil {
		return err
	}
	if err = c.session.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}
