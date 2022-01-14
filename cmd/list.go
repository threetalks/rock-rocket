/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "批量查询企业信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("legal", "", "", "法人姓名")
	listCmd.Flags().StringP("contact", "", "", "联系方式：手机号或者座机")
	listCmd.Flags().StringP("province", "", "", "省份")
	listCmd.Flags().StringP("city", "", "", "城市")
	listCmd.Flags().StringP("county", "", "", "区县")
	listCmd.Flags().StringP("capital", "", "", "资本范围,格式为：min|max; 单位万元, -1 表示没有注册资本")
	listCmd.Flags().StringP("registered", "", "", "注册时间,格式为：start|end; 时间格式为 2000-12-31|2021-12-21 ")
	listCmd.Flags().StringP("industry", "", "", "行业")
	listCmd.Flags().StringP("type", "", "", "公司类型")
	listCmd.Flags().StringP("address", "", "", "地址：采用模糊搜索")
	listCmd.Flags().StringP("scope", "", "", "经营范围:采用模糊搜索")
	listCmd.Flags().StringP("output", "", "", "数据文件路径,根据数据文件后缀确定数据文件格式"+
		"支持json、csv、vcf")
}
