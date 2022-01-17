/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version = "No Version Provided"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "展示 Rock-Rocket 版本",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("The Rock-Rocket version is: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
