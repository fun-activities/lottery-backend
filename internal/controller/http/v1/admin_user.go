package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminUser(api *gin.RouterGroup) {
	user := api.Group("user")
	{
		user.POST("login", h.adminLogin)
		user.GET("info", h.adminInfo)
		user.GET("logout", h.adminLogout)
	}

}
func (h *Handler) adminLogin(c *gin.Context) {
	c.JSON(200, gin.H{"code": 10000, "data": map[string]string{
		"token": "admin-token",
	}})
}

func (h *Handler) adminInfo(c *gin.Context) {
	c.JSON(200, gin.H{"code": 10000, "data": map[string]interface{}{
		"roles":        []string{"admin"},
		"token":        "admin-token",
		"introduction": "I am a super administrator",
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"name":         "Super Admin",
	}})
}

func (h *Handler) adminLogout(c *gin.Context) {
	log.Println("sss")
	c.JSON(200, "200")
}
