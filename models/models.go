package models

// CarrierModel provides an interface for serializing and deserializing
// against data fetching APIs. Supports `json` and `bson` binding.
type Carrier struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"firstName" bson:"firstName"`
	DBA  string `json:"lastName" bson:"lastName"`
}
