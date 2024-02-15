package service

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) DeleteConnectorConfig(connectorName string) error {
	_, err := s.Repository.DeleteOne(bson.M{"name": connectorName})
	if err != nil {
		return err
	}

	return nil
}
