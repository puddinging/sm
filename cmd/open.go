/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"sm/utils"
	"strings"

	"github.com/spf13/cobra"
)

var (
	alias string
)

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().StringVarP(&alias, "alias", "a", "", "alias")
	openCmd.Flags().SetAnnotation("alias", cobra.BashCompCustom, []string{"__sm_alias"})
}

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Connect to the specified server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := utils.FindConfigByAlias(configPath, alias)
		if err == nil {
			server.Connect()
			return nil
		} else {
			return errors.New("未找到匹配的服务器信息")
		}
	},
}

func __sm_alias(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 在这里实现自动补全的逻辑，可以根据你的服务器列表或配置文件提供建议
	suggestions := []string{"apaas-demo", "apaas-dev", "apaas-qa", "apaas-prod", "apaas-new-stable", "apaas-dev2"}
	filtered := make([]string, 0)

	// 过滤匹配的建议
	for _, s := range suggestions {
		if strings.HasPrefix(s, toComplete) {
			filtered = append(filtered, s)
		}
	}

	return filtered, cobra.ShellCompDirectiveNoFileComp
}
