package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTest(api *gin.RouterGroup) {
	test := api.Group("test")
	{
		test.GET("", h.test1)
	}
}
func (h *Handler) test1(c *gin.Context) {
	h.service.LotteryAdmin.Get(c)
	log.Println("sss")
	c.JSON(200, "200")
}
