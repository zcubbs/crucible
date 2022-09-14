package os

import (
	xos "crucible/x/os"
	"github.com/spf13/cobra"
)

// update represents the list command
var update = &cobra.Command{
	Use:   "update",
	Short: "OS Update",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := xos.Update()
		if err != nil {
			println(err.Error())
			panic("Could not Update OS")
		}
	},
}

func init() {
	Cmd.AddCommand(update)
}
