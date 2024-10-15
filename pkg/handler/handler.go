package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/building_base/model"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type buildingBase interface {
	AddBuilding(building model.Building) (int, error)
}

type Handler struct {
	service buildingBase
}

func Newhandler(service buildingBase) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/building", h.AddBuilding)
	return r
}
