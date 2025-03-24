// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	// Swagger docs.
	_ "github.com/fun-activities/lottery-backend/docs"
	v1 "github.com/fun-activities/lottery-backend/internal/controller/http/v1"
	"github.com/fun-activities/lottery-backend/internal/service"
	"github.com/fun-activities/lottery-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(router *gin.Engine, s service.Service, l logger.Interface) {
	// Options
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(Cors())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	router.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	handlerV1 := v1.NewHandler(s, l)
	// Routers
	h := router.Group("/api")
	{
		handlerV1.Init(h)
	}
}
