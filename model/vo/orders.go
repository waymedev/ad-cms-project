package vo

type OrderInput struct {
	SystemID     int      `json:"system_id"`
	CustomerName string   `json:"customer_name"`
	File         []File   `json:"file"`
	Department   []string `json:"department"`
	MakerID      int      `json:"maker_id"`
	Progress     string   `json:"progress"`
	DeadlineTime int      `json:"deadline_time"`
	Area         float64  `json:"area"`
	Price        float64  `json:"price"`
	Sum          float64  `json:"sum"`
	OrderStatus  int      `json:"order_status"`
	Note         string   `json:"note"`
	CreateTime   int      `json:"create_time"`
	Amount       float64  `json:"amount"`
	After        string   `json:"after"`
}

type OrderOutput struct {
	SystemID     int      `json:"system_id"`
	CustomerName string   `json:"customer_name"`
	File         []File   `json:"file"`
	Department   []string `json:"department"`
	Maker        string   `json:"maker"`
	Progress      string   `json:"progress"`
	CreateTime   int      `json:"create_time"`
	DeadlineTime int      `json:"deadline_time"`
	OrderStatus  int      `json:"order_status"`
	Area         float64  `json:"area"`
	Price        float64  `json:"price"`
	Sum          float64  `json:"sum"`
	After        string   `json:"after"`
	Note         string   `json:"note"`
	Amount       float64  `json:"amount"`
}

type File struct {
	FileName     string `json:"file_name"`
	MaterialName string `json:"material_name"`
}

type Material struct {
	MaterialID int    `json:"material_id"`
	Name       string `json:"material_name"`
	Number     int    `json:"material_num"`
}

type UpdateOrder struct {
	SystemID     int      `json:"system_id"`
	CustomerName string   `json:"customer_name"`
	File         []File   `json:"file"`
	Department   []string `json:"department"`
	Progress      string   `json:"progress"`
	DeadlineTime int      `json:"deadline_time"`
	OrderStatus  int      `json:"order_status"`
	Area         float64  `json:"area"`
	Price        float64  `json:"price"`
	Sum          float64  `json:"sum"`
	After        string   `json:"after"`
	Note         string   `json:"note"`
	Amount       float64  `json:"amount"`
}
