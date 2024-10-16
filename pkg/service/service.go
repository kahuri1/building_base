package service

import (
	"github.com/kahuri1/building_base/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	AddBuilding(building model.Building) (int, error)
	Building(filter model.Filter) (model.BuildingResponse, error)
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}
