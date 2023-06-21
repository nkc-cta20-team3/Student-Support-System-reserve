package api

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UpdateAuth(c *gin.Context) {

	request := model.UpdateAuthRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//必要な変数を定義
	var documentStatus int
	var teacherPosition int
	response := http.StatusBadRequest

	//認可ステータスの取得
	db.Table("absence_document").Select("status").Where("document_id = ?", request.DocumentID).Scan(&documentStatus)

	//役職IDの取得
	db.Table("teachers").Select("position_id").Where("teacher_id = ?", request.TeacherID).Scan(&teacherPosition)

	log.Println(request)
	log.Println(documentStatus)
	log.Println(teacherPosition)

	if teacherPosition == documentStatus+1 {
		//認可できる状態

		db.Table("absence_document").
			Where("document_id = ?", request.DocumentID).
			Updates(model.UpdateDocument{Status: teacherPosition, TeacherComment: request.TeacherComment})

		//エラーハンドリング
		if db.Error != nil {
			errMsg := "UPDATE ERROR"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		response = http.StatusOK

	}

	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
