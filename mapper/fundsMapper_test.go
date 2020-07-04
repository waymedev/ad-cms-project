package mapper

import (
	"cwm.wiki/ad-CMS/model"
	"testing"
	"time"
)

func TestInsertFund(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	i := model.Funds{
		Name: "材料进货",
		CreateTime: int(time.Now().Unix()),
		Amount: 123.123,
		OrderID: 1,
	}

	if err:= InsertFund(i); err != nil {
		t.Error(err)
	}
}

func TestSelectFunds(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var funds *[]model.Funds
	funds,err := SelectFunds()
	if err != nil {
		t.Error(err)
	}

	if funds != nil {
		t.Log(funds)
	}

}

func TestSelectFundById(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var fund *model.Funds
	fund,err := SelectFundById("1")
	if err != nil {
		t.Error(err)
	}

	if fund!= nil {
		t.Log(fund)
	}
}

func TestUpdateFund(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var fund *model.Funds
	u := model.Funds{
		SystemID: 1,
		Amount: 78.78,
	}

	fund,err := UpdateFund(u)
	if err != nil {
		t.Error(err)
	}

	t.Log(fund)
}

func TestDeleteFund(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	if err:=DeleteFund("1"); err!= nil {
		t.Error(err)
	}

}