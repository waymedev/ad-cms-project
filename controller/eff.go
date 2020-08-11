package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

//
//type adminInput struct {
//	SystemId int `json:"system_id"`
//	AdminStatus int `json:"admin_status"`
//}
//
type statusInput struct {
	SystemId int `json:"system_id"`
	OrderStatus int `json:"order_status"`
}
//
func GETEffective(c *gin.Context) {

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GETEffective", err)
		rest.Error(c, "请重新登录")
		return
	}

	// 根据制作人查询订单
	orderData, err := mapper.SelectOrderByMakerId(c.Param("id"))
	if err != nil {
		clog.Error("GETEffective",err)
		rest.Error(c,"查找失败")
		return
	}
	var rtv []vo.OrderOutput
	for _, v := range *orderData {

		var file []vo.File
		var department []string
		err := json.Unmarshal([]byte(v.File), &file)
		if err != nil {
			clog.Error("unmarsh erro")
		}

		err = json.Unmarshal([]byte(v.Department), &department)
		if err != nil {
			clog.Error("unmarsh erro")
		}

		user, err := mapper.SelectUser(strconv.Itoa(v.MakerID))
		if err != nil {
			clog.Error("用户不存在")
			return
		}

		order := vo.OrderOutput{
			SystemID:     v.SystemID,
			CustomerName: v.CustomerName,
			File:         file,
			Department:   department,
			Maker:        user.Username,
			Progress:      v.Progress,
			CreateTime:   v.CreateTime,
			DeadlineTime: v.DeadlineTime,
			OrderStatus:  v.OrderStatus,
			Area:         v.Area,
			Price:        v.Price,
			Sum:          v.Sum,
			After:        v.After,
			Note:         v.Note,
			Amount:       v.Amount,
		}
		rtv = append(rtv, order)
	}
	rest.Success(c,rtv)
}
//
//// 检查订单是否审核 这里需要联动材料
//func PatchAdmin(c *gin.Context) {
//	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
//	if err != nil {
//		clog.Error("PatchFund", err)
//		rest.Error(c, "请重新登录")
//		return
//	}
//
//	var input adminInput
//	if err = c.ShouldBindJSON(&input); err != nil {
//		clog.Error(err)
//		return
//	}
//
//	u := model.Orders{
//		SystemID: input.SystemId,
//		AdminStatus: input.AdminStatus,
//	}
//
//	// 查询订单
//	Originorder,err := mapper.SelectOrderById(strconv.Itoa(input.SystemId))
//	if err != nil {
//		clog.Error("PatchAdmin",err)
//		rest.Error(c,"查找失败")
//		return
//	}
//
//	byt := []byte(Originorder.Material)
//	var data []vo.Material
//	// 修改的材料列表
//	_ = json.Unmarshal(byt,&data)
//
//	// 即将要变成 1 已审核 0 未审核
//	if u.AdminStatus == 1 {
//		// 检查库存
//		for _,v := range data {
//			order,err := mapper.SelectMaterial(strconv.Itoa(v.MaterialID))
//			if err != nil {
//				clog.Error("PatchAdmin",err)
//				return
//			}
//			if order.Count - v.Number < 0 {
//				message := fmt.Sprintf("%s 库存不足，请检查材料库存",order.Name)
//				rest.Error(c, message)
//				return
//			}
//		}
//	}
//
//	// 修改订单状态
//	orderData, err := mapper.UpdateAdmin(u)
//
//	if orderData != nil && orderData.AdminStatus == 1 {
//		for _,v := range data {
//			err = mapper.DecMaterial(v)
//			if err != nil {
//				clog.Error("PatchAdmin",err)
//				return
//			}
//		}
//	}else if  orderData != nil && orderData.AdminStatus == 0 {
//		for _,v := range data {
//			err = mapper.IncMaterial(v)
//			if err != nil {
//				clog.Error("PatchAdmin",err)
//				return
//			}
//		}
//	}
//
//	var progress []string
//	err = json.Unmarshal([]byte(orderData.Process),&progress)
//	if err != nil {
//		clog.Error("unmarsh erro")
//	}
//
//	order := vo.OrderOutput{
//		SystemID: orderData.SystemID,
//		CustomerName: orderData.CustomerName,
//		FileName: orderData.FileName,
//		Department: orderData.Department,
//		MakerID: orderData.MakerID,
//		Process: progress,
//		CreateTime: orderData.CreateTime,
//		DeadlineTime: orderData.DeadlineTime,
//		OrderStatus: orderData.OrderStatus,
//		AdminStatus: orderData.AdminStatus,
//		OriginAmount: orderData.OriginAmount,
//		Discount: orderData.Discount,
//		Amount: orderData.Amount,
//	}
//
//	var materialData []vo.Material
//	err = json.Unmarshal([]byte(orderData.Material),&materialData)
//	if err != nil {
//		return
//	}
//
//	var material []vo.Material
//	for _,v := range materialData {
//		materialData,err := mapper.SelectMaterial(strconv.Itoa(v.MaterialID))
//		if err != nil {
//			return
//		}
//		m := vo.Material{
//			MaterialID: materialData.SystemID,
//			Name: materialData.Name,
//			Number: v.Number,
//		}
//
//		material = append(material,m)
//	}
//
//	order.Material = material
//
//	rest.Success(c,order)
//
//}
//
func PatchStatus(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("PatchFund", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input statusInput
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	u := model.Orders{
		SystemID: input.SystemId,
		OrderStatus: input.OrderStatus,
	}


	originOrder, err := mapper.UpdateStatus(u)
	if err != nil {
		clog.Error("PatchStatus",err)
		message := fmt.Sprintf("%s %s", err, "修改失败")
		rest.Error(c,message)
		return
	}

	// 联动资金记录
	if originOrder != nil && originOrder.OrderStatus == 0 {
		fund,err := mapper.SelectFundByOrderID(strconv.Itoa(originOrder.SystemID))
		if err != nil {
			clog.Error("PatchStatus SelectFundByOrderID",err)
			return
		}

		err = mapper.DeleteFund(strconv.Itoa(fund.SystemID))
		if err != nil {
			clog.Error("PatchStatus DeleteFund",err)
			return
		}
	} else if originOrder != nil && originOrder.OrderStatus == 1 {
		i := model.Funds{
			Name: "订单完成",
			CreateTime: int(time.Now().Unix()),
			Amount: originOrder.Amount,
			OrderID: originOrder.SystemID,
			Type: 1,
		}

		err = mapper.InsertFund(i)
		if err != nil {
			clog.Error("PatchStatus InsertFund",err)
			return
		}
	}

	rest.Success(c,originOrder)



}
//
//func FormatOneData( originData *model.Orders ) (*vo.OrderOutput,error) {
//	progress := []string{}
//	err := json.Unmarshal([]byte(originData.Process),&progress)
//	if err != nil {
//		clog.Error("unmarsh erro")
//	}
//
//	order := vo.OrderOutput{
//		SystemID: originData.SystemID,
//		CustomerName: originData.CustomerName,
//		FileName: originData.FileName,
//		Department: originData.Department,
//		MakerID: originData.MakerID,
//		Process: progress,
//		CreateTime: originData.CreateTime,
//		DeadlineTime: originData.DeadlineTime,
//		OrderStatus: originData.OrderStatus,
//		AdminStatus: originData.AdminStatus,
//		OriginAmount: originData.OriginAmount,
//		Discount: originData.Discount,
//		Amount: originData.Amount,
//	}
//
//	var materialData []vo.Material
//	err = json.Unmarshal([]byte(originData.Material),&materialData)
//	if err != nil {
//		return nil,err
//	}
//
//	material := []vo.Material{}
//	for _,v := range materialData {
//		materialData,err := mapper.SelectMaterial(strconv.Itoa(v.MaterialID))
//		if err != nil {
//			return nil,err
//		}
//		m := vo.Material{
//			MaterialID: materialData.SystemID,
//			Name: materialData.Name,
//			Number: v.Number,
//		}
//
//		material = append(material,m)
//	}
//
//	order.Material = material
//
//	return &order,nil
//}