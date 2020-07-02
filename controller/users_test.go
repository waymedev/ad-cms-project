package controller

import (
	"cwm.wiki/ad-CMS/model"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestLoginController(t *testing.T) {
	r := gin.Default()
	model.InitGormWithPath("../ad.db")

	r.POST("/test",LoginController)

	if err := r.Run(); err!=nil {
		t.Error(err)
	}
}


func TestGetUsers(t *testing.T) {
	r := gin.Default()
	model.InitGormWithPath("../ad.db")

	r.GET("/test",GetUsers)

	if err := r.Run(); err!=nil {
		t.Error(err)
	}
}
