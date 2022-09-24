package awx

import (
	"crucible/cli/cmd/helm"
	"crucible/cli/configs"
	"github.com/spf13/cobra"
)

var (
	kubeconfig   string
	awxVersion   string
	namespace    string
	chartVersion = "0.29.0"
	repoUrl      = "https://ansible.github.io/awx-operator/"
	repoName     = "awx-operator"
	chartName    = "awx-operator"
	chartValues  = map[string]interface{}{}
)

// install represents the list command
var install = &cobra.Command{
	Use:   "install",
	Short: "Install AWX Instance",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if namespace == "" {
			namespace = "awx"
		}
		if awxVersion == "" {
			awxVersion = "19.1.0"
		}

		kubeconfig = configs.Config.Kubeconfig.Path

		helm.ExecuteInstallChartCmd(
			kubeconfig,
			chartName,
			repoName,
			repoUrl,
			namespace,
			chartVersion,
			chartValues,
		)
	},
}

func init() {
	install.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace value")
	install.Flags().StringVarP(&awxVersion, "version", "v", "", "version value")

	Cmd.AddCommand(install)
}
