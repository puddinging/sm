/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"sm/model"
	"sm/utils"

	"io/fs"
	"os"

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
		config.ServerList = serverList
		WriteConfig(serverList, configPath)
		return nil
	},
}

func WriteConfig(config interface{}, filePath string) error {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New("无法序列化配置为 JSON 格式")
	}

	err = os.WriteFile(filePath, configJSON, fs.FileMode(0644))
	if err != nil {
		return errors.New("无法写入配置文件")
	}

	return nil
}
