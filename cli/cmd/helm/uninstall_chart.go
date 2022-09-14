package helm

import (
	"crucible/x/helm"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// uninstallChart represents the list command
var uninstallChart = &cobra.Command{
	Use:   "uninstall-chart",
	Short: "list all helm releases",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		chartName := "nginx"
		namespace := "default"

		// Install charts
		helm.UninstallChart(chartName, namespace)

		if err != nil {
			log.Fatal("Could not uninstall helm release", err)
		}
	},
}

func init() {
	Cmd.AddCommand(uninstallChart)
}
