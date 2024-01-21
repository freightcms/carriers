package services

import (
	"context"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"github.com/freightcms/carriers/schemas"
)

// CarrierService is the interface that provides carrier methods.
type CarrierService interface {
	// CreateCarrier Creates a new carrier and returns the created carrier. If the carrier
	// could not be created, an error is returned.
	CreateCarrier(ctx context.Context, schema *schemas.CreateCarrierSchema) (*schemas.CarrierSchema, error)
	// GetCarrier returns a carrier by id.
	GetCarrier(ctx context.Context, id string) (*schemas.CarrierSchema, error)
	// GetCarriers returns all carriers.
	GetCarriers(ctx context.Context) ([]*schemas.CarrierSchema, error)
	// UpdateCarrier updates a carrier.
	UpdateCarrier(ctx context.Context, schema *schemas.CarrierSchema) (*schemas.CarrierSchema, error)
	// DeleteCarrier deletes a carrier.
	DeleteCarrier(ctx context.Context, id string) error
}

// CarrierService is the interface that provides carrier methods.
type carrierService struct {
	db db.CarrierDb
}

func NewCarrierService(db db.CarrierDb) CarrierService {
	return &carrierService{db: db}
}

// CreateCarrier Creates a new carrier and returns the created carrier. If the carrier
// could not be created, an error is returned.
func (s *carrierService) CreateCarrier(ctx context.Context, schema *schemas.CreateCarrierSchema) (*schemas.CarrierSchema, error) {
	model := models.CreateFreightCarrier{
		Name: schema.Name,
		DBA:  schema.DBA,
	}
	carrier, err := s.db.CreateCarrier(ctx, &model)
	if err != nil {
		return nil, err
	}

	return &schemas.CarrierSchema{
		ID: carrier.ID,
		CreateCarrierSchema: schemas.CreateCarrierSchema{
			Name: carrier.Name,
			DBA:  carrier.DBA,
		},
	}, nil
}

// GetCarrier returns a carrier by id.
func (s *carrierService) GetCarrier(ctx context.Context, id string) (*schemas.CarrierSchema, error) {
	carrier, err := s.db.GetCarrier(ctx, id)
	if err != nil {
		return nil, err
	}

	return &schemas.CarrierSchema{
		ID: carrier.ID,
		CreateCarrierSchema: schemas.CreateCarrierSchema{
			Name: carrier.Name,
			DBA:  carrier.DBA,
		},
	}, nil
}

// GetCarriers returns all carriers.
func (s *carrierService) GetCarriers(ctx context.Context) ([]*schemas.CarrierSchema, error) {
	carriers, err := s.db.GetCarriers(ctx)
	if err != nil {
		return nil, err
	}

	carrierSlice := make([]*schemas.CarrierSchema, len(carriers))
	for i, carrier := range carriers {
		carrierSlice[i] = &schemas.CarrierSchema{
			ID: carrier.ID,
			CreateCarrierSchema: schemas.CreateCarrierSchema{
				Name: carrier.Name,
				DBA:  carrier.DBA,
			},
		}
	}

	return carrierSlice, nil
}

// UpdateCarrier updates a carrier.
func (s *carrierService) UpdateCarrier(ctx context.Context, schema *schemas.CarrierSchema) (*schemas.CarrierSchema, error) {
	model := models.FreightCarrier{
		ID: schema.ID,
		CreateFreightCarrier: models.CreateFreightCarrier{
			Name: schema.Name,
			DBA:  schema.DBA,
		},
	}
	carrier, err := s.db.UpdateCarrier(ctx, &model)
	if err != nil {
		return nil, err
	}

	return &schemas.CarrierSchema{
		ID: carrier.ID,
		CreateCarrierSchema: schemas.CreateCarrierSchema{
			Name: carrier.Name,
			DBA:  carrier.DBA,
		},
	}, nil
}

// DeleteCarrier deletes a carrier.
func (s *carrierService) DeleteCarrier(ctx context.Context, id string) error {
	return s.db.DeleteCarrier(ctx, id)
}
