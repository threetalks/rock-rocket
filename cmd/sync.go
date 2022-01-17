/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>

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
2. 远程同步需要添加服务器地址 比如: https://cos.rock_rocket.com/2022-01-01.csv
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync in todo")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringP("file", "f", "rock_rocket.csv", "同步文件路径")
}
