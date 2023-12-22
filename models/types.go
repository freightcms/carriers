package models

// carrier identification codes enum
// see https://kb.freightcms.com/carriers/ for more information
const (
	IATA = "IATA"
	ICC = "ICC"
	IMO = "IMO"
	MC = "MC"
	SCAC = "SCAC"
	USDOT = "USDOT"
)

type CarrierIdnetifycationCode struct {
	id string
	code string
}

type CarrierIntlAddress struct {
	id string
	address1 string
	address2 string
	address3 string
	region string
	stateOrProvince string
	zipOrPostalCode string
	country string
}

type CarrierContact struct {
	id string
	name string
	email string
	phone string
	fax string
	preferredContactMethod string
	reference string
}

type FreightCarrier struct {
	id string
	name string
	dba string
	website string
	contact CarrierContact
	mailingAddress CarrierIntlAddress
	billingAddress CarrierIntlAddress
	identificationCode []CarrierIdnetifycationCode
}