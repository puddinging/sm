/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"sm/utils"

	"github.com/spf13/cobra"
)

var (
	alias string
)

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().StringVarP(&alias, "alias", "a", "", "alias")
}

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open an ssh connection session via alias in the current console",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := utils.FindConfigByAlias(configPath, alias)
		if err == nil {
			server.Connect()
			return nil
		} else {
			return errors.New("no matching server found")
		}
	},
}
