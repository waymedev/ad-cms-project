package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"github.com/gin-gonic/gin"
	"time"
)

func LoginController(c *gin.Context) {
	var input vo.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	user, err := mapper.SelectUserByCondition(input)
	if err != nil {
		clog.Error(err)
		rest.Error(c, "用户名或密码错误")
		return
	}

	userInfo := jwt.UserInfo{
		SystemID: user.SystemID,
		Username: user.Username,
		Type:     user.Type,
	}

	token, err := jwt.GenerateToken(&userInfo, time.Hour*2)

	if err != nil {
		rest.Error(c, err)
		return
	}

	rtv := struct {
		jwt.UserInfo `json:"user"`
		Token        string `json:"token"`
	}{
		userInfo,
		token,
	}

	rest.Success(c, rtv)

}

func GetUsers(c *gin.Context) {

	var rtv *[]model.Users

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetUsers", err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv, err = mapper.SelectUsers()
	if err != nil {
		clog.Error("GetUsers", "查询出错")
		rest.Error(c, "查询用户出错")
		return
	}

	if rtv == nil {
		return
	}

	rest.Success(c, *rtv)

}

func GetUser(c *gin.Context) {

	var rtv *model.Users

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error(err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv, err = mapper.SelectUser(c.Param("id"))
	if err != nil {
		clog.Error("查询出错")
		rest.Error(c, "查询用户出错")
		return
	}

	if rtv == nil {
		return
	}
	rest.Success(c, *rtv)

}

func PatchUser(c *gin.Context) {

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetUsers", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input vo.UserInput
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	rtv,err := mapper.UpdateUser(input)
	if err !=nil {
		clog.Error("PatchUser", err.Error())
		rest.Error(c,err)
		return
	}

	rest.Success(c,rtv)

}

func DeleteUser(c *gin.Context) {

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetUsers", err)
		rest.Error(c, "请重新登录")
		return
	}

	err = mapper.DeleteUser(c.Param("id"))
	if err != nil {
		clog.Error("DeleteUser", err)
		rest.Error(c,"删除失败")
		return
	}

	rest.Success(c,true)

}

func PostUser(c *gin.Context) {

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetUsers", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input vo.UserInput
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	err = mapper.InsertUser(input)
	if err != nil {
		clog.Error("PostUser", err)
		rest.Error(c,err)
		return
	}

	rest.Success(c,true)
}

