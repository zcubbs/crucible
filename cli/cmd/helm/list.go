package helm

import (
	"crucible/x/helm"
	"crucible/x/utils"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

// list represents the list command
var list = &cobra.Command{
	Use:   "list",
	Short: "list all helm releases",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		_releases, err := helm.GetAllReleases(
			fmt.Sprintf("%s/%s", homedir, ".kube/config"),
		)

		if err != nil {
			log.Fatal("Could not list helm releases", err)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Namespace", "Release", "Name", "Version", "Status", "App Version", "Deployed"})
		for _, release := range _releases {
			t.AppendRows([]table.Row{
				{
					release.Namespace,
					release.Name,
					release.Chart.Metadata.Name,
					release.Chart.Metadata.Version,
					release.Info.Status,
					release.Chart.Metadata.AppVersion,
					utils.TimeElapsed(time.Now(), release.Info.FirstDeployed.Time, false),
				},
			})

		}
		t.AppendSeparator()
		t.SetStyle(table.StyleColoredDark)
		t.Style().Options.DrawBorder = false
		t.Style().Options.SeparateRows = false

		t.Render()
		if err != nil {
			println(err.Error())
			panic("Could not list helm releases")
		}

	},
}

func init() {
	Cmd.AddCommand(list)
}