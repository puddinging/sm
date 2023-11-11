/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("list filePath %s \n", configPath)
		config := utils.GetConfig(configPath)
		for _, server := range config.ServerList {
			fmt.Printf("Alias: %-10s IP: %-15s\n", server.Alias, server.IP)
			fmt.Println("--------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
