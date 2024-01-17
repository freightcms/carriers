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
	ID   string
	Code string
}

type CarrierIntlAddress struct {
	// ID is the unique identifier for the address
	ID string
	// Address1 is the street address
	Address1 string
	// Address2 is the building, etc.
	Address2 string
	// Address3 is the floor, suite, etc.
	Address3 string
	// Region is the state
	Region string
	// Locality is the city
	Locality string
	// ZipOrPostalCode is the zip code
	ZipOrPostalCode string
	// Country is the country code
	Country string
}

type CarrierContact struct {
	ID                     string
	Name                   string
	Email                  string
	Phone                  string
	Fax                    string
	PreferredContactMethod string
	Reference              string
}

type FreightCarrier struct {
	ID                 string
	Name               string
	DBA                string
	Website            string
	Contact            CarrierContact
	MailingAddress     CarrierIntlAddress
	BillingAddress     CarrierIntlAddress
	IdentificationCode []CarrierIdnetifycationCode
	CreatedAtUTC       string
	UpdatedAtUTC       string
}
