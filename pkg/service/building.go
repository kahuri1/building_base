package service

import (
	"github.com/kahuri1/building_base/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) Building(filter model.Filter) (model.BuildingResponse, error) {

	buildings, err := s.repo.Building(filter)
	if err != nil {
		log.Errorf("failed to get building: %w", err)
	}
	return buildings, nil
}
