package mongodb

import (
	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type carrierDb struct {
	client *mongo.Client
}

func CreateCarrierDb(client *mongo.Client) db.CarrierDb {
	return &carrierDb{client: client}
}

func (c *carrierDb) GetCarriers() ([]models.FreightCarrierModel, error) {
	return nil, nil
}

func (c *carrierDb) GetCarrier(id string) (*models.FreightCarrierModel, error) {
	return nil, nil
}

func (c *carrierDb) CreateCarrier(carrier *models.FreightCarrierModel) (*models.FreightCarrierModel, error) {
	return nil, nil
}

func (c *carrierDb) UpdateCarrier(id string, carrier *models.FreightCarrierModel) (*models.FreightCarrierModel, error) {
	return nil, nil
}

func (c *carrierDb) DeleteCarrier(id string) error {
	return nil
}
