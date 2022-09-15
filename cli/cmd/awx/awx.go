package awx

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "awx",
	Short: "Awx Instance Control Commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("awx called")
	},
}
