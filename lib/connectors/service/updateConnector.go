package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) UpdateConncetorConfig(connectorName string, connectorData interface{}) (*mongo.UpdateResult, error) {
	updateData := bson.M{"$set": connectorData}
	result, err := s.Repository.UpdateOne(bson.M{"name": connectorName}, updateData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
