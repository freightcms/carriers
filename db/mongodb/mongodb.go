package mongodb

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/squishedfox/webservice-prototype/db"
	"github.com/squishedfox/webservice-prototype/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PersonResourceManagerContextKey string

const (
	// ContextKey used to fetch or put the Person Resource Manager into the context
	ContextKey PersonResourceManagerContextKey = "personResourceManagerContextKey"
)

type resourceManager struct {
	session mongo.SessionContext
}

// Get implements db.PersonResourceManager.
func (r *resourceManager) Get(query *db.PeopleQuery) ([]*models.Person, error) {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")

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
	results := []*models.Person{}
	for cursor.Next(r.session) {
		var result models.Person
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf("Error occured fetching record %s\n", err.Error())
			continue
		}
		results = append(results, &result)
	}
	return results, nil

}

// WithContext fetches the mongo db session context from that passed argument (parent context)
// ,appends the person manager and returns all with the new context.
func WithContext(session mongo.SessionContext) context.Context {
	if session == nil {
		panic("Could not fetch session from context")
	}
	mgr := NewPersonManager(session)
	return context.WithValue(session, ContextKey, mgr)
}

// FromContext gets the Resource Manager from the context passsed.
func FromContext(ctx context.Context) db.PersonResourceManager {
	val := ctx.Value(ContextKey)
	if val == nil {
		panic(errors.New("could not fetch PersonResourceManager from context"))
	}

	return val.(*resourceManager)
}

// CreatePerson implements db.PersonResourceManager.
func (r *resourceManager) CreatePerson(person models.Person) (interface{}, error) {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	insertedResult, err := coll.InsertOne(r.session,
		&person,
		options.InsertOne(),
	)
	if err != nil {
		return nil, err
	}
	return insertedResult.InsertedID, nil
}

// DeletePerson implements db.PersonResourceManager.
func (r *resourceManager) DeletePerson(id interface{}) error {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	_, err := coll.DeleteOne(r.session, bson.M{"_id": id})
	return err
}

// GetById implements db.PersonResourceManager.
func (r *resourceManager) GetById(id interface{}) (*models.Person, error) {
	var result models.Person
	filter := bson.M{"_id": id}
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	if err := coll.FindOne(r.session, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePerson implements db.PersonResourceManager.
func (r *resourceManager) UpdatePerson(id interface{}, person models.Person) error {
	coll := r.session.Client().Database("graphql_mongo_prototype").Collection("people")
	result, err := coll.UpdateOne(r.session, bson.M{"_id": id}, person)

	if result.MatchedCount == 0 {
		return fmt.Errorf("could not find Person with id %s", id)
	}
	return err
}

func NewPersonManager(session mongo.SessionContext) db.PersonResourceManager {
	return &resourceManager{session: session}
}
