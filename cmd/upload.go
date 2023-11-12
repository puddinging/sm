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

var (
	toUploadAlias  string
	localFilePath  string
	remoteFilePath string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := utils.FindConfigByAlias(configPath, toUploadAlias)
		if err != nil {
			return errors.New("no matching server found")
		}

		fmt.Printf("the alias [%s] ", toUploadAlias)
		server.UploadFile(localFilePath, remoteFilePath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&toUploadAlias, "alias", "a", "", "alias")
	uploadCmd.Flags().StringVarP(&localFilePath, "localFilePath", "l", "", "alias")
	uploadCmd.Flags().StringVarP(&remoteFilePath, "remoteFilePath", "r", "", "alias")
}
