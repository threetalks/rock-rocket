/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "直接通过 sql 语句查询",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sql called args:%v,length:%d\n", args, len(args))
	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)
	sqlCmd.Flags().StringP("output", "", "", "数据文件路径,根据数据文件后缀确定数据文件格式"+
		"支持json、csv、vcf")
}
