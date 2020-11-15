package admin

import (
	"gin/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Login(gc *gin.Context) {
	//account := gc.Param("account")
	//password := gc.Param("password")
	paramId := gc.PostForm("id")
	id, err := strconv.Atoi(paramId)
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
