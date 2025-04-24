package handler

import (
	"github.com/SashaVolohov/mapsRequestServer/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.PUT("/:key/:value/:lifeTime", h.createValueByKey)
		api.GET("/:key", h.getValueByKey)
		api.DELETE("/:key", h.deleteValueByKey)
	}

	return router

}
