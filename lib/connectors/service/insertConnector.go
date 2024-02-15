package service

import (
	"errors"
	"kafkonnector_go/commons/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) InsertConncetorConfig(connectorData interface{}) (*mongo.InsertOneResult, error) {
	connector, ok := connectorData.(*database.Connector)
	if !ok {
		return nil, errors.New("connectorData is not of type *Connector")
	}

	result, err := s.Repository.InsertOne(connector)
	if err != nil {
		return nil, err
	}

	return result, nil
}
