package api

import (
	"fmt"
	"net/http"
	"strconv"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UnAuthorizationListGet(c *gin.Context) {

	//構造体定義
	document := []model.UnAuthorizeList{}
	take_class_name := model.TakeClassName{}

	//引数を取得する
	status, err := strconv.Atoi(c.Param("teacher_data"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//DB接続
	db := infra.DBInitGorm()

	db.Table("teachers").Select("position_id, class_name").Where("teacher_id = ?", status).First(&take_class_name)
	position := take_class_name.PositionID - 1

	if take_class_name.PositionID == 1 {
		//担任教員であるとき
		//担任クラスのみ出力する

		db.Table("absence_document AS ad").
			Select(
				"st.class_name",
				"st.student_name",
				"ar.absence_category",
				"ad.document_id").
			Joins("JOIN students AS st ON ad.student_id = st.student_id").
			Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
			Where("ad.status = ?", position).
			Where("st.class_name IN (?)", take_class_name.ClassName).
			Scan(&document)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}

	} else if take_class_name.PositionID >= 2 && take_class_name.PositionID <= 5 {

		//主任以上の場合
		//すべてのクラスを出力する

		db.Table("absence_document AS ad").
			Select(
				"st.class_name",
				"st.student_name",
				"ar.absence_category",
				"ad.document_id").
			Joins("JOIN students AS st ON ad.student_id = st.student_id").
			Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
			Where("ad.status = ?", position).
			Scan(&document)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}
	} else {
		//例外のとき
		//なにもしない
	}

	payload := gin.H{
		"document": document,
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)

}
