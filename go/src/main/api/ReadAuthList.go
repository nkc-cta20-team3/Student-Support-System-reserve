package api

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadAuthList(c *gin.Context) {

	response := []model.ReadAuthListResponse{}
	take_class_id := model.TakeClassID{}

	//引数用・引数がなくなったら消す
	request := model.ReadAuthListRequest{}
	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	// 本来ならミドルウェアの認証処理を挟み、取得したIDを格納する
	//id := 2

	//サブクエリ(副問い合わせ)の作成
	//所属クラス名を取得
	db.Table("user").Select("post_id, class_id").Where("user_id = ?", request.UserID).First(&take_class_id)
	// 実行予定のSQL
	// SELECT post_id, class_id FROM user WHERE user_id = 2;

	post := take_class_id.PostID - 1

	if take_class_id.PostID == 1 {

		//担任教員であるとき
		//担任クラスのみ出力する
		db.Table("oa").
			Select(
				"cs.class_name",
				"user.user_name",
				"dv.division_name",
				"oa.document_id").
			Joins("JOIN user ON oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", post).
			Where("cs.class_id = ?", take_class_id.ClassID).
			Scan(&response)
		// 実行予定のSQL
		// SELECT cs.class_name, user.user_name, dv.division_name, oa.document_id FROM oa JOIN user ON oa.user_id = user.user_id JOIN division AS dv ON oa.division_id = dv.division_id JOIN classification AS cs ON user.class_id = cs.class_id WHERE oa.status = 1 AND cs.class_id = 117;
		if db.Error != nil {
			fmt.Print("ERROR!")
		}

	} else if post >= 2 && post <= 5 {

		//主任以上の場合
		//すべてのクラスを出力する
		db.Table("oa").
			Select(
				"cs.class_name",
				"user.user_name",
				"dv.division_name",
				"oa.document_id").
			Joins("JOIN user ON oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", post).
			Scan(&response)
		// 実行予定のSQL
		// SELECT cs.class_name, user.user_name, dv.division_name, oa.document_id FROM oa JOIN user ON oa.user_id = user.user_id JOIN division AS dv ON oa.division_id = dv.division_id JOIN classification AS cs ON user.class_id = cs.class_id WHERE oa.status = 2;
		if db.Error != nil {
			fmt.Print("ERROR!")
		}
	}

	// ステータス200と、responseを返します
	c.JSON(http.StatusOK, gin.H{"document": response})
	return

}
