package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/building_base/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) AddBuilding(c *gin.Context) {
	var building model.Building

	d, err := c.GetRawData()
	if err != nil {
	}
	err = json.Unmarshal(d, &building)
	if err != nil {
		log.Errorf("unmarshal handlerError")
	}

	id, err := h.service.AddBuilding(building)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
