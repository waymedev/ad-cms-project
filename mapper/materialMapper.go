package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"strconv"
)


// 添加材料
func InsertMaterial(m model.Materials) error {
	err := model.DB.Create(&m).Error

	return err
}

// 查询所有材料
func SelectMaterials() (*[]model.Materials, error) {
	var materials []model.Materials
	if err := model.DB.Find(&materials).Error; err != nil {
		clog.Error("SelectMaterials", err)
		return nil, err
	}

	return &materials, nil
}

// 查询单个材料
func SelectMaterial(id string) (*model.Materials, error) {
	var material model.Materials
	if err := model.DB.Where("system_id = ?", id).First(&material).Error; err != nil {
		clog.Error("SelectMaterial", err)
		return nil, err
	}

	return &material, nil
}

// 修改材料记录
func UpdateMaterial(u model.Materials) (*model.Materials, error) {
	var material *model.Materials
	material, err := SelectMaterial(strconv.Itoa(u.SystemID))
	if err != nil {
		clog.Error("UpdateUser", err)
		return nil, err
	}

	if material != nil {
		model.DB.Model(&material).Update(u)
	}

	return material, nil
}

// 删除材料
func DeleteMaterial(id string) error {
	err := model.DB.Where("system_id = ?", id).Delete(model.Materials{}).Error

	return err
}