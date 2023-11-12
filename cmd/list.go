/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"sm/utils"

	"github.com/spf13/cobra"
)

type ServerList struct {
	Alias    string `json:"alias"`
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Servers []ServerList `json:"serverList"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all stored server information",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("list filePath %s \n", configPath)
		config, err := utils.GetConfig(configPath)
		if err != nil {
			return errors.New("配置文件未找到")
		}
		// 计算 Alias 字段的最大宽度
		maxAliasWidth := 0
		for _, server := range config.ServerList {
			if len(server.Alias) > maxAliasWidth {
				maxAliasWidth = len(server.Alias)
			}
		}
		for _, server := range config.ServerList {
			fmt.Printf("Alias: %- *s\tIP: %-15s\n", maxAliasWidth, server.Alias, server.IP)
			fmt.Println("-----------------------------------------------")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}