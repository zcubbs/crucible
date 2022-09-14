package helm

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "helm",
	Short: "Helm Helper Commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("helm called")
	},
}
