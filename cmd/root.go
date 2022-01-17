/*
Copyright © 2022 Three Talks  <bytebody@icloud.com>
Rock-Rocket 方便实用的本地企业查询工具
*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gorm.io/gorm"
	"os"
	"rock-rocket/internal/db"
	"rock-rocket/internal/rock_rocket"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	rockRocket *rock_rocket.RockRocket
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rock_rocket",
	Short: "Rock-Rocket 是一个支持命令行和本地网页查询企业信息的工具库",
	Long: `Rock-Rocket 作为一个企业查询工具，支持如下功能：
1. 支持命令行根据企业名称、统一社会信用码、法人名字、注册地址等条件查询企业信息
2. 支持通过本地 WEB 查询
3. 支持企业信息通讯录导出`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	if version != "" {
		Version = version
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initRockRocket)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rock_rocket.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".rock_rocket" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rock_rocket")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

func initRockRocket() {
	driverName := viper.GetString("driverName")
	dsn := viper.GetString("dsn")

	if err := db.Initial(dsn, driverName, &gorm.Config{}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rockRocket = rock_rocket.NewRockRocket(db.DB())
}
