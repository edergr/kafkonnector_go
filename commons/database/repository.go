package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connector struct {
	Name       string `bson:"name"`
	Delimiter  string `bson:"delimiter"`
	Topic      string `bson:"topic"`
	FieldNames string `bson:"fieldNames"`
	Filters    Filter `bson:"filters"`
	Retry      *bool  `bson:"retry,omitempty"`
}

type Filter struct {
	Sequence string `bson:"sequence"`
	Jobs     []Job  `bson:"jobs"`
}

type Job struct {
	Name         string `bson:"name"`
	Type         string `bson:"type"`
	Field        string `bson:"field,omitempty"`
	Target       string `bson:"target,omitempty"`
	FirstField   string `bson:"firstField,omitempty"`
	SecondField  string `bson:"secondField,omitempty"`
	NewFieldName string `bson:"newFieldName,omitempty"`
}

type Repository struct {
	Collection *mongo.Collection
}

func ConnectorRepository(client *mongo.Client) *Repository {
	return &Repository{
		Collection: client.Database("kafkonnector").Collection("connectors"),
	}
}

func withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func handleError(operation string, err error) error {
	log.Printf("Error %s: %v\n", operation, err)
	return err
}

type Projection struct {
	IncludeFields []string `bson:"includeFields"`
	ExcludeFields []string `bson:"excludeFields"`
}

func (r *Repository) Find(filter bson.M, projection Projection) ([]Connector, error) {
	ctx, cancel := withTimeout()
	defer cancel()

	findOptions := options.Find()
	projectionFields := bson.M{}

	if len(projection.IncludeFields) > 0 {
		for _, field := range projection.IncludeFields {
			projectionFields[field] = 1
		}
	} else if len(projection.ExcludeFields) > 0 {
		for _, field := range projection.ExcludeFields {
			projectionFields[field] = 0
		}
	}

	findOptions.SetProjection(projectionFields)

	cur, err := r.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, handleError("Error finding documents", err)
	}
	defer cur.Close(ctx)

	var results []Connector
	for cur.Next(ctx) {
		var result Connector
		err := cur.Decode(&result)
		if err != nil {
			return nil, handleError("Error decoding document", err)
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return []Connector{}, nil
	}

	if err := cur.Err(); err != nil {
		return nil, handleError("Error iterating over documents", err)
	}

	return results, nil
}

func (r *Repository) FindOne(filter bson.M) (*Connector, error) {
	ctx, cancel := withTimeout()
	defer cancel()

	var result Connector
	err := r.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &Connector{}, nil
		}
		return nil, handleError("Error finding document", err)
	}
	return &result, nil
}

func (r *Repository) UpdateOne(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := withTimeout()
	defer cancel()

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, handleError("Error updating document", err)
	}
	return result, nil
}

func (r *Repository) InsertOne(connector *Connector) (*mongo.InsertOneResult, error) {
	ctx, cancel := withTimeout()
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, connector)
	if err != nil {
		return nil, handleError("Error inserting document", err)
	}
	return result, nil
}

func (r *Repository) DeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	ctx, cancel := withTimeout()
	defer cancel()

	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, handleError("Error deleting document", err)
	}
	return result, nil
}
