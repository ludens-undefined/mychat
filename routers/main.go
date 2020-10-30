package routers

import "github.com/gin-gonic/gin"

func GetRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	adminRoutes(v1)
}
