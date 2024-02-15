package routes

import (
	"net/http"

	"kafkonnector_go/lib/connectors/controller"
	"kafkonnector_go/lib/connectors/service"
)

func handleGet(service *service.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		controller.Get(service, res, req)
	}
}

func handlePost(service *service.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		controller.Post(service, res, req)
	}
}

func handleDelete(service *service.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		controller.Delete(service, res, req)
	}
}

func Router(service *service.Service) {
	http.Handle("/connectors", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		switch req.Method {
		case http.MethodGet:
			handleGet(service)(res, req)
		case http.MethodPost:
			handlePost(service)(res, req)
		case http.MethodDelete:
			handleDelete(service)(res, req)
		default:
			http.NotFound(res, req)
		}
	}))
}
