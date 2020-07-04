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


CREATE TABLE "orders"
(
	system_id integer
		constraint orders_pk
			primary key autoincrement,
	customer_name text,
	file_name text,
	department text,
	material_id text,
	maker_id integer,
	process text,
	create_time integer,
	deadline_time integer,
	order_status integer,
	admin_status integer
)

JSON Sample
-------------------------------------
{    "file_name": "MWyKZvGjvaAwJIAfWjrYoMiHP",    "deadline_time": 59,    "process": "UMnLokwGbFewZewDuTpOytFpc",    "create_time": 1,    "order_status": 27,    "system_id": 50,    "customer_name": "kEhjVVNngWgflRXHVHBVBTFTP",    "department": "IDfjXHrJwvSkmlMGKmfWtXZxa",    "material_id": "eAkNSnmEBsvsjGllWZXdMpdyh",    "maker_id": 44,    "admin_status": 32}



*/

// Orders struct is a row record of the orders table in the main database
type Orders struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int32 `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] customer_name                                  text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	CustomerName sql.NullString `gorm:"column:customer_name;type:text;" json:"customer_name"`
	//[ 2] file_name                                      text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	FileName sql.NullString `gorm:"column:file_name;type:text;" json:"file_name"`
	//[ 3] department                                     text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Department sql.NullString `gorm:"column:department;type:text;" json:"department"`
	//[ 4] material_id                                    text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	MaterialID sql.NullString `gorm:"column:material_id;type:text;" json:"material_id"`
	//[ 5] maker_id                                       integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	MakerID sql.NullInt64 `gorm:"column:maker_id;type:integer;" json:"maker_id"`
	//[ 6] process                                        text                 null: true   primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Process sql.NullString `gorm:"column:process;type:text;" json:"process"`
	//[ 7] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime sql.NullInt64 `gorm:"column:create_time;type:integer;" json:"create_time"`
	//[ 8] deadline_time                                  integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	DeadlineTime sql.NullInt64 `gorm:"column:deadline_time;type:integer;" json:"deadline_time"`
	//[ 9] order_status                                   integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	OrderStatus sql.NullInt64 `gorm:"column:order_status;type:integer;" json:"order_status"`
	//[10] admin_status                                   integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	AdminStatus sql.NullInt64 `gorm:"column:admin_status;type:integer;" json:"admin_status"`
}

var ordersTableInfo = &TableInfo{
	Name: "orders",
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
			Name:               "customer_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "CustomerName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "customer_name",
			ProtobufFieldName:  "customer_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "file_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "FileName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "file_name",
			ProtobufFieldName:  "file_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "department",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "Department",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "department",
			ProtobufFieldName:  "department",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "material_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "MaterialID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "material_id",
			ProtobufFieldName:  "material_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "maker_id",
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
			GoFieldName:        "MakerID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "maker_id",
			ProtobufFieldName:  "maker_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "process",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       -1,
			GoFieldName:        "Process",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "process",
			ProtobufFieldName:  "process",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "deadline_time",
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
			GoFieldName:        "DeadlineTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "deadline_time",
			ProtobufFieldName:  "deadline_time",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "order_status",
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
			GoFieldName:        "OrderStatus",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "order_status",
			ProtobufFieldName:  "order_status",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "admin_status",
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
			GoFieldName:        "AdminStatus",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "admin_status",
			ProtobufFieldName:  "admin_status",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Orders) TableName() string {
	return "orders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Orders) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Orders) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Orders) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Orders) TableInfo() *TableInfo {
	return ordersTableInfo
}
