package model

/*
DB Table Details
-------------------------------------


CREATE TABLE users
(
	system_id integer
		constraint users_pk
			primary key autoincrement,
	username text not null,
	password text not null,
	type integer not null,
	create_time integer,
	update_time integer
)

JSON Sample
-------------------------------------
{    "password": "piYPpEXKiuhEIKrkrATIaxUXn",    "type": 33,    "create_time": 81,    "update_time": 70,    "system_id": 17,    "username": "SHyhBFXkbZiccXNYixBEWCkVy"}



*/

// Users struct is a row record of the users table in the main database
type Users struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] username                                       text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Username string `gorm:"column:username;type:text;" json:"username"`
	//[ 2] password                                       text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Password string `gorm:"column:password;type:text;" json:"password"`
	//[ 3] type                                           integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Type int `gorm:"column:type;type:integer;" json:"type"`
	//[ 4] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime int `gorm:"column:create_time;type:integer;" json:"create_time"`
	//[ 5] update_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	UpdateTime int `gorm:"column:update_time;type:integer;" json:"update_time"`
}
