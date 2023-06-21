package api

import (
	"net/http"
	"log"

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
	db.Debug().Table("absence_document AS ab").
		Select(
			"ab.document_id",
			"ab.request_date",
			"ab.student_id",
			"st.class_name",
			"st.student_name",
			"ab.absence_start_date",
			"ab.absence_start_flame",
			"ab.absence_end_date",
			"ab.absence_end_flame",
			"ab.location",
			"ab.student_comment",
			"ab.teacher_comment").
		Joins("JOIN students AS st ON ab.student_id = st.student_id").
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
