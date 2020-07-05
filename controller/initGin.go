package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func InitGin(port string) {
	r := gin.New()
	// 开启服务端 log 颜色
	gin.ForceConsoleColor()
	// 跨域
	mwCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	})
	r.Use(mwCORS)

	// router
	api := r.Group("/api")
	{
		api.POST("/login", LoginController)
		api.GET("/user", GetUsers)
		api.GET("/user/:id", GetUser)
		api.PATCH("/user",PatchUser)
		api.DELETE("/user/:id",DeleteUser)
		api.POST("/user",PostUser)

		// 订单管理
		api.GET("/order",GetOrders)
		api.GET("/order/:id",GetOrder)
		api.DELETE("/order/:id",DeleteOrder)
		api.POST("/order",PostOrder)
		api.PATCH("/order",PatchOrder)

		api.POST("/demo",Demo)
	}


	if err := r.Run(":" + port); err != nil {
		return
	}


}

//func RigesterGin(port string) {
//
//	r := gin.New()
//	// 跨域
//	mwCORS := cors.New(cors.Config{
//		AllowOrigins:     []string{"*"},
//		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
//		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
//		ExposeHeaders:    []string{"Content-Type"},
//		AllowCredentials: true,
//		AllowOriginFunc: func(origin string) bool {
//			return true
//		},
//		MaxAge: 2400 * time.Hour,
//	})
//	r.Use(mwCORS)
//
//	//// 路由
//	//r.GET("/books", controller.FindBooks)
//	//r.POST("/books", controller.CreateBook)
//	//r.GET("/books/:id", controller.FindBook)
//	//r.PATCH("/books/:id", controller.UpdateBook)
//	//r.DELETE("/books/:id", controller.DeleteBooks)
//	//
//	//// fwt demo
//	//r.POST("/login", controller.LoginController)
//
//	if err := r.Run(":" + port); err != nil {
//		return
//	}
//
//}
