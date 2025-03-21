package models

import (
	"time"

	commonModels "github.com/freightcms/common/models"
	organizationModels "github.com/freightcms/organizations/models"
)

type CarrierIdentifierType string

const (
	IATA      CarrierIdentifierType = "IATA"
	DOTNUMBER CarrierIdentifierType = "DOT"
	SCAC      CarrierIdentifierType = "SCAC"
	MCNUMBER  CarrierIdentifierType = "MC"
)

// InsuranceInfo provides an interface for serializing and deserializing
// against data fetching APIs. Supports `json` and `bson` binding.
type InsuranceInfo struct {
	PolicyProvider string    `json:"policyProvier" bson:"policyProvider"`
	PolicyNumber   string    `json:"policyNumber" bson:"policyNumber"`
	EffectiveDate  time.Time `json:"effectiveDate" bson:"effectiveDate"`
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationDate"`
	InsuranceType  string    `json:"insuranceType" bson:"insuranceType"`
	InsuredAmount  float32   `json:"insuredAmount" bson:"insuredAmount"`
}

type CarrierIdentifier struct {
	// Type should be an identifier type such as an IATA, SCAC, ALPHA Code,
	// DOTNUMBER, etc.
	Type string `json:"type" bson:"type"`
	// Value should be the
	Value string `json:"value" value:"value"`
}

// CarrierModel provides an interface for serializing and deserializing
// against data fetching APIs. Supports `json` and `bson` binding.
type Carrier struct {
	organizationModels.Organization
	// Whether the carrier is active within a network
	IsActive bool
	// Insurance provides the different coverages that the carrier may hold as
	// a way of
	// covering
	Insurance []*InsuranceInfo
	// Modes provides the modes which are supposed by the Carrier
	Modes []commonModels.TransportationMode
}
