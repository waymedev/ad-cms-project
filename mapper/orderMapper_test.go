package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"encoding/json"
	"testing"
)

func TestInsertOrder(t *testing.T) {



}

func TestSelectOrders(t *testing.T) {
	//model.InitGormWithPath("../ad.db")
	//
	//var orders *[]model.Orders
	//orders, err := SelectOrders()
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//t.Log(*orders)

	progress := []string{}
	err := json.Unmarshal([]byte("[\"铁皮\",\"木头\"]"),&progress)
	if err != nil {
		clog.Error("unmarsh erro")
	}

	t.Log(progress)
}

func TestSelectOrderById(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var order *model.Orders
	order, err := SelectOrderById("2")
	if err != nil {
		t.Error(err)
	}

	t.Log(order)
}

func TestUpdateOrder(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	// materialId, _ := json.Marshal([]int{1, 2, 3, 4, 5})

	i := model.Orders{
		SystemID:   2,
		// MaterialID: string(materialId),
	}

	var order *model.Orders
	order,err := UpdateOrder(i)
	if err != nil {
		t.Error(err)
	}

	t.Log(*order)

}

func TestDeleteOrder(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	if err:= DeleteOrder("1"); err != nil {
		t.Error(err)
	}
}

func TestSelectOrderByMakerId(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	rtv, err := SelectOrderByMakerId("1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(rtv)
}