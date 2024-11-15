package k8s

import (
	"crucible/cli/configs"
	"crucible/core/kubectl"
	"github.com/spf13/cobra"
	"log"
)

var namespace string

// createNamespace represents the list command
var createNamespace = &cobra.Command{
	Use:   "create-ns",
	Short: "Show k8s create-ns list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteCreateNamespaceCmd(
			configs.Config.Kubeconfig.Path,
			namespace,
		)
	},
}

func ExecuteCreateNamespaceCmd(kubeconfig, namespace string) {
	err := kubectl.CreateNamespace(
		kubeconfig,
		namespace,
	)
	if err != nil {
		println("error:", err.Error())
	}
}

func init() {
	createNamespace.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace value")
	if err := createNamespace.MarkFlagRequired("namespace"); err != nil {
		log.Println(err)
	}

	Cmd.AddCommand(createNamespace)
}
