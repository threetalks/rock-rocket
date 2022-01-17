/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
)

// codeCmd represents the query command
var idCmd = &cobra.Command{
	Use:   "id",
	Short: "通过 ID 查询企业",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("have to provide id")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("get company by id error: %s\n", err)
			return
		}
		company, err := rockRocket.GetCompanyByID(uint(id))
		if err != nil {
			fmt.Printf("get company by id:%s error:%s\n", args[0], err)
			return
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("get company by id:%s error:%s\n", args[0], err)
			return
		}
		format, _ := cmd.Flags().GetString("format")
		var data []byte
		switch format {
		case "vcf":
			data, err = json.Marshal(company)
		case "csv":
			data, err = json.Marshal(company)
		default:
			data, err = json.Marshal(company)
		}
		if output != "" {
			err = ioutil.WriteFile(output, data, os.ModePerm)
			if err != nil {
				fmt.Printf("save company:%s to %s error:%s\n", string(data), output, err)
				return
			}
			fmt.Printf("save company:%s to %s success\n", string(data), output)
		} else {
			fmt.Printf("get company by id:%s\n %s", args[0], string(data))
		}
	},
}

func init() {
	rootCmd.AddCommand(idCmd)
	idCmd.Flags().StringP("output", "o", "", "数据文件路径")
	idCmd.Flags().StringP("format", "", "json", "文件格式：json、csv、vcf")
}
