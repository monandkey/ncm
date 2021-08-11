package cmd

import (
	"github.com/spf13/cobra"
	"local.packages/ncm"
)

type nodeConfig struct {
	init bool
}

func init() {
	nodeCmd := &cobra.Command{}
	nodeCmd.Use = "node"
	nodeCmd.Short = "Control the resources of the worker node"

	nodeConfig := nodeConfig{
		init: false,
	}

	nodeCmd.Flags().BoolVar(&nodeConfig.init, "init", nodeConfig.init, "Create node config")

	nodeCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if nodeConfig.init {
			userInfo := ncm.UserInfo{}
			err := userInfo.GetUserInfomation()
			if err != nil {
				return err
			}
		}
		return nil
	}
	rootCmd.AddCommand(nodeCmd)
}
