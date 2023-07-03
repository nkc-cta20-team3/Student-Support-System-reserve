package api

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadPosition(c *gin.Context) {
	request := model.UserData{}

	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responceDocument := model.UserPosition{}
	db := infra.DBInitGorm()

	//役職IDを取得する
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&responceDocument)

	if db.Error != nil {
		log.Println("SQL ERROR!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SQL ERROR"})
	}

	//戻り値
	c.JSON(http.StatusOK, gin.H{"message": responceDocument})
	return

}
