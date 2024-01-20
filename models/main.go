package models

// CarrierIdentifyingCode is a string enum representing the type of carrier identification code
type CarrierIdentifyingCode string

// carrier identification codes enum
// see https://kb.freightcms.com/carriers/ for more information
const (
	IATA  CarrierIdentifyingCode = "IATA"
	ICC   CarrierIdentifyingCode = "ICC"
	IMO   CarrierIdentifyingCode = "IMO"
	MC    CarrierIdentifyingCode = "MC"
	SCAC  CarrierIdentifyingCode = "SCAC"
	USDOT CarrierIdentifyingCode = "USDOT"
)

type CarrierIdnetifycationCode struct {
	ID   string `json:"id" bson:"_id"`
	Code string `json:"code" bson:"code"`
}

type CarrierIntlAddress struct {
	// ID is the unique identifier for the address
	ID string `json:"id" bson:"_id"`
	// Address1 is the street address
	Address1 string `json:"address1" bson:"address1"`
	// Address2 is the building, etc.
	Address2 string `json:"address2" bson:"address2"`
	// Address3 is the floor, suite, etc.
	Address3 string `json:"address3" bson:"address3"`
	// Region is the state
	Region string `json:"region" bson:"region"`
	// Locality is the city
	Locality string `json:"locality" bson:"locality"`
	// ZipOrPostalCode is the zip code
	ZipOrPostalCode string `json:"zipOrPostalCode" bson:"zipOrPostalCode"`
	// Country is the country code
	Country string `json:"country" bson:"country"`
}

type CarrierContact struct {
	ID                     string `json:"id" bson:"_id"`
	Name                   string `json:"name" bson:"name"`
	Email                  string `json:"email" bson:"email"`
	Phone                  string `json:"phone" bson:"phone"`
	Fax                    string `json:"fax" bson:"fax"`
	PreferredContactMethod string `json:"preferredContactMethod" bson:"preferredContactMethod"`
	Reference              string `json:"reference" bson:"reference"`
}

type CreateFreightCarrier struct {
	Name               string                      `json:"name" bson:"name"`
	DBA                string                      `json:"dba" bson:"dba"`
	Website            string                      `json:"website" bson:"website"`
	Contact            CarrierContact              `json:"contact" bson:"contact"`
	MailingAddress     CarrierIntlAddress          `json:"mailingAddress" bson:"mailingAddress"`
	BillingAddress     CarrierIntlAddress          `json:"billingAddress" bson:"billingAddress"`
	IdentificationCode []CarrierIdnetifycationCode `json:"identificationCode" bson:"identificationCode"`
}

type FreightCarrier struct {
	CreateFreightCarrier `bson:",inline"`
	ID                   string `json:"id" bson:"_id,omitempty"`
	CreatedAtUTC         string `json:"createdAtUTC" bson:"createdAtUTC"`
	UpdatedAtUTC         string `json:"updatedAtUTC" bson:"updatedAtUTC"`
}
