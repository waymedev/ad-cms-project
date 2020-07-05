package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"github.com/gin-gonic/gin"
)

func GETEffective(c *gin.Context) {

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GETEffective", err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv, err := mapper.SelectOrderByMakerId(c.Param("id"))
	if err != nil {
		clog.Error("GETEffective",err)
		rest.Error(c,"查找失败")
		return
	}

	rest.Success(c,rtv)
}
