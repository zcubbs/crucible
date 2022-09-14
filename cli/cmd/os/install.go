package os

import (
	xos "crucible/x/os"
	"github.com/spf13/cobra"
)

// update represents the list command
var install = &cobra.Command{
	Use:   "install",
	Short: "OS install packages",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := xos.Install(args...)
		if err != nil {
			println(err.Error())
			panic("Could not install packages")
		}
	},
}

func init() {
	Cmd.AddCommand(install)
}
