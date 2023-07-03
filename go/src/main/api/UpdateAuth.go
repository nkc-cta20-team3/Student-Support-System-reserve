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
	var post int
	response := http.StatusBadRequest

	//認可ステータスの取得
	db.Table("oa").Select("status").Where("document_id = ?", request.DocumentID).Scan(&documentStatus)

	//役職IDの取得
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)

	log.Println(request)
	log.Println(documentStatus)
	log.Println(post)

	if post == documentStatus+1 && documentStatus >= 0 {

		if post == 1 {

			//担任が認可するとき
			//所属クラス

			//認可書類の提出者と認可者が学生・担任関係にあるかの確認
			//var classID int
			//db.Table("")

			db.Table("oa").
				Where("document_id = ?", request.DocumentID).
				Updates(model.UpdateDocument{Status: post, TeacherComment: request.TeacherComment})

		} else if post >= 2 {

			//教員が認可する
			db.Table("oa").
				Where("document_id = ?", request.DocumentID).
				Updates(model.UpdateDocument{Status: post})

		}

		//エラーハンドリング
		if db.Error != nil {
			errMsg := "UPDATE ERROR"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		response = http.StatusOK

	} else {
		log.Print("認可書類が適切でない、または認可者が適切ではありません")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
