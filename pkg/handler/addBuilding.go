package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/building_base/model"
	log "github.com/sirupsen/logrus"
)

// AddBuilding добавляет новое здание в систему.
// @Summary Add Building
// @Description add building in db
// @Tags Buildings
// @Accept  json
// @Produce  json
// @Param building body model.AddBuildingResponse true "Building data"
// @Success 200 {object} map[string]int "ID new building"
// @Failure 400 {object} model.ErrorResponse "failed validation"
// @Failure 500 {object} model.ErrorResponse "failed server"
// @Router /building [post]
func (h *Handler) AddBuilding(c *gin.Context) {
	var building model.Building

	d, err := c.GetRawData()
	if err != nil {
		c.JSON(400, model.ErrorResponse{"Failed to read request data"})
		return
	}
	err = json.Unmarshal(d, &building)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		c.JSON(400, model.ErrorResponse{"Invalid data format"})
		return
	}

	id, err := h.service.AddBuilding(building)
	if err != nil {
		log.Printf("Failed to create building: %v", err)
		c.JSON(500, model.ErrorResponse{"Failed to create building"})
		return
	}
	c.JSON(200, gin.H{"id": id})
}
