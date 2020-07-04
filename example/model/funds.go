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
{    "system_id": 27,    "name": "DrdDLvdyeCYZilMbfMjoyNdvC",    "amount": 58,    "create_time": 33}



*/

// Funds struct is a row record of the funds table in the main database
type Funds struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int32 `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] name                                           text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Name string `gorm:"column:name;type:text;" json:"name"`
	//[ 2] amount                                         integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Amount int32 `gorm:"column:amount;type:integer;" json:"amount"`
	//[ 3] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime sql.NullInt64 `gorm:"column:create_time;type:integer;" json:"create_time"`
}

var fundsTableInfo = &TableInfo{
	Name: "funds",
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
			Name:               "amount",
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
			GoFieldName:        "Amount",
			GoFieldType:        "int32",
			JSONFieldName:      "amount",
			ProtobufFieldName:  "amount",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "create_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "integer",
			DatabaseTypePretty: "integer",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "integer",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *Funds) TableName() string {
	return "funds"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *Funds) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *Funds) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *Funds) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *Funds) TableInfo() *TableInfo {
	return fundsTableInfo
}
