package v1

import (
	"log"

	"github.com/fun-activities/lottery-backend/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminPrize(api *gin.RouterGroup) {
	prize := api.Group("prize")
	{
		prize.GET("get", h.adminGetPrize)
		prize.POST("create", h.adminCreatePrize)
		prize.POST("update", h.adminUpdatePrize)
		prize.POST("updatestatus", h.adminUpdateStatusPrize)
	}
}
func (h *Handler) adminGetPrize(c *gin.Context) {
	query := entity.PrizeQuery{}
	c.ShouldBindQuery(&query)
	prizes, total,err := h.service.PrizeAdmin.Get(c, query)
	if err != nil {
		c.JSON(200, gin.H{"code": 40001, "msg": err.Error()})
		return
	}
	if len(prizes) == 0 {
		c.JSON(200, gin.H{"code": 20001, "msg": "it is empty"})
		return
	}
	c.JSON(200, gin.H{"code": 10000, "msg": "success", "data": map[string]interface{}{
		"list":  prizes,
		"total": total,
	}})
}

func (h *Handler) adminCreatePrize(c *gin.Context) {
	prizeEntity := entity.Prize{}
	c.ShouldBindJSON(&prizeEntity)
	if id, err := h.service.PrizeAdmin.Create(c, prizeEntity); id > 0 && err == nil {
		c.JSON(200, gin.H{"code": 10000})
	} else {
		c.JSON(200, gin.H{"code": 40001, "msg": err.Error()})
	}
}

func (h *Handler) adminUpdatePrize(c *gin.Context) {
	prizeEntity := entity.Prize{}
	c.ShouldBindJSON(&prizeEntity)
	if len(prizeEntity.PrizeId) == 0 {
		c.JSON(200, gin.H{"code": 30001, "msg": "prize_id param is empty"})
		return
	}
	if ok, err := h.service.PrizeAdmin.Update(c, prizeEntity); ok > 0 && err == nil {
		c.JSON(200, gin.H{"code": 10000})
	} else {
		log.Println(ok, err)
		c.JSON(200, gin.H{"code": 40001})
	}
}

func (h *Handler) adminUpdateStatusPrize(c *gin.Context) {
	prizeUpdateStatus := entity.PrizeUpdateStatus{}
	c.ShouldBindQuery(&prizeUpdateStatus)
	if len(prizeUpdateStatus.PrizeId) == 0 {
		c.JSON(200, gin.H{"code": 30001, "msg": "prize_id param is empty"})
		return
	}
	if ok, err := h.service.PrizeAdmin.UpdateStatus(c, prizeUpdateStatus.PrizeId,
		entity.Status(prizeUpdateStatus.Status)); ok > 0 && err == nil {
		c.JSON(200, gin.H{"code": 10000})
	} else {
		c.JSON(200, gin.H{"code": 40001})
	}
}
