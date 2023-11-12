/*
Copyright © 2023 NAME HERE <jiefeng.wang@outlook.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"sm/utils"

	"github.com/spf13/cobra"
)

var (
	alilsName string
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a server information by alias",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if alilsName == "" {
			return errors.New("alils cannot be empty")
		}
		config, err := utils.GetConfig(configPath)
		if err != nil {
			return errors.New("configuration file not found")
		}
		for i, server := range config.ServerList {
			if server.Alias == alilsName {
				fmt.Printf("The server will be deleted: %s", server.Alias)
				config.ServerList = append(config.ServerList[:i], config.ServerList[i+1:]...)
			}
		}
		utils.WriteConfig(config, configPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&alilsName, "alias", "a", "", "alias")
}
