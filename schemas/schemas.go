package schemas

import locations "github.com/freightcms/locations/models"

type CreateFreightCarrierAddress struct {
	Line1       string                `json:"line1" binding:"required" validation:"max=100,min=1"`     // Street address, P.O. box, company name, c/o
	Line2       *string               `json:"line2"`                                                   // Apartment, suite, unit, building, floor, etc.
	Line3       *string               `json:"line3"`                                                   // Floor, Bin, Section of Warehouse, Port # etc.
	Local       string                `json:"local"`                                                   // City, town, village, etc.
	Region      string                `json:"region" binding:"required" validation:"max=100,min=5"`    // State, province within Country
	PostalCode  string                `json:"postalCode" binding:"required" validation:"max=11,min=5"` // Postal code
	Country     locations.CountryCode `json:"country" binding:"required"`                              // Country
	Description *string               `json:"description" validation:"max=200,min=0"`                  // Description of the address
	Attention   *string               `json:"attention" validation:"max=100,min=0"`                    // Attention of the address
	Type        locations.AddressType `json:"type" binding:"required"`                                 // Type of address, e.g. "home", "work", "billing", "shipping", "other"
}

type CreateFreightCarrier struct {
	Name string `json:"name" binding:"required"` // display name of the freight carrier
	DBA  string `json:"dba" binding:"required"`  // doing business as, references to what the actual company name may be in a legal sense.
}

type CreateIdentificationCodeModel struct {
	Code string `json:"code"` // the code of the carrier such as MC number, DOT number, SCAC, etc.
	Type string `json:"type"` // the type of the code such as MC, DOT, SCAC, etc.
}
