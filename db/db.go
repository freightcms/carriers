package db

import "github.com/freightcms/carriers/models"

type CarrierQuery struct {
	// Page to start sorting by. Indexing at 1
	Page int
	// PageSize tells the query how many results to return based on search criteria
	PageSize int
	// SortBy Tells the query how it should be sorting the results
	SortBy string
	// Fields to include in the return statement
	Fields []string
}

// NewQuery creates a new query object the default values. This is the preferred method for creating
// a new query object. You should then set the other values of the query struct as necessary.
func NewQuery() *CarrierQuery {
	return &CarrierQuery{
		Page:     0,
		PageSize: 10,
		SortBy:   "_id",
		Fields:   []string{},
	}
}

func (q *CarrierQuery) SetPage(page int) *CarrierQuery {
	q.Page = page
	return q
}

func (q *CarrierQuery) SetPageSize(pageSize int) *CarrierQuery {
	q.PageSize = pageSize
	return q
}

func (q *CarrierQuery) SetSortBy(sortBy string) *CarrierQuery {
	q.SortBy = sortBy
	return q
}

func (q *CarrierQuery) SetFields(fields []string) *CarrierQuery {
	q.Fields = fields
	return q
}

// CarrierResourceManager provides an abstract interface for managing Carrier Resources to a database provider such as
// postgres, ms sql server, couchdb, mongodb, redis, dynamodb, etc.
type CarrierResourceManager interface { // alternatively this can be named to CarrierEntityDb or CarrierEntityManager, DbContext, etc.
	// CreateCarrier function puts a new person resource into the database and returns the ID of the newly
	// created Carrier Resource. if there is an error while attempting to create the Carrier resource it is
	// returned with a nil for the ID.
	CreateCarrier(person models.Carrier) (interface{}, error)

	// DeleteCarrier deletes a Carrier resource from the target database system. If there is an error attempting
	// to delete the resource the error is returned. If the resource does not exist no error is returned.
	DeleteCarrier(id interface{}) error

	// UpdateCarrier modifies and updates a person resource. If there is an error attempting to update the
	// resource or a resource could not be found an error is returned.
	UpdateCarrier(id interface{}, person models.Carrier) error

	// GetById fetches a Carrier resource by it's identifier. If no resource is found then nil, nil is returned
	// as a successfully "failed" attempt. If there is an issue communicating with the database system the error
	// is returned and nil for the resource.
	GetById(id interface{}) (*models.Carrier, error)

	// Get fetches all Carrier resources from target database/resource storage. If none are found an empty slice
	// is returned. If there is an error fetching one or more recrods the error is immediately returned at the
	// opperation is cancelled.
	Get(query *CarrierQuery) ([]*models.Carrier, error)

	// AddIdentifier adds a carrier identifier to the carrier associated with the id. An error is returned if the
	// carrier does not exist or there is a duplicate identifier with the carrier
	AddIdentifier(id interface{}, identifier models.CarrierIdentifier) error
	// TODO: add query availability as well so we can search for resources based on properties
}
