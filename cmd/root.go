/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
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

			// 创建文件
			file, err := os.Create(configPath)
			if err != nil {
				fmt.Println("无法创建文件：", err)
				return
			}
			defer file.Close()

			fmt.Println("文件创建成功！")
		} else {
			fmt.Println("无法访问文件：", err)
		}
		return
	}
}
