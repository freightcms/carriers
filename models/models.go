package models

// PersonModel provides an interface for serializing and deserializing
// against data fetching APIs. Supports `json` and `bson` binding.
type Person struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}
