package service

import (
	"github.com/kahuri1/building_base/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) AddBuilding(building model.Building) (int, error) {
	id, err := s.repo.AddBuilding(building)
	if err != nil {
		log.Errorf("failed to create name group: %w", err)
	}
	return id, nil
}
