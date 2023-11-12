/*
Copyright © 2023 NAME HERE <jiefeng.wang@outlook.com>
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
	Short: "add a server infomation",
	Long:  ``,
	// DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if alias == "" {
			return errors.New("the flag alias cannot be empty")
		}

		if ip == "" {
			return errors.New("the flag ip cannot be empty")
		}

		if username == "" {
			return errors.New("the username ip cannot be empty")
		}

		if password == "" {
			return errors.New("the flag password cannot be empty")
		}

		fmt.Printf("add commond ip: %s, username: %s, password: %s.\n", ip, username, password)
		config, err := utils.GetConfig(configPath)
		if err != nil {
			return errors.New("configuration file not found")
		}
		server := model.Server{
			Alias:    alias,
			IP:       ip,
			Username: username,
			Password: password,
		}
		serverList := config.ServerList
		for _, server := range serverList {
			if server.Alias == alias {
				errMsg := fmt.Sprintf("the alias [%s] already exists", alias)
				return errors.New(errMsg)
			}
		}
		serverList = append(serverList, server)
		config.ServerList = serverList
		utils.WriteConfig(config, configPath)
		return nil
	},
}
