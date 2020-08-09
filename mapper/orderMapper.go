package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"strconv"
)

// 添加订单
func InsertOrder(order model.Orders) error {
	err := model.DB.Create(&order).Error

	return err
}

//  查找全部订单
func SelectOrders() (*[]model.Orders, error) {
	var orders []model.Orders
	if err := model.DB.Find(&orders).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &orders, nil
}

func SelectOrderByFilter(s vo.SearchOrder) (*[]model.Orders, error) {

	var orders []model.Orders
	var db = model.DB
	if s.Name != "" {
		db = db.Where("customer_name like ?", "%"+s.Name+"%")
	}
	if s.Start != 0 {
		db = db.Where("create_time >= ?", s.Start)
	}
	if s.End != 0 {
		db = db.Where("create_time <= ?", s.End)
	}

	if err := db.Find(&orders).Error; err != nil {
		clog.Error("SelectOrderByFilter", err)
		return nil, err
	}

	return &orders, nil

}

func SelectOrderByMakerId(id string) (*[]model.Orders, error) {
	var orders []model.Orders
	if err := model.DB.Where("maker_id = ?", id).Find(&orders).Error; err != nil {
		clog.Error("SelectOrderByMakerId", err)
		return nil, err
	}

	return &orders, nil

}

// 根据ID查找单个订单
func SelectOrderById(id string) (*model.Orders, error) {
	var order model.Orders
	if err := model.DB.Where("system_id = ?", id).First(&order).Error; err != nil {
		clog.Error("SelectOrderById", err)
		return nil, err
	}

	return &order, nil
}

// 修改订单
func UpdateOrder(update model.Orders) (*model.Orders, error) {
	var order *model.Orders
	order, err := SelectOrderById(strconv.Itoa(update.SystemID))
	if err != nil {
		clog.Error("UpdateUser", err)
		return nil, err
	}

	//if order.AdminStatus == 1 {
	//	return nil,errors.New("订单已审核，请联系管理员")
	//}

	if order != nil {
		model.DB.Model(&order).Update(update)
	}

	return order, nil
}

// 删除订单
func DeleteOrder(id string) error {
	err := model.DB.Where("system_id = ? ", id).Delete(model.Orders{}).Error

	return err
}

// 更新审核状态
func UpdateAdmin(u model.Orders) (*model.Orders, error) {

	var order *model.Orders
	order, err := SelectOrderById(strconv.Itoa(u.SystemID))
	if err != nil {
		clog.Error("UpdateUser", err)
		return nil, err
	}

	//err = model.DB.Model(&order).Where("system_id = ?", u.SystemID).Update("admin_status", u.AdminStatus).Error

	return order, err
}

// 更新完成状态
func UpdateStatus(u model.Orders) (*model.Orders, error) {

	var order *model.Orders
	order, err := SelectOrderById(strconv.Itoa(u.SystemID))
	if err != nil {
		clog.Error("UpdateUser", err)
		return nil, err
	}

	//if order.AdminStatus == 0 {
	//	return nil,errors.New("订单未审核")
	//}

	err = model.DB.Model(&order).Where("system_id = ?", u.SystemID).Update("order_status", u.OrderStatus).Error

	return order, err
}
