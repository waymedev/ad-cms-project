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
	"strconv"
	"time"
)

type adminInput struct {
	SystemId int `json:"system_id"`
	AdminStatus int `json:"admin_status"`
}

type statusInput struct {
	SystemId int `json:"system_id"`
	OrderStatus int `json:"order_status"`
}

func GETEffective(c *gin.Context) {

	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GETEffective", err)
		rest.Error(c, "请重新登录")
		return
	}

	orderData, err := mapper.SelectOrderByMakerId(c.Param("id"))
	if err != nil {
		clog.Error("GETEffective",err)
		rest.Error(c,"查找失败")
		return
	}
	rtv := []vo.OrderOutput{}
	for _, v := range *orderData {

		progress := []string{}
		err := json.Unmarshal([]byte(v.Process),&progress)
		if err != nil {
			clog.Error("unmarsh erro")
		}

		order := vo.OrderOutput{
			SystemID: v.SystemID,
			CustomerName: v.CustomerName,
			FileName: v.FileName,
			Department: v.Department,
			MakerID: v.MakerID,
			Process: progress,
			CreateTime: v.CreateTime,
			DeadlineTime: v.DeadlineTime,
			OrderStatus: v.OrderStatus,
			AdminStatus: v.AdminStatus,
			OriginAmount: v.OriginAmount,
			Discount: v.Discount,
			Amount: v.Amount,
		}

		materialID := []int{}
		err = json.Unmarshal([]byte(v.MaterialID),&materialID)
		if err != nil {
			return
		}

		material := []vo.Material{}
		for _,v := range materialID {
			materialData,err := mapper.SelectMaterial(strconv.Itoa(v))
			if err != nil {
				return
			}
			m := vo.Material{
				MaterialID: v,
				Name: materialData.Name,
			}

			material = append(material,m)
		}

		order.Material = material

		rtv = append(rtv, order)
	}
	rest.Success(c,rtv)
}

func PatchAdmin(c *gin.Context) {
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("PatchFund", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input adminInput
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	u := model.Orders{
		SystemID: input.SystemId,
		AdminStatus: input.AdminStatus,
	}

	// 查询订单
	Originorder,err := mapper.SelectOrderById(strconv.Itoa(input.SystemId))
	if err != nil {
		clog.Error("PatchAdmin",err)
		rest.Error(c,"查找失败")
		return
	}

	byt := []byte(Originorder.MaterialID)
	var data []int
	// 修改的材料列表
	_ = json.Unmarshal(byt,&data)

	if u.AdminStatus == 1 {
		// 检查库存
		for _,v := range data {
			order,err := mapper.SelectMaterial(strconv.Itoa(v))
			if err != nil {
				clog.Error("PatchAdmin",err)
				return
			}
			if order.Count <= 0 {
				rest.Error(c,"库存不足，请检查库存")
				return
			}
		}
	}

	// 修改订单状态
	orderData, err := mapper.UpdateAdmin(u)

	if orderData != nil && orderData.AdminStatus == 1 {
		err = mapper.DecMaterial(data)
		if err != nil {
			clog.Error("PatchAdmin",err)
			return
		}
	}else if  orderData != nil && orderData.AdminStatus == 0 {
		err = mapper.IncMaterial(data)
		if err != nil {
			clog.Error("PatchAdmin",err)
			return
		}
	}

	progress := []string{}
	err = json.Unmarshal([]byte(orderData.Process),&progress)
	if err != nil {
		clog.Error("unmarsh erro")
	}

	order := vo.OrderOutput{
		SystemID: orderData.SystemID,
		CustomerName: orderData.CustomerName,
		FileName: orderData.FileName,
		Department: orderData.Department,
		MakerID: orderData.MakerID,
		Process: progress,
		CreateTime: orderData.CreateTime,
		DeadlineTime: orderData.DeadlineTime,
		OrderStatus: orderData.OrderStatus,
		AdminStatus: orderData.AdminStatus,
		OriginAmount: orderData.OriginAmount,
		Discount: orderData.Discount,
		Amount: orderData.Amount,
	}

	materialID := []int{}
	err = json.Unmarshal([]byte(orderData.MaterialID),&materialID)
	if err != nil {
		return
	}

	material := []vo.Material{}
	for _,v := range materialID {
		materialData,err := mapper.SelectMaterial(strconv.Itoa(v))
		if err != nil {
			return
		}
		m := vo.Material{
			MaterialID: v,
			Name: materialData.Name,
		}

		material = append(material,m)
	}

	order.Material = material

	rest.Success(c,order)

}

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
		rest.Error(c,"修改失败")
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
		}

		err = mapper.InsertFund(i)
		if err != nil {
			clog.Error("PatchStatus InsertFund",err)
			return
		}
	}


	rtv,err := FormatOneData(originOrder)
	if err != nil {
		return
	}

	rest.Success(c,rtv)



}

func FormatOneData( originData *model.Orders ) (*vo.OrderOutput,error) {
	progress := []string{}
	err := json.Unmarshal([]byte(originData.Process),&progress)
	if err != nil {
		clog.Error("unmarsh erro")
	}

	order := vo.OrderOutput{
		SystemID: originData.SystemID,
		CustomerName: originData.CustomerName,
		FileName: originData.FileName,
		Department: originData.Department,
		MakerID: originData.MakerID,
		Process: progress,
		CreateTime: originData.CreateTime,
		DeadlineTime: originData.DeadlineTime,
		OrderStatus: originData.OrderStatus,
		AdminStatus: originData.AdminStatus,
		OriginAmount: originData.OriginAmount,
		Discount: originData.Discount,
		Amount: originData.Amount,
	}

	materialID := []int{}
	err = json.Unmarshal([]byte(originData.MaterialID),&materialID)
	if err != nil {
		return nil,err
	}

	material := []vo.Material{}
	for _,v := range materialID {
		materialData,err := mapper.SelectMaterial(strconv.Itoa(v))
		if err != nil {
			return nil,err
		}
		m := vo.Material{
			MaterialID: v,
			Name: materialData.Name,
		}

		material = append(material,m)
	}

	order.Material = material

	return &order,nil
}