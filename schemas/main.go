package schemas

import "github.com/freightcms/carriers/models"

type CarrierIdentifyingCodes struct {
	Code  models.CarrierIdentifyingCode
	Value string
}

type CreateCarrierSchema struct {
	Name string `json:"name"`
	DBA  string `json:"dba"`
}

type CarrierSchema struct {
	CreateCarrierSchema
	ID string `json:"id"`
}
