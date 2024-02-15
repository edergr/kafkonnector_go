package service

import (
	"kafkonnector_go/commons/database"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) GetConnectorsNames() ([]string, error) {
	results, err := s.Repository.Find(bson.M{}, database.Projection{IncludeFields: []string{"name"}})
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return []string{}, nil
	}

	var names []string
	for _, result := range results {
		names = append(names, result.Name)
	}

	return names, nil
}

func (s *Service) GetConnectorConfig(connectorName string) (*database.Connector, error) {
	result, err := s.Repository.FindOne(bson.M{"name": connectorName})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return &database.Connector{}, nil
	}

	return result, nil
}
