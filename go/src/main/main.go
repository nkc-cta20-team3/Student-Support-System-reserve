package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// ローカルモジュールのインポート
	"main/api"
	"main/handler"
	"main/infra"
	"main/service"
)

func main() {

	//
	engine := infra.DBInit()
	factory := service.NewService(engine)
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

	//
	g := gin.Default()
	g.Use(service.ServiceFactoryMiddleware(factory))
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

	//
	routes := g.Group("/v1")
	{
		// user
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)

	}

	routes2 := g.Group("/v2")
	{
		// hello world
		routes2.GET("/hello", handler.Hello)
	}

	routeapi := g.Group("/api")
	{
		//apiフォルダ内のapiをルーティング
		routeapi.POST("/ral", api.ReadAuthList)
		routeapi.POST("/rd", api.ReadDocument)
		routeapi.POST("/ua", api.UpdateAuth)

		//実装予定の管理者向けのAPI
		//routeapi.POST("/cd", api.CreateDocument)
		//routeapi.POST("/dd", api.DeleteDocument)
	}

	g.Run(":8080")
}
