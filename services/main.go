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
	CreateCarrier(schema *schemas.CreateCarrierSchema) (*schemas.CarrierSchema, error)
	// GetCarrier returns a carrier by id.
	GetCarrier(id string) (*schemas.CarrierSchema, error)
	// GetCarriers returns all carriers.
	GetCarriers() ([]*schemas.CarrierSchema, error)
	// UpdateCarrier updates a carrier.
	UpdateCarrier(schema *schemas.CarrierSchema) (*schemas.CarrierSchema, error)
	// DeleteCarrier deletes a carrier.
	DeleteCarrier(id string) error
}

type carrierService struct {
	ctx *context.Context
	db  db.CarrierDb
}

// NewCarrierService returns a new carrier service. When creating a new service it should pass
// a context which can be used to cancel long running operations. The db parameter is the
// database implementation that the service should use to perform queries.
// The service should be created in the api package and passed to the handlers.
// example:
//
//	func main() {
//		db := db.NewInMemoryDb()
//		service := services.NewCarrierService(context.Background(), db)
//		api := api.NewApi(service)
//		api.Start()
//	}
func NewCarrierService(ctx context.Context, db db.CarrierDb) CarrierService {
	return &carrierService{
		ctx: &ctx,
		db:  db,
	}
}

// CreateCarrier Creates a new carrier and returns the created carrier. If the carrier
// could not be created, an error is returned.
func (s *carrierService) CreateCarrier(schema *schemas.CreateCarrierSchema) (*schemas.CarrierSchema, error) {
	model := models.CreateFreightCarrier{
		Name: schema.Name,
		DBA:  schema.DBA,
	}
	carrier, err := s.db.CreateCarrier(s.ctx, &model)
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
func (s *carrierService) GetCarrier(id string) (*schemas.CarrierSchema, error) {
	carrier, err := s.db.GetCarrier(s.ctx, id)
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
func (s *carrierService) GetCarriers() ([]*schemas.CarrierSchema, error) {
	carriers, err := s.db.GetCarriers(s.ctx)
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
func (s *carrierService) UpdateCarrier(schema *schemas.CarrierSchema) (*schemas.CarrierSchema, error) {
	model := models.FreightCarrier{
		ID: schema.ID,
		CreateFreightCarrier: models.CreateFreightCarrier{
			Name: schema.Name,
			DBA:  schema.DBA,
		},
	}
	carrier, err := s.db.UpdateCarrier(s.ctx, &model)
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
func (s *carrierService) DeleteCarrier(id string) error {
	return s.db.DeleteCarrier(s.ctx, id)
}
