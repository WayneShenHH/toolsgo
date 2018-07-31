// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// database
	{
		viper.SetDefault("database.username", "root")
		viper.SetDefault("database.password", "123456")
		viper.SetDefault("database.host", "localhost")
		viper.SetDefault("database.port", "3306")
		viper.SetDefault("database.encoding", "utf8")
		viper.SetDefault("database.name", "dev_sbodds")
		viper.SetDefault("database.timeout", "5000")
		viper.SetDefault("database.maxidleconns", "10")
		viper.SetDefault("database.maxconns", "10")
		viper.SetDefault("database.logmode", "false")
	}
	// redis
	{
		viper.SetDefault("redis.host", "localhost")
		viper.SetDefault("redis.port", "6379")
		viper.SetDefault("redis.index", "0")
		viper.SetDefault("redis.maxidleconns", "10")
		viper.SetDefault("redis.maxconns", "10")
	}
	// notify
	{
		viper.SetDefault("notify.slackenable", "false")
	}
	// HTTP
	{
		viper.SetDefault("http.addr", ":18086")
		viper.SetDefault("http.pingaddr", "http://127.0.0.1:18086")
		viper.SetDefault("http.b2baddr", ":18087")
		viper.SetDefault("http.pingb2baddr", "http://127.0.0.1:18087")
		viper.SetDefault("http.wsaddr", ":18088")
		viper.SetDefault("http.pingwsaddr", "http://127.0.0.1:18088")
	}
	// swagger
	{
		viper.SetDefault("swagger.enable", "false")
		viper.SetDefault("swagger.filepath", "./swagger.json")
	}
	// websocket
	{
		viper.SetDefault("websocket.pingdelay", "10")
	}
	// logger
	{
		viper.SetDefault("logger.stdlevel", "debug")
		viper.SetDefault("logger.filelevel", "debug")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// 吃設定檔 & 環境變數
	viper.AutomaticEnv()        // read in environment variables that match
	viper.SetEnvPrefix("libgo") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigType("yaml")
	if cfgFile != "" {
		// Use config file from the flag.
		readConfigFile()
	}

	app.Config()
	// 初始化 logger
	logger.Init()
}

func readConfigFile() {
	var b []byte
	var err error
	viper.SetConfigFile(cfgFile)
	b, err = ioutil.ReadFile(cfgFile)
	if err == nil {
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err == nil {
			fmt.Println(fmt.Sprint("Using config file: ", viper.ConfigFileUsed()))

		} else {
			fmt.Println(fmt.Sprintf("Read config failed: %v", err))
		}
	}
}
