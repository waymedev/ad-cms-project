package model

/*
DB Table Details
-------------------------------------


CREATE TABLE funds
(
	system_id integer
		constraint funds_pk
			primary key autoincrement,
	name text not null,
	amount integer not null,
	create_time integer
)

JSON Sample
-------------------------------------
{    "create_time": 60,    "system_id": 95,    "name": "ypaSbxNXBNSClBiMHXkDyQHYT",    "amount": 39}



*/

// Funds struct is a row record of the funds table in the main database
type Funds struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] name                                           text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Name string `gorm:"column:name;type:text;" json:"name"`
	//[ 2] amount                                         integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Amount float64 `gorm:"column:amount;type:real;" json:"amount"`
	//[ 3] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime int `gorm:"column:create_time;type:integer;" json:"create_time"`

	OrderID int `gorm:"column:order_id;type:integer;" json:"order_id"`
}
