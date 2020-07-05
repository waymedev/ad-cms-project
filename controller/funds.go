package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
	"github.com/gin-gonic/gin"
)

// 财务列表
func GetFunds(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetFunds", err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv,err := mapper.SelectFunds()
	if err != nil {
		clog.Error("GetFunds",err)
		rest.Error(c,"查询失败")
		return
	}

	rest.Success(c,rtv)
}

// 单个列表
func GetFund(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetFund", err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv,err := mapper.SelectFundById(c.Param("id"))
	if err != nil {
		clog.Error("GetFund", err)
		rest.Error(c,"查询失败")
		return
	}

	rest.Success(c,rtv)
}

// 更新财务
func PatchFund(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("PatchFund", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input model.Funds
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	rtv,err := mapper.UpdateFund(input)
	if err != nil {
		clog.Error("PatchFund",err)
		rest.Error(c,"修改失败")
		return
	}

	rest.Success(c,rtv)
}

// 删除财务
func DeleteFund(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("DeleteFund", err)
		rest.Error(c, "请重新登录")
		return
	}

	if err = mapper.DeleteFund(c.Param("id")); err!=nil{
		clog.Error("DeleteFund", err)
		rest.Error(c,"删除失败")
		return
	}

	rest.Success(c,true)

}

// 添加财务
func PostFund(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("PostFund", err)
		rest.Error(c, "请重新登录")
		return
	}
	var input model.Funds
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error("PostFund",err)
		return
	}

	if err = mapper.InsertFund(input); err!=nil {
		clog.Error("PostFund",err)
		rest.Error(c,"添加失败")
		return
	}

	rest.Success(c,true)

}