package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
	"github.com/gin-gonic/gin"
)

// 新增材料
func PostMaterial(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("PostMaterial", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input model.Materials
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	if err = mapper.InsertMaterial(input); err!=nil {
		clog.Error("PostMaterial", err)
		rest.Error(c,"添加失败")
		return
	}

	rest.Success(c,true)

}

// 获取材料列表
func GetMaterials(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetMaterials", err)
		rest.Error(c, "请重新登录")
		return
	}

	var rtv *[]model.Materials

	rtv,err = mapper.SelectMaterials()
	if err != nil {
		clog.Error("GetMaterials", err)
		rest.Error(c,"查询失败")
		return
	}

	rest.Success(c,rtv)
}

// 获取单个材料
func GetMaterial(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetMaterial", err)
		rest.Error(c, "请重新登录")
		return
	}

	var rtv *model.Materials

	rtv,err = mapper.SelectMaterial(c.Param("id"))
	if err != nil {
		clog.Error("GetMaterial",err)
		rest.Error(c,"查询失败")
		return
	}

	rest.Success(c,rtv)
}

// 修改材料
func PatchMaterial(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("PatchMaterial", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input model.Materials
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	rtv,err := mapper.UpdateMaterial(input)
	if err != nil {
		clog.Error("PatchMaterial", err)
		rest.Error(c,"修改失败")
		return
	}

	rest.Success(c,rtv)


}

// 删除材料
func DeleteMaterial(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("DeleteMaterial", err)
		rest.Error(c, "请重新登录")
		return
	}

	if err = mapper.DeleteMaterial(c.Param("id")); err != nil {
		clog.Error("DeleteMaterial", err)
		rest.Error(c,"删除失败")
		return
	}

	rest.Success(c,true)
}