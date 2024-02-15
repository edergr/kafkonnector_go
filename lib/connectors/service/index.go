package service

import (
	"kafkonnector_go/commons/database"
)

type Service struct {
	Repository *database.Repository
}

func NewService(repo *database.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
