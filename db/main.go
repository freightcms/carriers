package db

import (
	"context"

	"github.com/freightcms/carriers/models"
)

type CarrierDb interface {
	// GetCarrier returns a carrier by id.
	GetCarrier(ctx *context.Context, id string) (*models.FreightCarrier, error)
	// GetCarriers returns all carriers.
	GetCarriers(ctx *context.Context) ([]*models.FreightCarrier, error)
	// CreateCarrier creates a new carrier.
	CreateCarrier(ctx *context.Context, carrier *models.CreateFreightCarrier) (*models.FreightCarrier, error)
	// UpdateCarrier updates a carrier.
	UpdateCarrier(ctx *context.Context, carrier *models.FreightCarrier) (*models.FreightCarrier, error)
	// DeleteCarrier deletes a carrier.
	DeleteCarrier(ctx *context.Context, id string) error
	// Close closes the database connections. This function should be defered after
	// creating a new database connection.
	Close() error
}
