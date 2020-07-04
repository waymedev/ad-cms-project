package mapper

import (
	"cwm.wiki/ad-CMS/model"
	"encoding/json"
	"testing"
	"time"
)

func TestInsertOrder(t *testing.T) {

	model.InitGormWithPath("../ad.db")

	materialId, _ := json.Marshal([]int{1, 2, 3, 4})
	process, _ := json.Marshal([]string{"铁皮", "木头"})

	i := model.Orders{
		CustomerName: "招商银行",
		FileName:     "zs.jpg",
		Department:   "铁皮部",
		MaterialID:   string(materialId),
		Process:      string(process),
		MakerID:      1,
		CreateTime:   int(time.Now().Unix()),
		DeadlineTime: int(time.Now().Unix()),
		OrderStatus:  0,
		AdminStatus:  0,
	}

	if err := InsertOrder(i); err != nil {
		t.Fatal(err)
	}

}

func TestSelectOrders(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var orders *[]model.Orders
	orders, err := SelectOrders()
	if err != nil {
		t.Error(err)
	}

	t.Log(*orders)
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

	materialId, _ := json.Marshal([]int{1, 2, 3, 4, 5})

	i := model.Orders{
		SystemID:   2,
		MaterialID: string(materialId),
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
