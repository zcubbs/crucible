package helm

import (
	"crucible/cli/cmd/k8s"
	"crucible/x/helm"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	repoName     string
	repoUrl      string
	chartName    string
	namespace    string
	chartVersion string
	chartValues  map[string]interface{}
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

		fmt.Println("repoName: ", repoName)
		fmt.Println("repoUrl: ", repoUrl)
		fmt.Println("chartName: ", chartName)
		fmt.Println("namespace: ", namespace)
		fmt.Println("chartVersion: ", chartVersion)
		fmt.Println("chartValues: ", chartValues)

		// Execute Command
		ExecuteInstallChartCmd(chartName, repoName, repoUrl, namespace, chartVersion, chartValues)

	},
}

func init() {
	installChart.Flags().StringVarP(&repoName, "repo-name", "r", "", "Helm repo name")
	installChart.Flags().StringVarP(&repoUrl, "repo-url", "u", "", "Helm repo url")
	installChart.Flags().StringVarP(&chartName, "chart-name", "c", "", "Helm chart name")
	installChart.Flags().StringVarP(&namespace, "namespace", "n", "", "Helm chart namespace")
	installChart.Flags().StringVarP(&chartVersion, "chart-version", "v", "", "Helm chart version")

	if err := installChart.MarkFlagRequired("repo-name"); err != nil {
		log.Println(err)
	}
	if err := installChart.MarkFlagRequired("repo-url"); err != nil {
		log.Println(err)
	}
	if err := installChart.MarkFlagRequired("chart-name"); err != nil {
		log.Println(err)
	}
	if err := installChart.MarkFlagRequired("namespace"); err != nil {
		log.Println(err)
	}

	Cmd.AddCommand(installChart)
}

func ExecuteInstallChartCmd(chartName, repoName, repoUrl, namespace, chartVersion string, chartValues map[string]interface{}) {
	// Add helm repo
	helm.RepoAdd(repoName, repoUrl)
	// Update charts from the helm repo
	helm.RepoUpdate()
	// Create Namespace
	k8s.ExecuteCreateNamespaceCmd(namespace)
	// Install charts
	helm.InstallChart(chartName, repoName, namespace, chartVersion, chartName, chartValues)
	// List helm releases
	ExecuteHelmListCmd()
}
