/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"sm/model"
	"sm/utils"

	"github.com/spf13/cobra"
)

var (
	ip       string
	username string
	password string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&alias, "alias", "a", "", "alias")
	addCmd.Flags().StringVarP(&ip, "ip", "i", "", "id adress")
	addCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "password")
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("add commond ip: %s, username: %s, password: %s.\n", ip, username, password)
		config, err := utils.GetConfig(configPath)
		if err != nil {
			return errors.New("配置文件未找到")
		}
		server := model.Server{
			Alias:    alias,
			IP:       ip,
			Username: username,
			Password: password,
		}
		serverList := config.ServerList
		serverList = append(serverList, server)
		return nil
	},
}
