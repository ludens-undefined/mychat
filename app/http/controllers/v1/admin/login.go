package admin

import (
	"gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(gc *gin.Context) {
	// account := gc.Param("account")
	// password := gc.Param("password")
	paramId := gc.PostForm("id")
	// 参数转化为数字
	id, _ := strconv.Atoi(paramId)
	data, err := model.UsersModel().GetUsersById(id)

	if err != nil {
		gc.JSON(http.StatusOK, gin.H{
			"data":  data,
			"error": err,
		})
	} else {
		gc.JSON(http.StatusBadRequest, gin.H{
			"data":  data,
			"error": err,
		})
	}
}
