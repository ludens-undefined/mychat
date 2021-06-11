package admin

import (
	"gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(gc *gin.Context) {
	//account := gc.Param("account")
	//password := gc.Param("password")
	paramId := gc.PostForm("id")
	id, _ := strconv.Atoi(paramId)
	v, err := model.BusinessModel().GetBusinessById(id)

	if err == nil {
		gc.JSON(200, gin.H{
			"message": v,
		})
	} else {
		gc.JSON(400, gin.H{
			"message": err,
		})
	}
}
