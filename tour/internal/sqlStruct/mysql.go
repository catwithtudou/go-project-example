package sqlStruct

import (
	"database/sql"
	"errors"
	"fmt"
)

/**
 *@Author tudou
 *@Date 2020/7/22
 **/

//数据库连接核心对象
type DBModel struct{
	DBEngine *sql.DB
	DBInfo *DBInfo
	DBTypeTrans map[string]string
}

//存储连接MySQL的基本信息
type DBInfo struct{
	DBType string
	Host string
	Username string
	Password string
	Charset string
}

//存储COLUMNS表中需要的字段
type TableColumn struct{
	ColumnName string
	DataType string
	IsNullable string
	ColumnKey string
	ColumnType string
	ColumnComment string
}


func NewDBModel(info *DBInfo) *DBModel{
	return &DBModel{DBInfo: info,DBTypeTrans:GetMySQLDBTypeStruct()}
}

var dBTypeStruct map[string]map[string]string

func dBTypeTransGet(DBType string)map[string]string{
	return dBTypeStruct[DBType]
}

func GetMySQLDBTypeStruct()map[string]string{
	return  dBTypeTransGet("mysql")
}




//连接数据库
func (m *DBModel) Connect()(err error){
	connect:="%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	dsn:=fmt.Sprintf(connect,m.DBInfo.Username,m.DBInfo.Password,m.DBInfo.Host,m.DBInfo.Charset)
	m.DBEngine,err = sql.Open(m.DBInfo.DBType,dsn)
	return
}


//获取表中列的信息
func (m *DBModel) GetColumns(dbName,tableName string) (tableColumns []*TableColumn,err error){
	query:="SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows,err:=m.DBEngine.Query(query,dbName,tableName)
	if err!=nil{
		return nil,err
	}
	if rows==nil{
		return nil,errors.New("查询COLUMNS表失败")
	}

	for rows.Next(){
		column:=TableColumn{}
		err = rows.Scan(&column.ColumnName,&column.DataType,&column.ColumnKey,&column.IsNullable,&column.ColumnType,&column.ColumnComment)
		if err!=nil{
			return nil,err
		}
		tableColumns = append(tableColumns,&column)
	}
	return
}

//MySQL表字段类型映射
var mySQLTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func init(){
	dBTypeStruct["mysql"]=mySQLTypeToStructType
}