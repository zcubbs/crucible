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

		// Install charts
		helm.UninstallChart(chartName, namespace)

		if err != nil {
			log.Fatal("Could not uninstall helm release", err)
		}
	},
}

func init() {
	uninstallChart.Flags().StringVarP(&chartName, "chart-name", "c", "", "Helm chart name")
	uninstallChart.Flags().StringVarP(&namespace, "namespace", "n", "", "Helm chart namespace")

	if err := uninstallChart.MarkFlagRequired("chart-name"); err != nil {
		log.Println(err)
	}
	if err := uninstallChart.MarkFlagRequired("namespace"); err != nil {
		log.Println(err)
	}

	Cmd.AddCommand(uninstallChart)
}
