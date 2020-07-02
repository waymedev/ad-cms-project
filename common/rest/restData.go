package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context,m interface{}) {
	c.JSON(http.StatusOK,gin.H{
		"code" : 0,
		"data" : m,
	})
}

func Error(c *gin.Context,m interface{}) {
	c.JSON(http.StatusBadRequest,gin.H{
		"code" : 1,
		"msg" : m,
	})
}