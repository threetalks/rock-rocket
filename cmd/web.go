/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "启动本地浏览器进行访问查询 默认端口为 8520",
	Long:  `启动本地浏览器进行访问查询 默认访问地址为 http://0.0.0.0:8520`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web in todo")
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().IntP("port", "p", 8520, "服务监听端口号")
	webCmd.Flags().StringP("server", "s", "0.0.0.0", "本地服务IP")
}
