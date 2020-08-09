package model

/*
DB Table Details
-------------------------------------


CREATE TABLE orders
(
	system_id integer
		constraint orders_pk
			primary key autoincrement,
	customer_name text,
	file_name text,
	department text,
	material_id text,
	maker text,
	process text,
	create_time integer,
	deadline_time integer
, order_status integer, admin_status integer)

JSON Sample
-------------------------------------
{    "department": "OpHuuLvCVJubZdAedviSlYeuV",    "maker": "IWxIepDMcNgxQJifVoIwYXceZ",    "deadline_time": 45,    "system_id": 15,    "file_name": "WHJZFdHBWKLantdhYHDBfKwZV",    "process": "DmeKfboNTtqjqdHyRhmfSQLeZ",    "create_time": 18,    "order_status": 97,    "admin_status": 66,    "customer_name": "FoWXEIldbXurTcrGDrTTLFnSM",    "material_id": "JmagLmaNPnwgvIHpyOZDebdsk"}



*/

// Orders struct is a row record of the orders table in the main database
type Orders struct {
	SystemID     int     `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	CustomerName string  `gorm:"column:customer_name;type:text;" json:"customer_name"`
	File         string  `gorm:"column:file;type:text;" json:"file"`
	Department   string  `gorm:"column:department;type:text;" json:"department"`
	MakerID      int     `gorm:"column:maker_id;type:integer;" json:"maker_id"`
	Progress      string  `gorm:"column:progress;" json:"progress"`
	CreateTime   int     `gorm:"column:create_time;type:integer;" json:"create_time"`
	DeadlineTime int     `gorm:"column:deadline_time;type:integer;" json:"deadline_time"`
	OrderStatus  int     `gorm:"column:order_status;type:integer;" json:"order_status"`
	Amount       float64 `gorm:"column:amount;type:real;" json:"amount"`
	Area         float64 `gorm:"column:area;type:real;" json:"area"`
	Price        float64 `gorm:"column:price;type:real;" json:"price"`
	Sum          float64 `gorm:"column:sum;type:real;" json:"sum"`
	After        string  `gorm:"column:after;type:text;" json:"after"`
	Note         string  `gorm:"column:note;type:text;" json:"note"`
}
