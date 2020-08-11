package mapper

import (
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"testing"
	"time"
)

func TestInsertFund(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	i := model.Funds{
		Name: "材料名称",
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

func TestTime(t *testing.T) {

	time,_ := time.Parse("2006-01-02", "2019-01-02")
	t.Log(time.Unix())
}

func TestSelectFundsByFileter(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	search := vo.SearchFund{
		Start:1596988800,
		End:1597161600,
	}
	funds,_ := SelectFundsByFileter(search)
	t.Log(funds)

}