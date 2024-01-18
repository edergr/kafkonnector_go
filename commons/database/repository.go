package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func (r *Repository) FindOne(filter bson.M, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		log.Println("Error finding document:", err)
		return err
	}
	return nil
}

func (r *Repository) UpdateOne(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating document:", err)
		return nil, err
	}
	return result, nil
}

func (r *Repository) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, document)
	if err != nil {
		log.Println("Error inserting document:", err)
		return nil, err
	}
	return result, nil
}
