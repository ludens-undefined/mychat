package main

import (
	_ "gin/model"
	"gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.GetRoutes(r)

	_ = r.Run("localhost:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
