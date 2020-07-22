package cmd

import (
	"github.com/spf13/cobra"
	"go-project-example/tour/internal/sqlStruct"
	"log"
)

/**
 *@Author tudou
 *@Date 2020/7/22
 **/

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql语句转换和处理",
	Long: "sql语句转换和处理",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var sqlStructCmd = &cobra.Command{
	Use: "struct",
	Short: "sql语句转换",
	Long: "sql语句转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sqlStruct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sqlStruct.NewDBModel(dbInfo)
		err:= dbModel.Connect()
		if err!=nil{
			log.Fatalf("dbModel.Connect err: %s",err)
		}
		columns,err := dbModel.GetColumns(dbName,tableName)
		if err!=nil{
			log.Fatalf("dbModel.GetColumns err: %s",err)
		}

		template:=sqlStruct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName,templateColumns)
		if err!=nil{
			log.Fatalf("template.Gernerate err: %s",err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sqlStructCmd)
	sqlStructCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sqlStructCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sqlStructCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sqlStructCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sqlStructCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sqlStructCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sqlStructCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
