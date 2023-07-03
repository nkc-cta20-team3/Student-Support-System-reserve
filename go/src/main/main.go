package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// ローカルモジュールのインポート
	"main/api"
)

func main() {

	//
	g := gin.Default()
	g.Use(cors.New(cors.Config{

		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
		},

		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"x-requested-with",
		},

		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,

		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// ルーティング
	routes := g.Group("/api")
	{
		// apiフォルダ内のapiをルーティング

		// curl -X POST http://localhost:8080/api/ral
		routes.POST("/ral", api.ReadAuthList)

		routes.POST("/rd", api.ReadDocument)

		routes.POST("/ua", api.UpdateAuth)

		routes.POST("/rp", api.ReadPosition)

		routes.POST("/ra", api.RejectAuth)

		// 実装予定の管理者向けのAPI
		// routes.POST("/cd", api.CreateDocument)
		// routes.POST("/dd", api.DeleteDocument)

		// 動作確認用
		// curl -X GET http://localhost:8080/api/hello
		routes.GET("/hello", api.Hello)

	}

	g.Run(":8080")
}
