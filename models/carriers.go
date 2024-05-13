package models

import (
	partyModels "github.com/freightcms/parties/models"
)

type FreightCarrierModel struct {
	ID string `json:"ID" bson:"_id"`
	partyModels.Company
	IdentificationCodes []IdentificationCodeModel      `json:"IdentificationCodes" bson:"IdentificationCodes"`
	Insurance           []FreightCarrierInsuranceModel `json:"Insurance" bson:"Insurance"`
}

type IdentificationCodeModel struct {
	Code string `json:"Code" bson:"Code"` // the code of the carrier such as MC number, DOT number, SCAC, etc.
	Type string `json:"Type" bson:"Type"` // the type of the code such as MC, DOT, SCAC, etc.
}

type FreightCarrierInsuranceModel struct {
	ID             string `json:"ID" bson:"_id"`
	PolicyHolder   string `json:"PolicyHolder" bson:"PolicyHolder"`     // the name of the policy holder. The person or entity which the insurance policy is issued to.
	PolicyNumber   string `json:"PolicyNumber" bson:"PolicyNumber"`     // the policy number of the insurance with the policy holder.
	Insurer        string `json:"Insurer" bson:"Insurer"`               // the name of the insurance company
	InsuranceType  string `json:"InsuranceType" bson:"InsuranceType"`   // the type of insurance such as Cargo, Liability, etc.
	Amount         string `json:"Amount" bson:"Amount"`                 // the amount of insurance total cost coverage
	EffectiveDate  string `json:"EffectiveDate" bson:"EffectiveDate"`   // the effective date of the insurance
	ExpirationDate string `json:"ExpirationDate" bson:"ExpirationDate"` // the expiration date of the insurance
}
