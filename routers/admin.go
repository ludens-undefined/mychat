package routers

import (
	"gin/controllers/v1/admin"

	"github.com/gin-gonic/gin"
)

func adminRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", admin.Login)
}
