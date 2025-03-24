// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/fun-activities/lottery-backend/internal/service"
	"github.com/fun-activities/lottery-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
	logger  logger.Interface
}

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewHandler(s service.Service, l logger.Interface) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		admin := v1.Group("admin") 
		{
			h.initTest(admin)
			h.initAdminPrize(admin)
			h.initAdminUser(admin)
			h.initAdminLottery(admin)
		}
		

	}

}
