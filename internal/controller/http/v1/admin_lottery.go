package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminLottery(api *gin.RouterGroup) {
	user := api.Group("lottery")
	{
		user.GET("get", h.adminGetLottery)
		user.POST("create", h.adminCreateLottery)
		user.POST("update", h.adminUpdateLottery)
	}
}


func (h *Handler) adminGetLottery(c *gin.Context) {	
}

func (h *Handler) adminCreateLottery(c *gin.Context) {
}

func (h *Handler) adminUpdateLottery(c *gin.Context) {
	
}
