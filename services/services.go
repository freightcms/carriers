package services

import (
	"errors"

	"github.com/freightcms/carriers/schemas"
)

type FreightCarrierService interface {
	CreateCarrier(carrierSchema *schemas.CreateFreightCarrierSchema) (*schemas.FreightCarrierSchema, errors.Error)
	DeleteCarrier(id interface{}) (bool, errors.Error)
	UpdateCarrier(id interface{}, carrierSchema *schemas.UpdateFreightCarrierSchema) (*schemas.FreghtCarrierSchema, errors.Error)
	GetCarrier(id interface{}) (*schemas.FreightCarrierSchema, errors.Error)
	GetCarriers(filter interface{}) errors.Error
}
