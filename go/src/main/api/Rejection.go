package api

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func Rejection(c *gin.Context) {

	request := model.DocumentRejection{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//必要な変数
	var documentstatus int
	response := http.StatusBadRequest

	//却下ステータスの取得
	db.Table("absence_document").Select("status").Where("document_id = ?", request.DocumentID).Scan(&documentstatus)

	log.Println(request)
	log.Println(documentstatus)

	//却下
	db.Table("absence_document").
		Where("document_id = ?", request.DocumentID).
		Updates(model.UpdateDocument{Status: -1})

	//エラーハンドリング
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	response = http.StatusOK
	
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}