package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/building_base/model"
	"net/http"
)

// Building returns a list of buildings based on filters.
// @Summary Get buildings by filters
// @Description Returns a list of buildings based on the specified filters (city, year, floors).
// @Tags Buildings
// @Accept  json
// @Produce  json
// @Param City query string false "City to filter"
// @Param Year query string false "Year to filter"
// @Param Floors query string false "Number of floors to filter"
// @Success 200 {object} map[string]interface{} "List of buildings"
// @Failure 400 {object} model.ErrorResponse "Error when receiving buildings"
// @Router /building [get]
func (h *Handler) Building(c *gin.Context) {

	city := c.Query("City")
	year := c.Query("Year")
	floors := c.Query("Floors")

	filter, err := checkRequest(city, year, floors)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{"Invalid filters"})
		return
	}

	Buildings, err := h.service.Building(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{"Failed to retrieve buildings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Buildings": Buildings})
}

func checkRequest(city, year, floors string) (model.Filter, error) {
	var filter model.Filter

	filter.City = city

	if year != "" {
		var y int
		_, err := fmt.Sscanf(year, "%d", &y)
		if err == nil {
			filter.Year = &y
		}
	}

	if floors != "" {
		var f int
		_, err := fmt.Sscanf(floors, "%d", &f)
		if err == nil {
			filter.Floors = &f
		}
	}
	return filter, nil
}
