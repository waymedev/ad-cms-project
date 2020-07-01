package initStep

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
