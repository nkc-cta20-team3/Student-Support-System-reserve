package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

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

		routeapi.GET("/ad/:document_id", api.GetAbsenceDocuments)
		routeapi.GET("/aa/:absence_data", api.AuthorizeAbsence)
		routeapi.GET("/da/:absence_list", api.DeleteAbsence)
		// routeapi.GET("/ra/:absence_list", api.RegisterAbsence)
		routeapi.POST("/ual", api.UnAuthorizationList)
		routeapi.GET("/ual/:teacher_data", api.UnAuthorizationListGet)
	}

	g.Run(":8080")
}
