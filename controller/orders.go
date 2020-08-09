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
	_ "github.com/go-playground/locales/ka"
	"strconv"
	"time"
)

//
//func formate(input vo.OrderInput) model.Orders {
//	materialId, _ := json.Marshal(input.MaterialID)
//	process, _ := json.Marshal(input.Process)
//
//	m := model.Orders{
//		SystemID:     input.SystemID,
//		CustomerName: input.CustomerName,
//		FileName:     input.FileName,
//		Department:   input.Department,
//		MaterialID:   string(materialId),
//		MakerID:      input.MakerID,
//		Process:      string(process),
//		DeadlineTime: input.DeadlineTime,
//		OriginAmount: input.OriginAmount,
//		Discount:     input.Discount,
//		Amount:       input.OriginAmount * input.Discount,
//		OrderStatus:  input.OrderStatus,
//		AdminStatus:  input.AdminStatus,
//	}
//
//	return m
//}
//
//获取所有订单
func GetOrders(c *gin.Context) {

	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var orderData *[]model.Orders
	orderData, err = mapper.SelectOrders()
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "查询失败")
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

	rest.Success(c, rtv)

}

////
// 获取单个订单
func GetOrder(c *gin.Context) {

	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	orderData, err := mapper.SelectOrderById(c.Param("id"))
	if err != nil {
		clog.Error("GetOrder", err)
		rest.Error(c, "查询失败")
		return
	}

	var file []vo.File
	var department []string
	err = json.Unmarshal([]byte(orderData.File), &file)
	if err != nil {
		clog.Error("unmarsh erro")
	}

	err = json.Unmarshal([]byte(orderData.Department), &department)
	if err != nil {
		clog.Error("unmarsh erro")
	}

	user, err := mapper.SelectUser(strconv.Itoa(orderData.MakerID))
	if err != nil {
		clog.Error("用户不存在")
		return
	}

	order := vo.OrderOutput{
		SystemID:     orderData.SystemID,
		CustomerName: orderData.CustomerName,
		File:         file,
		Department:   department,
		Maker:        user.Username,
		Progress:      orderData.Progress,
		CreateTime:   orderData.CreateTime,
		DeadlineTime: orderData.DeadlineTime,
		OrderStatus:  orderData.OrderStatus,
		Area:         orderData.Area,
		Price:        orderData.Price,
		Sum:          orderData.Sum,
		After:        orderData.After,
		Note:         orderData.Note,
		Amount:       orderData.Amount,
	}

	rest.Success(c, order)
}

// 新增订单
func PostOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input vo.OrderInput
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	file, _ := json.Marshal(input.File)
	department, _ := json.Marshal(input.Department)

	m := model.Orders{
		CustomerName: input.CustomerName,
		File:         string(file),
		Department:   string(department),
		Area:         input.Area,
		Price:        input.Price,
		Sum:          input.Area * input.Price,
		MakerID:      input.MakerID,
		Progress:     input.Progress,
		CreateTime:   int(time.Now().Unix()),
		DeadlineTime: input.DeadlineTime,
		Amount:       input.Amount,
		Note:         input.Note,
		After:        input.After,
	}

	if err = mapper.InsertOrder(m); err != nil {
		clog.Error("PostOrder", err)
		rest.Error(c, "添加订单失败")
		return
	}

	rest.Success(c, true)
}

////
// 修改订单
func PatchOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input vo.UpdateOrder
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	clog.Info("PatchOrder", input)

	file,err := json.Marshal(input.File)
	if err != nil {
		clog.Error("序列化失败")
	}
	department,err := json.Marshal(input.Department)
	if err != nil {
		clog.Error("序列化失败")
	}


	m := model.Orders{
		SystemID:     input.SystemID,
		CustomerName: input.CustomerName,
		File:     string(file),
		Department:   string(department),
		DeadlineTime: input.DeadlineTime,
		Area: input.Area,
		Price: input.Price,
		After: input.After,
		Progress: input.Progress,
		Amount:      input.Amount,
		Note: input.Note,
	}

	rtv, err := mapper.UpdateOrder(m)
	if err != nil {
		clog.Error("PatchOrder", err.Error())
		rest.Error(c, err.Error())
		return
	}

	rest.Success(c, rtv)
}

////
// 删除订单
func DeleteOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}


	if err = mapper.DeleteOrder(c.Param("id")); err != nil {
		clog.Error("", err)
		rest.Error(c, "删除失败,订单已审核或不存在")
		return
	}

	rest.Success(c, true)
}
//
//
//func GetDownloadById(c *gin.Context) {
//
//	f := excelize.NewFile()
//
//	order, err := mapper.SelectOrderById(c.Param("id"))
//	if err != nil {
//		clog.Error("GetOrder", err)
//		rest.Error(c, "查询失败")
//		return
//	}
//
//	// 查询制作者用户名
//	makerName, err := mapper.SelectUser(strconv.Itoa(order.MakerID))
//	if err != nil {
//		clog.Error("用户名查询错误")
//		return
//	}
//
//	// 查询使用到的材质
//	byt := []byte(order.Material)
//	var material []vo.Material
//	// 修改的材料列表
//	_ = json.Unmarshal(byt, &material)
//	var materialName string
//	for k, v := range material {
//		m, err := mapper.SelectMaterial(strconv.Itoa(v.MaterialID))
//		if err != nil {
//			clog.Error("格式化错误")
//			return
//		}
//		if k == 0 {
//			materialName = m.Name
//		} else {
//			materialName = materialName + "," + m.Name
//
//		}
//	}
//
//	// 序列化制作工艺
//	byt = []byte(order.Process)
//	var progress []string
//	var progressName string
//	_ = json.Unmarshal(byt, &progress)
//	for k, v := range progress {
//		if k == 0 {
//			progressName = v
//		} else {
//			progressName = progressName + "," + v
//		}
//	}
//
//	// 格式化时间
//	tm := time.Unix(int64(order.CreateTime), 0)
//	createTime := tm.Format("2006-01-02")
//	tm = time.Unix(int64(order.DeadlineTime), 0)
//	endTime := tm.Format("2006-01-02")
//
//	var (
//		addr  string
//		sheet = "Sheet1"
//		style int
//		data  = map[int][]interface{}{
//			1: {"文件内容"},
//			2: {"单号", "客户名称", "制作人", "订单名称", "加工部门", "材质", "制作工艺", "金额", "下单日期", "出货日期"},
//			3: {
//				order.SystemID,
//				order.CustomerName,
//				makerName.Username,
//				order.FileName,
//				order.Department,
//				materialName,
//				progressName,
//				order.OriginAmount,
//				createTime,
//				endTime,
//			},
//		}
//	)
//
//	for r, row := range data {
//
//		if addr, err = excelize.JoinCellName("A", r); err != nil {
//			fmt.Println(err)
//			return
//		}
//		if err = f.SetSheetRow(sheet, addr, &row); err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//
//	// 设置列宽
//	if err = f.SetColWidth(sheet, "A", "J", 12); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 合并月份单元格
//	if err = f.MergeCell(sheet, "A1", "J1"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 设置月份单元格样式
//	if style, err = f.NewStyle(&excelize.Style{
//		Alignment: &excelize.Alignment{Horizontal: "center"},
//	}); err != nil {
//		fmt.Println("123", err)
//		return
//	}
//
//	// 设置月份单元格字体
//	if err = f.SetColStyle(sheet, "A:J", style); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
//	c.Header("Content-Transfer-Encoding", "binary")
//
//	//回写到web 流媒体 形成下载
//	_ = f.Write(c.Writer)
//
//}
//
//func GetAllDownload(c *gin.Context) {
//
//	f := MakeFile()
//
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
//	c.Header("Content-Transfer-Encoding", "binary")
//
//	//回写到web 流媒体 形成下载
//	_ = f.Write(c.Writer)
//
//}
//
//func MakeFile() *excelize.File {
//	f := excelize.NewFile()
//
//	orders, err := mapper.SelectOrders()
//	if err != nil {
//		clog.Error("GetOrder", err)
//		return nil
//	}
//
//	var (
//		addr  string
//		sheet = "Sheet1"
//		style int
//		data  = map[int][]interface{}{
//			1: {"文件内容"},
//			2: {"单号", "客户名称", "制作人", "文件名", "加工部门", "材质", "制作工艺", "金额", "下单日期", "出货日期"},
//		}
//	)
//
//	index := 3
//	for _, order := range *orders {
//		// 查询制作者用户名
//		makerName, err := mapper.SelectUser(strconv.Itoa(order.MakerID))
//		if err != nil {
//			clog.Error("用户名查询错误")
//			return nil
//		}
//
//		// 查询使用到的材质
//		byt := []byte(order.Material)
//		var material []vo.Material
//		// 修改的材料列表
//		_ = json.Unmarshal(byt, &material)
//		materialName := ""
//		for k, v := range material {
//			m, err := mapper.SelectMaterial(strconv.Itoa(v.MaterialID))
//			if err != nil {
//				clog.Error("格式化错误")
//				return nil
//			}
//			if k == 0 {
//				materialName = m.Name
//			} else {
//				materialName = materialName + "," + m.Name
//
//			}
//		}
//
//		// 序列化制作工艺
//		byt = []byte(order.Process)
//		var progress []string
//		progressName := ""
//		_ = json.Unmarshal(byt, &progress)
//		for k, v := range progress {
//			if k == 0 {
//				progressName = v
//			} else {
//				progressName = progressName + "," + v
//			}
//		}
//
//		// 格式化时间
//		tm := time.Unix(int64(order.CreateTime), 0)
//		createTime := tm.Format("2006-01-02")
//		tm = time.Unix(int64(order.DeadlineTime), 0)
//		endTime := tm.Format("2006-01-02")
//
//		data[index] = []interface{}{
//			order.SystemID,
//			order.CustomerName,
//			makerName.Username,
//			order.FileName,
//			order.Department,
//			materialName,
//			progressName,
//			order.OriginAmount,
//			createTime,
//			endTime,
//		}
//		index++
//	}
//
//	for r, row := range data {
//
//		if addr, err = excelize.JoinCellName("A", r); err != nil {
//			fmt.Println(err)
//			return nil
//		}
//		if err = f.SetSheetRow(sheet, addr, &row); err != nil {
//			fmt.Println(err)
//			return nil
//		}
//	}
//
//	// 设置列宽
//	if err = f.SetColWidth(sheet, "A", "J", 12); err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	// 合并月份单元格
//	if err = f.MergeCell(sheet, "A1", "J1"); err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	// 设置月份单元格样式
//	if style, err = f.NewStyle(&excelize.Style{
//		Alignment: &excelize.Alignment{Horizontal: "center"},
//	}); err != nil {
//		fmt.Println("123", err)
//		return nil
//	}
//
//	// 设置月份单元格字体
//	if err = f.SetColStyle(sheet, "A:J", style); err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	return f
//}
