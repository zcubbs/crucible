package k3s

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "k3s",
	Short: "K3s Helper Commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("k3s called")
	},
}
