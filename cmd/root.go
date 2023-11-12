/*
Copyright © 2023 NAME HERE <jiefeng.wang@outlook.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sm/model"
	"sm/utils"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sm",
	Short: "Manage all your server connection information",
	Long:  `Manage all your server connection information`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 获取配置路径
	userHomeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println("无法获取用户主目录：", err)
		return
	}
	configPath = filepath.Join(userHomeDir, ".sm", "config.json")

	// 获取文件所在目录
	dir := filepath.Dir(configPath)

	// 确保目录存在
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("无法创建目录：", err)
		return
	}

	// 检查文件是否存在
	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在，正在创建：", configPath)
			config := model.Config{
				ServerList: []model.Server{},
			}
			utils.WriteConfig(config, configPath)
		} else {
			fmt.Println("无法访问文件：", err)
		}
		return
	}
}
