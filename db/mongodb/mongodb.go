package mongodb

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CarrierResourceManagerContextKey string

const (
	// ContextKey used to fetch or put the Carrier Resource Manager into the context
	ContextKey CarrierResourceManagerContextKey = "carrierResourceManagerContextKey"
)

type resourceManager struct {
	session mongo.SessionContext
}

// Get implements db.CarrierResourceManager.
func (r *resourceManager) Get(query *db.CarrierQuery) ([]*models.Carrier, error) {
	coll := r.session.Client().Database("carriers").Collection("carriers")

	projection := bson.D{}

	// see https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/
	for _, fieldName := range query.Fields {
		// for security reasons we only want people to be able to query the objects that they should be able to
		if slices.Contains([]string{"id", "firstName", "lastName"}, fieldName) {
			projection = append(projection, bson.E{
				Key:   fieldName,
				Value: 1,
			})
		}
	}
	if len(query.SortBy) != 0 {
		if !slices.Contains([]string{"_id", "id"}, query.SortBy) {
			return nil, fmt.Errorf("%s is not a valid sortBy option", query.SortBy)
		}
	}
	sort := bson.D{bson.E{Key: query.SortBy, Value: 1}}
	opts := options.Find().
		SetSort(sort).
		SetLimit(int64(query.PageSize)).
		SetSkip(int64((query.Page) * query.PageSize)).
		SetProjection(projection)

	cursor, err := coll.Find(r.session, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	results := []*models.Carrier{}
	for cursor.Next(r.session) {
		var result models.Carrier
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf("Error occured fetching record %s\n", err.Error())
			continue
		}
		results = append(results, &result)
	}
	return results, nil

}

// WithContext fetches the mongo db session context from that passed argument (parent context)
// ,appends the carrier manager and returns all with the new context.
func WithContext(session mongo.SessionContext) context.Context {
	if session == nil {
		panic("Could not fetch session from context")
	}
	mgr := NewCarrierManager(session)
	return context.WithValue(session, ContextKey, mgr)
}

// FromContext gets the Resource Manager from the context passsed.
func FromContext(ctx context.Context) db.CarrierResourceManager {
	val := ctx.Value(ContextKey)
	if val == nil {
		panic(errors.New("could not fetch CarrierResourceManager from context"))
	}

	return val.(*resourceManager)
}

// CreateCarrier implements db.CarrierResourceManager.
func (r *resourceManager) CreateCarrier(carrier models.Carrier) (interface{}, error) {
	insertedResult, err := r.collection().InsertOne(r.session,
		&bson.M{
			"firstName": carrier.FirstName,
			"lastName":  carrier.LastName,
		},
		options.InsertOne(),
	)
	if err != nil {
		return nil, err
	}
	id := insertedResult.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil
}

// DeleteCarrier implements db.CarrierResourceManager.
func (r *resourceManager) DeleteCarrier(id interface{}) error {
	coll := r.session.Client().Database("carriers").Collection("carriers")
	_, err := coll.DeleteOne(r.session, bson.M{"_id": id})
	return err
}

// GetById implements db.CarrierResourceManager.
func (r *resourceManager) GetById(id interface{}) (*models.Carrier, error) {
	var result models.Carrier
	filter := bson.M{"_id": id}
	coll := r.session.Client().Database("carriers").Collection("carriers")
	if err := coll.FindOne(r.session, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCarrier implements db.CarrierResourceManager.
func (r *resourceManager) UpdateCarrier(id interface{}, carrier models.Carrier) error {
	coll := r.session.Client().Database("carriers").Collection("carriers")
	result, err := coll.UpdateOne(r.session, bson.M{"_id": id}, carrier)

	if result.MatchedCount == 0 {
		return fmt.Errorf("could not find Carrier with id %s", id)
	}
	return err
}

func NewCarrierManager(session mongo.SessionContext) db.CarrierResourceManager {
	return &resourceManager{session: session}
}

func (r *resourceManager) collection() *mongo.Collection {
	coll := r.session.Client().Database("freightcms").Collection("carriers")
	return coll
}
