package k3s

import (
	"crucible/x/k3s"
	"github.com/spf13/cobra"
)

// install represents the list command
var install = &cobra.Command{
	Use:   "install",
	Short: "k3s install",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := k3s.Install()
		if err != nil {
			println(err.Error())
			panic("Could not install k3s")
		}
	},
}

func init() {
	Cmd.AddCommand(install)
}
