/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the query command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "通过ID、企业名称、统一社会信用码查询，存在时返回最近注册的企业",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find one called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("name", "", "", "企业名称")
	getCmd.Flags().StringP("code", "", "", "统一社会信用码")
	getCmd.Flags().StringP("id", "", "", "通过ID查询")
	getCmd.Flags().StringP("output", "", "", "数据文件路径,根据数据文件后缀确定数据文件格式"+
		"支持json、csv、vcf")
}
