package helm

import (
	"crucible/x/helm"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// installChart represents the list command
var installChart = &cobra.Command{
	Use:   "install-chart",
	Short: "list all helm releases",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		repoName := "bitnami"
		repoUrl := "https://charts.bitnami.com/bitnami"
		chartName := "nginx"
		namespace := "default"
		chartVersion := "13.2.4"
		chartValues := map[string]interface{}{
			"replicaCount": 3,
		}

		// Add helm repo
		helm.RepoAdd(repoName, repoUrl)
		// Update charts from the helm repo
		helm.RepoUpdate()
		// Install charts
		helm.InstallChart(chartName, repoName, namespace, chartVersion, chartName, chartValues)

		if err != nil {
			log.Fatal("Could not install helm release", err)
		}
	},
}

func init() {
	Cmd.AddCommand(installChart)
}
