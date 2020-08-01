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
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] customer_name                                  text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	CustomerName string `gorm:"column:customer_name;type:text;" json:"customer_name"`
	//[ 2] file_name                                      text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	FileName string `gorm:"column:file_name;type:text;" json:"file_name"`
	//[ 3] department                                     text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Department string `gorm:"column:department;type:text;" json:"department"`
	//[ 4] material_id                                    text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Material string `gorm:"column:material;" json:"material"`
	//[ 5] maker                                          text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	MakerID int `gorm:"column:maker_id;type:text;" json:"maker_id"`
	//[ 6] process                                        text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Process string `gorm:"column:process;" json:"process"`
	//[ 7] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime int `gorm:"column:create_time;type:integer;" json:"create_time"`
	//[ 8] deadline_time                                  integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	DeadlineTime int `gorm:"column:deadline_time;type:integer;" json:"deadline_time"`
	//[ 9] order_status                                   integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	OrderStatus int `gorm:"column:order_status;type:integer;" json:"order_status"`
	//[10] admin_status                                   integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	AdminStatus  int     `gorm:"column:admin_status;type:integer;" json:"admin_status"`
	OriginAmount float64 `gorm:"column:origin_amount;type:real;" json:"origin_amount"`
	Discount     float64 `gorm:"column:discount;type:real;" json:"discount"`
	Amount       float64 `gorm:"column:amount;type:real;" json:"amount"`
}
