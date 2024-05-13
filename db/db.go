package db

import "github.com/freightcms/carriers/models"

type CarrierDb interface {
	GetCarriers() ([]models.FreightCarrierModel, error)
	GetCarrier(id string) (models.FreightCarrierModel, error)
	CreateCarrier(carrier models.FreightCarrierModel) (models.FreightCarrierModel, error)
	UpdateCarrier(id string, carrier models.FreightCarrierModel) (models.FreightCarrierModel, error)
	DeleteCarrier(id string) error
}
