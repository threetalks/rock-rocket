/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "展示当前Rock-Rocket配置",
	Long:  `展示当前Rock-Rocket配置`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, key := range viper.AllKeys() {
			fmt.Printf("%s : %+v\n", key, viper.Get(key))
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
