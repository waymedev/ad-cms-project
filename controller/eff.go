package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
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

	rtv, err := mapper.SelectOrderByMakerId(c.Param("id"))
	if err != nil {
		clog.Error("GETEffective",err)
		rest.Error(c,"查找失败")
		return
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
	order,err := mapper.SelectOrderById(strconv.Itoa(input.SystemId))
	if err != nil {
		clog.Error("PatchAdmin",err)
		rest.Error(c,"查找失败")
		return
	}

	byt := []byte(order.MaterialID)
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
	rtv, err := mapper.UpdateAdmin(u)

	if rtv != nil && rtv.AdminStatus == 1 {
		err = mapper.DecMaterial(data)
		if err != nil {
			clog.Error("PatchAdmin",err)
			return
		}
	}else if  rtv != nil && rtv.AdminStatus == 0 {
		err = mapper.IncMaterial(data)
		if err != nil {
			clog.Error("PatchAdmin",err)
			return
		}
	}

	rest.Success(c,rtv)

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

	rtv, err := mapper.UpdateStatus(u)
	if err != nil {
		clog.Error("PatchStatus",err)
		rest.Error(c,"修改失败")
	}

	// 联动资金记录
	if rtv != nil && rtv.OrderStatus == 0 {
		fund,err := mapper.SelectFundByOrderID(strconv.Itoa(rtv.SystemID))
		if err != nil {
			clog.Error("PatchStatus SelectFundByOrderID",err)
			return
		}

		err = mapper.DeleteFund(strconv.Itoa(fund.SystemID))
		if err != nil {
			clog.Error("PatchStatus DeleteFund",err)
			return
		}
	} else if rtv != nil && rtv.OrderStatus == 1 {
		i := model.Funds{
			Name: "订单完成",
			CreateTime: int(time.Now().Unix()),
			Amount: rtv.Amount,
			OrderID: rtv.SystemID,
		}

		err = mapper.InsertFund(i)
		if err != nil {
			clog.Error("PatchStatus InsertFund",err)
			return
		}
	}

	rest.Success(c,rtv)



}