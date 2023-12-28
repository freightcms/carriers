package api

import "github.com/freightcms/carriers/models"

type CarrierIdentifyingCodes struct {
	code  models.CarrierIdentifyingCode
	value string
}

type CreateCarrierSchema struct {
	Name string `json:"name"`
	DBA  string `json:"dba"`
}
