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
{    "system_id": 80,    "username": "jTlANtGOucEaivdiXlLruqESo",    "password": "IFEESPfqqxPVPSdZAjAiaEgBU",    "type": 42,    "create_time": 34,    "update_time": 63}



*/

// Users struct is a row record of the users table in the main database
type Users struct {
	//[ 0] system_id                                      integer              null: false  primary: true   isArray: false  auto: true   col: integer         len: -1      default: []
	SystemID int32 `gorm:"primary_key;AUTO_INCREMENT;column:system_id;type:integer;" json:"system_id"`
	//[ 1] username                                       text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Username string `gorm:"column:username;type:text;" json:"username"`
	//[ 2] password                                       text                 null: false  primary: false  isArray: false  auto: false  col: text            len: -1      default: []
	Password string `gorm:"column:password;type:text;" json:"password"`
	//[ 3] type                                           integer              null: false  primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	Type int32 `gorm:"column:type;type:integer;" json:"type"`
	//[ 4] create_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	CreateTime sql.NullInt64 `gorm:"column:create_time;type:integer;" json:"create_time"`
	//[ 5] update_time                                    integer              null: true   primary: false  isArray: false  auto: false  col: integer         len: -1      default: []
	UpdateTime sql.NullInt64 `gorm:"column:update_time;type:integer;" json:"update_time"`
}

var usersTableInfo = &TableInfo{
	Name: "users",
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
			Name:               "username",
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
			GoFieldName:        "Username",
			GoFieldType:        "string",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "password",
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
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "update_time",
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
			GoFieldName:        "UpdateTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Users) TableInfo() *TableInfo {
	return usersTableInfo
}
