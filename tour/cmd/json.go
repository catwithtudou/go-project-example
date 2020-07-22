package cmd

import (
	"github.com/spf13/cobra"
	"go-project-example/tour/internal/jsonStruct"
	"log"
)

/**
 *@Author tudou
 *@Date 2020/7/22
 **/


var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json转换和处理",
	Long:  "json转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var jsonStructCmd = &cobra.Command{
	Use:   "struct",
	Short: "json转换",
	Long:  "json转换",
	Run: func(cmd *cobra.Command, args []string) {
		parser, err := jsonStruct.NewParser(str)
		if err != nil {
			log.Fatalf("jsonStruct.NewParser err: %v", err)
		}
		content := parser.JsonStruct()
		log.Printf("输出结果: %s", content)
	},
}

func init() {
	jsonCmd.AddCommand(jsonStructCmd)
	jsonStructCmd.Flags().StringVarP(&str, "str", "s", "", "请输入json字符串")
}
