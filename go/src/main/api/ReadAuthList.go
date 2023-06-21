package api

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadAuthList(c *gin.Context) {

	request := model.ReadAuthListRequest{}
	response := []model.ReadAuthListResponse{}
	take_class_name := model.TakeClassName{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//サブクエリ(副問い合わせ)の作成

	//所属クラス名を取得
	db.Table("teachers").Select("position_id, class_name").Where("teacher_id = ?", request.TeacherID).First(&take_class_name)
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
			Where("st.class_name = ?", take_class_name.ClassName).
			Scan(&response)
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
			Scan(&response)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}

	} 

	// ステータス200と、responseを返します
	c.JSON(http.StatusOK, gin.H{"document": response})
	return

}
