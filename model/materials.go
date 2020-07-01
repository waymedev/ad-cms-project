package model

/*
DB Table Details
-------------------------------------


CREATE TABLE materials
(
	system_id integer
		constraint materials_pk
			primary key autoincrement,
	name text not null,
	count integer not null,
	total integer not null
)

JSON Sample
-------------------------------------
{    "total": 38,    "system_id": 5,    "name": "OnaePCgSEwccmTMHfsNtpKwAm",    "count": 60}



*/

// Materials struct is a row record of the materials table in the main database
type Materials struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int32 `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] name                                           text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Name string `gorm:"column:name;type:text;" json:"name"`
	//[ 2] count                                          integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Count int32 `gorm:"column:count;type:integer;" json:"count"`
	//[ 3] total                                          integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Total int32 `gorm:"column:total;type:integer;" json:"total"`
}
