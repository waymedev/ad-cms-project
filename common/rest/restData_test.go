package rest

import (
	"cwm.wiki/ad-CMS/model/vo"
	"github.com/gin-gonic/gin"
	"testing"
)

func success(c *gin.Context) {
	m := vo.UserInput{
		Username: "admin",
		Password: "admin",
		SystemId: 123,
	}

	Success(c,m)
}

func TestSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/test",success)

	err := r.Run()
	if err != nil {
		t.Error(err)
	}

}

func errorr(c *gin.Context) {


	Error(c,"用户名或密码错误")
}

func TestError(t *testing.T) {
	r := gin.Default()
	r.GET("/test",errorr)

	err := r.Run()
	if err != nil {
		t.Error(err)
	}

}
