package routers

import (
	"gin/app/http/controllers/v1/admin"

	"github.com/gin-gonic/gin"
)

func adminRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", admin.Login)
}
