package mapper

import (
	"cwm.wiki/ad-CMS/model"
	"testing"
)

func TestInsertMaterial(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	i := model.Materials{
		Name:  "木板",
		Count: 10,
	}

	if err := InsertMaterial(i); err != nil {
		t.Error(err)
	}

}

func TestSelectMaterials(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var materials *[]model.Materials

	materials,err := SelectMaterials()
	if err != nil {
		t.Error(err)
	}

	if materials != nil {
		t.Log(*materials)
	}
}

func TestSelectMaterial(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	var m *model.Materials

	m,err := SelectMaterial("1")
	if err != nil {
		t.Error(err)
	}

	if m != nil {
		t.Log(*m)
	}
}

func TestUpdateMaterial(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	u := model.Materials{
		SystemID: 1,
		Count: 20,
	}

	var m *model.Materials

	m,err := UpdateMaterial(u)
	if err != nil {
		t.Error(err)
	}

	t.Log(m)
}

func TestDeleteMaterial(t *testing.T) {
	model.InitGormWithPath("../ad.db")

	if err:= DeleteMaterial("2"); err!=nil {
		t.Error(err)
	}
}

// func TestIncMaterial(t *testing.T) {
// 	model.InitGormWithPath("../ad.db")

// 	if err := IncMaterial([]int{1,4}); err!=nil {
// 		t.Error(err)
// 	}
// }

// func TestDecMaterial(t *testing.T) {
// 	model.InitGormWithPath("../ad.db")

// 	if err := DecMaterial([]int{1,4}); err!=nil {
// 		t.Error(err)
// 	}
// }