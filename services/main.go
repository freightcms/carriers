package services

import (
	"github.com/freightcms/carriers/models"
	"github.com/freightcms/carriers/schemas"
)

type CarrierDb interface {
	// GetCarrier returns a carrier by id.
	GetCarrier(id string) (*models.Carrier, error)
	// GetCarriers returns all carriers.
	GetCarriers() ([]*models.Carrier, error)
	// CreateCarrier creates a new carrier.
	CreateCarrier(carrier *models.Carrier) (*models.Carrier, error)
	// UpdateCarrier updates a carrier.
	UpdateCarrier(carrier *models.Carrier) (*models.Carrier, error)
	// DeleteCarrier deletes a carrier.
	DeleteCarrier(id string) error
}

// CarrierService is the interface that provides carrier methods.
type CarrierService struct {
	db CarrierDb
}

func NewCarrierService(db CarrierDb) *CarrierService {
	return &CarrierService{db: db}
}

// CreateCarrier Creates a new carrier and returns the created carrier. If the carrier
// could not be created, an error is returned.
func (s *CarrierService) CreateCarrier(schema *schemas.CreateCarrierSchema) (*schemas.CarrierSchema, error) {
	model := models.Carrier{
		Name: schema.Name,
		DBA:  schema.DBA,
	}
	carrier, err := s.db.CreateCarrier(&model)
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
func (s *CarrierService) GetCarrier(id string) (*schemas.CarrierSchema, error) {
	carrier, err := s.db.GetCarrier(id)
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
func (s *CarrierService) GetCarriers() ([]*schemas.CarrierSchema, error) {
	carriers, err := s.db.GetCarriers()
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
func (s *CarrierService) UpdateCarrier(schema *schemas.CarrierSchema) (*schemas.CarrierSchema, error) {
	model := models.Carrier{
		ID:   schema.ID,
		Name: schema.Name,
		DBA:  schema.DBA,
	}
	carrier, err := s.db.UpdateCarrier(&model)
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
func (s *CarrierService) DeleteCarrier(id string) error {
	return s.db.DeleteCarrier(id)
}
