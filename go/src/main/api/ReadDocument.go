package api

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadDocument(c *gin.Context) {

	request := model.ReadDocumentRequest{}
	response := model.ReadDocumentResponse{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Print(request.DocumentID)

	//データベース接続
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	/*
		type ReadDocumentResponse struct {
			DocumentID int `json:"document_id"`
			RequestDate time.Time `json:"request_date"`
			StudentID int `json:"student_id"`
			ClassName string `json:"class_name"`
			StudentName string `json:"student_name"`
			StartDate time.Time `json:"absence_start_date"`
			StartFlame int `json:"start_flame"`
			EndDate time.Time `json:"end_date"`
			EndFlame int `json:"end_flame"`
			Location string `json:"location"`
			StudentComment string `json:"student_comment"`
			TeacherComment string `json:"teacher_comment"`
		}
	*/

	// データベースからデータを取得する
	db.Debug().Table("oa").
		Select(
			"oa.document_id",
			"oa.request_at",
			"oa.start_date",
			"oa.start_flame",
			"oa.end_date",
			"oa.end_flame",
			"oa.location",
			"oa.student_comment",
			"oa.teacher_comment",
			"user.user_uuid",
			"cs.class_name",
			"user.user_name").
		Joins("JOIN user ON oa.user_id = user.user_id").
		Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
		Where("document_id = ?", request.DocumentID).
		First(&response)
	if db.Error != nil {
		errMsg := "データベースからデータを取得できませんでした"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	log.Println(response)

	c.JSON(http.StatusOK, gin.H{"document": response})

}
