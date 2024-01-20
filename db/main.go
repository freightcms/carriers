package db

import "github.com/freightcms/carriers/models"

type CarrierDb interface {
	// GetCarrier returns a carrier by id.
	GetCarrier(id string) (*models.FreightCarrier, error)
	// GetCarriers returns all carriers.
	GetCarriers() ([]*models.FreightCarrier, error)
	// CreateCarrier creates a new carrier.
	CreateCarrier(carrier *models.FreightCarrier) (*models.FreightCarrier, error)
	// UpdateCarrier updates a carrier.
	UpdateCarrier(carrier *models.FreightCarrier) (*models.FreightCarrier, error)
	// DeleteCarrier deletes a carrier.
	DeleteCarrier(id string) error
	// Close closes the database connections. This function should be defered after
	// creating a new database connection.
	Close() error
}
