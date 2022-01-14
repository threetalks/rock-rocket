/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webCmd.Flags().IntP("port", "p", 8520, "服务监听端口号")
	webCmd.Flags().StringP("server", "s", "0.0.0.0", "本地服务IP")
}
