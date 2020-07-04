package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

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
{    "total": 15,    "system_id": 86,    "name": "RSdoFpmegIcXPTGJgqhxLvode",    "count": 52}



*/

// Materials struct is a row record of the materials table in the main database
type Materials struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] name                                           text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Name string `gorm:"column:name;type:text;" json:"name"`
	//[ 2] count                                          integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Count int `gorm:"column:count;type:integer;" json:"count"`
	//[ 3] total                                          integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Total int `gorm:"column:total;type:integer;" json:"total"`
}

var materialsTableInfo = &TableInfo{
	Name: "materials",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "system_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "integer",
			DatabaseTypePretty: "integer",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "integer",
			ColumnLength:       -1,
			GoFieldName:        "SystemID",
			GoFieldType:        "int32",
			JSONFieldName:      "system_id",
			ProtobufFieldName:  "system_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "count",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "integer",
			DatabaseTypePretty: "integer",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "integer",
			ColumnLength:       -1,
			GoFieldName:        "Count",
			GoFieldType:        "int32",
			JSONFieldName:      "count",
			ProtobufFieldName:  "count",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "total",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "integer",
			DatabaseTypePretty: "integer",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "integer",
			ColumnLength:       -1,
			GoFieldName:        "Total",
			GoFieldType:        "int32",
			JSONFieldName:      "total",
			ProtobufFieldName:  "total",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Materials) TableName() string {
	return "materials"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Materials) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Materials) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Materials) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Materials) TableInfo() *TableInfo {
	return materialsTableInfo
}
