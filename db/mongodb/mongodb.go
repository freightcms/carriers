package mongodb

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	organizationModels "github.com/freightcms/organizations/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	jsonToBsonMap map[string]string = make(map[string]string)
)

func init() {
	modelType := reflect.TypeOf(models.Carrier{})
	for i := range modelType.NumField() {
		field := modelType.Field(i)
		jsonToBsonMap[field.Tag.Get("json")] = field.Tag.Get("bson")
	}
	orgModelType := reflect.TypeOf(organizationModels.Organization{})
	for i := range orgModelType.NumField() {
		field := orgModelType.Field(i)
		jsonToBsonMap[field.Tag.Get("json")] = field.Tag.Get("bson")
	}
	delete(jsonToBsonMap, "")
}

type resourceManager struct {
	session mongo.SessionContext
}

func (r *resourceManager) getProjection(fields []string) bson.D {
	projection := bson.D{}
	// see https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/
	for jsonKey, bsonKey := range jsonToBsonMap {
		// for security reasons we only want people to be able to query the objects that they should be able to
		tmp := bson.E{
			Key:   bsonKey,
			Value: 0,
		}
		if slices.Contains(fields, jsonKey) {
			tmp.Value = 1
		}
		projection = append(projection, tmp)
	}
	return projection
}

// Get implements db.CarrierResourceManager.
func (r *resourceManager) Get(query *db.CarrierQuery) ([]*models.Carrier, int64, error) {
	if len(query.SortBy) != 0 {
		if !slices.Contains([]string{"_id", "id"}, query.SortBy) {
			return nil, 0, fmt.Errorf("%s is not a valid sortBy option", query.SortBy)
		}
	}
	projection := r.getProjection(query.Fields)
	sort := bson.D{bson.E{Key: query.SortBy, Value: 1}}
	opts := options.Find().
		SetSort(sort).
		SetLimit(int64(query.PageSize)).
		SetSkip(int64((query.Page) * query.PageSize)).
		SetProjection(projection)

	cursor, err := r.collection().Find(r.session, bson.D{}, opts)
	if err != nil {
		return nil, 0, err
	}
	results := []*models.Carrier{}
	for cursor.Next(r.session) {
		var result models.Carrier
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf("Error occured fetching Carrier record %s\n", err.Error())
			continue
		}
		results = append(results, &result)
	}

	count, err := r.collection().CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return nil, 0, err
	}
	return results, count, nil
}

// CreateCarrier implements db.CarrierResourceManager.
func (r *resourceManager) CreateCarrier(carrier models.Carrier) (any, error) {
	insertedResult, err := r.collection().InsertOne(r.session,
		carrier,
		options.InsertOne(),
	)
	if err != nil {
		return nil, err
	}
	id := insertedResult.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil
}

// DeleteCarrier implements db.CarrierResourceManager.
func (r *resourceManager) DeleteCarrier(id any) error {
	if reflect.TypeOf(id).Kind() != reflect.String {
		return fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}

	if err = r.session.StartTransaction(); err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	if _, err = r.collection().DeleteOne(r.session, filter); err != nil {
		return err
	}
	err = r.session.CommitTransaction(r.session)
	return err
}

// GetById implements db.CarrierResourceManager.
func (r *resourceManager) GetById(id any) (*models.Carrier, error) {
	var result models.Carrier

	if reflect.TypeOf(id).Kind() != reflect.String {
		return nil, fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}
	if err := r.collection().FindOne(r.session, &filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCarrier implements db.CarrierResourceManager.
func (r *resourceManager) UpdateCarrier(id any, carrier models.Carrier) error {
	if reflect.TypeOf(id).Kind() != reflect.String {
		return fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	result, err := r.collection().UpdateOne(r.session, filter, carrier)

	if result.MatchedCount == 0 {
		return fmt.Errorf("could not find Carrier with id %s", id)
	}
	return err
}

func (r *resourceManager) AddIdentifier(id any, identifier models.CarrierIdentifier) error {
	if reflect.TypeOf(id).Kind() != reflect.String {
		return fmt.Errorf("cannot use typeof %s as id parameter", reflect.TypeOf(id).String())
	}

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}

	filter := bson.D{
		{Key: "_id", Value: objectId},
	}
	update := bson.D{
		//bson.D{{"$addFields", bson.D{{
		//	"identifiers", identifier,
		//}},
		//}},
	}

	result := r.collection().FindOneAndUpdate(r.session, filter, update)

	return result.Err()
}

func NewCarrierManager(session mongo.SessionContext) db.CarrierResourceManager {
	return &resourceManager{session: session}
}

func (r *resourceManager) collection() *mongo.Collection {
	coll := r.session.Client().Database("freightcms").Collection("carriers")
	return coll
}
