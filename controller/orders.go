package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func formate(input vo.OrderInput) model.Orders {
	materialId,_ := json.Marshal(input.MaterialID)
	process,_ := json.Marshal(input.Process)
	deadline,_ := time.Parse("2006-01-02",input.DeadlineTime)

	m := model.Orders{
		SystemID: input.SystemID,
		CustomerName: input.CustomerName,
		FileName: input.FileName,
		Department: input.Department,
		MaterialID: string(materialId),
		MakerID: input.MakerID,
		Process: string(process),
		DeadlineTime:int(deadline.Unix()),
		OriginAmount: input.OriginAmount,
		Discount: input.Discount,
		Amount: input.OriginAmount * input.Discount,
		OrderStatus: input.OrderStatus,
		AdminStatus: input.AdminStatus,
	}

	return m
}

// 获取所有订单
func GetOrders(c *gin.Context) {

	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var rtv *[]model.Orders
	rtv, err = mapper.SelectOrders()
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "查询失败")
		return
	}

	rest.Success(c, *rtv)

}

// 获取单个订单
func GetOrder(c *gin.Context) {
	var rtv *model.Orders

	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv, err = mapper.SelectOrderById(c.Param("id"))
	if err != nil {
		clog.Error("GetOrder", err)
		rest.Error(c, "查询失败")
		return
	}

	rest.Success(c, rtv)
}

// 新增订单
func PostOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input vo.OrderInput
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	materialId,_ := json.Marshal(input.MaterialID)
	process,_ := json.Marshal(input.Process)
	deadline,_ := time.Parse("2006-01-02",input.DeadlineTime)

	m := model.Orders{
		CustomerName: input.CustomerName,
		FileName: input.FileName,
		Department: input.Department,
		MaterialID: string(materialId),
		MakerID: input.MakerID,
		Process: string(process),
		CreateTime: int(time.Now().Unix()),
		DeadlineTime:int(deadline.Unix()),
		OriginAmount: input.OriginAmount,
		Discount: input.Discount,
		Amount: input.OriginAmount * input.Discount,

	}


	if err = mapper.InsertOrder(m); err != nil {
		clog.Error("PostOrder",err)
		rest.Error(c,"添加订单失败")
		return
	}

	// TODO 与材料联动

	rest.Success(c,true)
}

// 修改订单 // 审核订单 // 修改订单完成状态
func PatchOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input model.Orders
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	clog.Info("PatchOrder", input)

	rtv,err := mapper.UpdateOrder(input)
	if err !=nil {
		clog.Error("PatchUser", err.Error())
		rest.Error(c,err)
		return
	}

	rest.Success(c,rtv)
}

// 删除订单
func DeleteOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	if err = mapper.DeleteOrder(c.Param("id")); err != nil {
		clog.Error("",err)
		rest.Error(c,"删除失败")
		return
	}

	rest.Success(c,true)
}

