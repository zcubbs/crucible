package os

import (
	xos "crucible/core/os"
	"github.com/spf13/cobra"
)

// upgrade represents the list command
var upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "OS Upgrade",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := xos.Upgrade()
		if err != nil {
			println(err.Error())
			panic("Could not Upgrade OS")
		}
	},
}

func init() {
	Cmd.AddCommand(upgrade)
}
