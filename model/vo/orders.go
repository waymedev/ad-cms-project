package vo

type OrderInput struct {
	SystemID     int      `json:"system_id"`
	CustomerName string   `json:"customer_name"`
	FileName     string   `json:"file_name"`
	Department   string   `json:"department"`
	MaterialID   []int    `json:"material_id"`
	MakerID      int      `json:"maker_id"`
	Process      []string `json:"process"`
	DeadlineTime int   `json:"deadline_time"`
	OriginAmount float64  `json:"origin_amount"`
	Discount     float64  `json:"discount"`
	OrderStatus  int      `json:"order_status"`
	AdminStatus  int      `json:"admin_status"`
	CreateTime   int   `json:"create_time"`
}


