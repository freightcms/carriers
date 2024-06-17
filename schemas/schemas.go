package schemas

import locations "github.com/freightcms/locations/models"

type CreateFreightCarrierAddress struct {
	Line1       string                `json:"line1" binding:"required" validate:"max=100,min=1"`     // Street address, P.O. box, company name, c/o
	Line2       *string               `json:"line2" validate:"min=0,max=100"`                        // Apartment, suite, unit, building, floor, etc.
	Line3       *string               `json:"line3" validate:"min=0,max=100"`                        // Floor, Bin, Section of Warehouse, Port # etc.
	Local       string                `json:"local" binding:"required" validate:"min=1,max=100"`     // City, town, village, etc.
	Region      string                `json:"region" binding:"required" validate:"max=100,min=5"`    // State, province within Country
	PostalCode  string                `json:"postalCode" binding:"required" validate:"max=11,min=5"` // Postal code
	Country     locations.CountryCode `json:"country" binding:"required"`                            // Country
	Description *string               `json:"description" validate:"max=200,min=0"`                  // Description of the address
	Attention   *string               `json:"attention" validate:"max=100,min=0"`                    // Attention of the address
	Type        locations.AddressType `json:"type" binding:"required"`                               // Type of address, e.g. "home", "work", "billing", "shipping", "other"
}

type CreateFreightCarrier struct {
	Name                string                                   `json:"name" binding:"required" validate:"min=1,max=100"` // display name of the freight carrier
	DBA                 string                                   `json:"dba" validate:"min=0,max=100"`                     // doing business as, references to what the actual company name may be in a legal sense.
	PhysicalAddress     CreateFreightCarrierAddress              `json:"physicalAddress" binding:"required"`
	MailingAddress      CreateFreightCarrierAddress              `json:"mailingAddress" binding:"required"`
	Insurance           []CreateFreightInsurance                 `json:"insurance" binding:"required" validate:"min=1"`
	IdentificationCodes []CreateFreightCarrierIdentificationCode `json:"identificationCodes" binding:"required" validate:"min=1"`
}

type CreateFreightCarrierIdentificationCode struct {
	Code string `json:"code" binding:"required" validate:"min=4,max=50"`
	Type string `json:"type" binding:"required" validate:"min=4,max=50"`
}

type CreateFreightInsurance struct {
	PolicyHolder   string  `json:"policyHolder" binding:"required" validate:"min=1,max=200"`   // the name of the policy holder. The person or entity which the insurance policy is issued to.
	PolicyNumber   string  `json:"policyNumber" binding:"required" validate:"min=1,max=200"`   // the policy number of the insurance with the policy holder.
	Insurer        string  `json:"insurer" binding:"required" validate:"min=1,max=200"`        // the name of the insurance company
	InsuranceType  string  `json:"insuranceType" binding:"required" validate:"min=1,max=200"`  // the type of insurance such as Cargo, Liability, etc.
	Amount         float32 `json:"amount" binding:"required" validate:"gt=0"`                  // the amount of insurance total cost coverage.
	EffectiveDate  string  `json:"effectiveDate" binding:"required" validate:"min=10,max=10"`  // the effective date of the insurance
	ExpirationDate string  `json:"expirationDate" binding:"required" validate:"min=10,max=10"` // the expiration date of the insurance
}

type CreateIdentificationCodeModel struct {
	Code string `json:"code"` // the code of the carrier such as MC number, DOT number, SCAC, etc.
	Type string `json:"type"` // the type of the code such as MC, DOT, SCAC, etc.
}
