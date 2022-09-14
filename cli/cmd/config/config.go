package config

import (
	"crucible/cli/configs"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "config",
	Short: "list cli configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		jsonConfig, err := json.MarshalIndent(&configs.Config, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", string(jsonConfig))
	},
}

func init() {
}
