package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"kafkonnector_go/commons/database"
	"kafkonnector_go/lib/connectors/service"
)

func Get(service *service.Service, res http.ResponseWriter, req *http.Request) {
	var requestBody *database.Connector

	var err error
	if req.ContentLength != 0 {
		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			log.Println("Failed to parse JSON request body:", err)
			http.Error(res, "Failed to parse JSON request body", http.StatusBadRequest)
			return
		}
	}

	var connector string = ""
	if requestBody != nil {
		connector = requestBody.Name
	}

	var connectorData interface{}
	if connector != "" {
		connectorData, err = service.GetConnectorConfig(connector)
	} else {
		connectorData, err = service.GetConnectorsNames()
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if connectorData != nil {
		json.NewEncoder(res).Encode(connectorData)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

func Post(service *service.Service, res http.ResponseWriter, req *http.Request) {
	var requestBody *database.Connector

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(res, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	connector := requestBody.Name

	var connectorData *database.Connector
	connectorData, err = service.GetConnectorConfig(connector)

	if !reflect.DeepEqual(connectorData, &database.Connector{}) {
		service.UpdateConncetorConfig(connector, requestBody)
	} else {
		service.InsertConncetorConfig(requestBody)
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Delete(service *service.Service, res http.ResponseWriter, req *http.Request) {
	var requestBody *database.Connector
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(res, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	connector := requestBody.Name

	service.DeleteConnectorConfig(connector)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
