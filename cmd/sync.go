/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Rock-Rocket 企业数据同步",
	Long: `企业数据同步：支持 https、http 远程和本地json、csv 文件同步
1. 同步根据文件后缀判定文件类型
2. 远程同步需要添加服务器地址 比如: https://cos.rock-rocket.com/2022-01-01.csv
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync in todo")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	syncCmd.Flags().StringP("file", "f", "rock-rocket.csv", "同步文件路径")
}
