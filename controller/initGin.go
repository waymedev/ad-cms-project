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


	//r.GET("/index/*url", func(c *gin.Context) {
	//	pre, exist := c.Params.Get("url")
	//	if !exist {
	//		return
	//	}
	//	c.Redirect(http.StatusMovedPermanently, "/index/"+pre)
	//})
	//r.GET("/favicon.ico", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/index/favicon.ico")
	//})
	//r.GET("/index.html", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/index")
	//})

	// router
	api := r.Group("/api")
	{
		api.POST("/login", LoginController)
		api.GET("/user", GetUsers)
		api.GET("/user/:id", GetUser)
		api.PATCH("/user", PatchUser)
		api.DELETE("/user/:id", DeleteUser)
		api.POST("/user", PostUser)

		// 订单管理
		api.GET("/order", GetOrders)
		api.GET("/order/:id", GetOrder)
		api.DELETE("/order/:id", DeleteOrder)
		api.POST("/order", PostOrder)
		api.PATCH("/order", PatchOrder)
		api.POST("/order/search",OrderSearch)
		////// 导出
		api.GET("/download",GetAllDownload)
		api.GET("/download/:id",GetDownloadById)

		// 材料管理
		api.GET("/m", GetMaterials)
		api.GET("/m/:id", GetMaterial)
		api.POST("/m", PostMaterial)
		api.PATCH("/m", PatchMaterial)
		api.DELETE("/m/:id", DeleteMaterial)

		// 绩效管理
		api.GET("/eff/:id", GETEffective)
		//api.PATCH("/order/admin", PatchAdmin)
		api.PATCH("/order/status", PatchStatus)

		// 资金管理
		api.GET("/fund", GetFunds)
		api.GET("/fund/:id", GetFund)
		api.PATCH("/fund", PatchFund)
		api.DELETE("/fund/:id", DeleteFund)
		api.POST("/fund", PostFund)

	}
	if err := r.Run(":" + port); err != nil {
		return
	}

}

