package controller

import (
	"cwm.wiki/ad-CMS/common/jwt"
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/common/rest"
	"cwm.wiki/ad-CMS/mapper"
	"cwm.wiki/ad-CMS/model"
	"cwm.wiki/ad-CMS/model/vo"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/locales/ka"
	"strconv"
	"time"
)

func formate(input vo.OrderInput) model.Orders {
	materialId,_ := json.Marshal(input.MaterialID)
	process,_ := json.Marshal(input.Process)

	m := model.Orders{
		SystemID: input.SystemID,
		CustomerName: input.CustomerName,
		FileName: input.FileName,
		Department: input.Department,
		MaterialID: string(materialId),
		MakerID: input.MakerID,
		Process: string(process),
		DeadlineTime:input.DeadlineTime,
		OriginAmount: input.OriginAmount,
		Discount: input.Discount,
		Amount: input.OriginAmount * input.Discount,
		OrderStatus: input.OrderStatus,
		AdminStatus: input.AdminStatus,
	}

	return m
}

// 获取所有订单
func GetOrders(c *gin.Context) {

	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var rtv *[]model.Orders
	rtv, err = mapper.SelectOrders()
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "查询失败")
		return
	}

	rest.Success(c, *rtv)

}

// 获取单个订单
func GetOrder(c *gin.Context) {
	var rtv *model.Orders

	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	rtv, err = mapper.SelectOrderById(c.Param("id"))
	if err != nil {
		clog.Error("GetOrder", err)
		rest.Error(c, "查询失败")
		return
	}

	rest.Success(c, rtv)
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

	materialId,_ := json.Marshal(input.MaterialID)
	process,_ := json.Marshal(input.Process)

	m := model.Orders{
		CustomerName: input.CustomerName,
		FileName: input.FileName,
		Department: input.Department,
		MaterialID: string(materialId),
		MakerID: input.MakerID,
		Process: string(process),
		CreateTime: int(time.Now().Unix()),
		DeadlineTime:input.DeadlineTime,
		OriginAmount: input.OriginAmount,
		Discount: input.Discount,
		Amount: input.OriginAmount * input.Discount,

	}


	if err = mapper.InsertOrder(m); err != nil {
		clog.Error("PostOrder",err)
		rest.Error(c,"添加订单失败")
		return
	}

	// TODO 与材料联动

	rest.Success(c,true)
}

// 修改订单 // 审核订单 // 修改订单完成状态
func PatchOrder(c *gin.Context) {
	// 检查 token
	_, err := jwt.ParseUser(c.GetHeader("Authorization"))
	if err != nil {
		clog.Error("GetOrders", err)
		rest.Error(c, "请重新登录")
		return
	}

	var input model.Orders
	if err = c.ShouldBindJSON(&input); err != nil {
		clog.Error(err)
		return
	}

	clog.Info("PatchOrder", input)

	rtv,err := mapper.UpdateOrder(input)
	if err !=nil {
		clog.Error("PatchUser", err.Error())
		rest.Error(c,err)
		return
	}

	rest.Success(c,rtv)
}

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
		clog.Error("",err)
		rest.Error(c,"删除失败")
		return
	}

	rest.Success(c,true)
}

func GetDownloadById(c *gin.Context) {

	f := excelize.NewFile()

	order, err := mapper.SelectOrderById(c.Param("id"))
	if err != nil {
		clog.Error("GetOrder", err)
		rest.Error(c, "查询失败")
		return
	}

	// 查询制作者用户名
	makerName,err := mapper.SelectUser(strconv.Itoa(order.MakerID))
	if err != nil {
		clog.Error("用户名查询错误")
		return
	}

	// 查询使用到的材质
	byt := []byte(order.MaterialID)
	var material []int
	// 修改的材料列表
	_ = json.Unmarshal(byt,&material)
	var materialName string
	for k,v := range material {
		m,err := mapper.SelectMaterial(strconv.Itoa(v))
		if err != nil {
			clog.Error("格式化错误")
			return
		}
		if k == 0 {
			materialName = m.Name
		} else {
			materialName = materialName + "," + m.Name

		}
	}

	// 序列化制作工艺
	byt = []byte(order.Process)
	var progress []string
	var progressName string
	_ = json.Unmarshal(byt,&progress)
	for k, v := range progress {
		if k == 0 {
			progressName = v
		}else  {
			progressName = progressName + "," + v
		}
	}


	// 格式化时间
	tm := time.Unix(int64(order.CreateTime), 0)
	createTime := tm.Format("2006-01-02")
	tm = time.Unix(int64(order.DeadlineTime), 0)
	endTime := tm.Format("2006-01-02")


	var(
		addr string
		sheet = "Sheet1"
		style int
		data = map[int][]interface{}{
			1: {"文件内容"},
			2:{"单号","客户名称","制作人","文件名","加工部门","材质","制作工艺","金额","下单日期","出货日期"},
			3: {
				order.SystemID,
				order.CustomerName,
				makerName.Username,
				order.FileName,
				order.Department,
				materialName,
				progressName,
				order.OriginAmount,
				createTime,
				endTime,
			},

		}
	)


	for r, row := range data {

		if addr, err = excelize.JoinCellName("A", r); err != nil {
			fmt.Println(err)
			return
		}
		if err = f.SetSheetRow(sheet, addr, &row); err != nil {
			fmt.Println(err)
			return
		}
	}

	// 设置列宽
	if err = f.SetColWidth(sheet, "A", "J", 12); err != nil {
		fmt.Println(err)
		return
	}

	// 合并月份单元格
	if err = f.MergeCell(sheet, "A1", "J1"); err != nil {
		fmt.Println(err)
		return
	}

	// 设置月份单元格样式
	if style, err = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
	}); err != nil {
		fmt.Println("123",err)
		return
	}

	// 设置月份单元格字体
	if err = f.SetColStyle(sheet, "A:J", style); err != nil {
		fmt.Println(err)
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	//回写到web 流媒体 形成下载
	_ = f.Write(c.Writer)



}

func GetAllDownload(c * gin.Context) {

	f := MakeFile()

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	//回写到web 流媒体 形成下载
	_ = f.Write(c.Writer)


}


func MakeFile() *excelize.File {
	f := excelize.NewFile()


	orders, err := mapper.SelectOrders()
	if err != nil {
		clog.Error("GetOrder", err)
		return nil
	}


	var(
		addr string
		sheet = "Sheet1"
		style int
		data = map[int][]interface{}{
			1: {"文件内容"},
			2:{"单号","客户名称","制作人","文件名","加工部门","材质","制作工艺","金额","下单日期","出货日期"},
		}
	)

	index := 3
	for _,order := range *orders {
		// 查询制作者用户名
		makerName,err := mapper.SelectUser(strconv.Itoa(order.MakerID))
		if err != nil {
			clog.Error("用户名查询错误")
			return nil
		}

		// 查询使用到的材质
		byt := []byte(order.MaterialID)
		var material []int
		// 修改的材料列表
		_ = json.Unmarshal(byt,&material)
		materialName := ""
		for k,v := range material {
			m,err := mapper.SelectMaterial(strconv.Itoa(v))
			if err != nil {
				clog.Error("格式化错误")
				return nil
			}
			if k == 0 {
				materialName = m.Name
			} else {
				materialName = materialName + "," + m.Name

			}
		}

		// 序列化制作工艺
		byt = []byte(order.Process)
		var progress []string
		progressName := ""
		_ = json.Unmarshal(byt,&progress)
		for k, v := range progress {
			if k == 0 {
				progressName = v
			}else  {
				progressName = progressName + "," + v
			}
		}


		// 格式化时间
		tm := time.Unix(int64(order.CreateTime), 0)
		createTime := tm.Format("2006-01-02")
		tm = time.Unix(int64(order.DeadlineTime), 0)
		endTime := tm.Format("2006-01-02")


		data[index] = []interface{}{
			order.SystemID,
			order.CustomerName,
			makerName.Username,
			order.FileName,
			order.Department,
			materialName,
			progressName,
			order.OriginAmount,
			createTime,
			endTime,
		}
		index ++
	}


	for r, row := range data {

		if addr, err = excelize.JoinCellName("A", r); err != nil {
			fmt.Println(err)
			return nil
		}
		if err = f.SetSheetRow(sheet, addr, &row); err != nil {
			fmt.Println(err)
			return nil
		}
	}

	// 设置列宽
	if err = f.SetColWidth(sheet, "A", "J", 12); err != nil {
		fmt.Println(err)
		return nil
	}

	// 合并月份单元格
	if err = f.MergeCell(sheet, "A1", "J1"); err != nil {
		fmt.Println(err)
		return nil
	}

	// 设置月份单元格样式
	if style, err = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
	}); err != nil {
		fmt.Println("123",err)
		return nil
	}

	// 设置月份单元格字体
	if err = f.SetColStyle(sheet, "A:J", style); err != nil {
		fmt.Println(err)
		return nil
	}

	return f
}