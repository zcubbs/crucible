package k8s

import (
	"crucible/x/kubectl"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var namespace string

// createNamespace represents the list command
var createNamespace = &cobra.Command{
	Use:   "create-ns",
	Short: "Show k8s create-ns list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteCreateNamespaceCmd(namespace)
	},
}

func ExecuteCreateNamespaceCmd(namespace string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	homedir = fmt.Sprintf("%s/.kube/config", homedir)
	err = kubectl.CreateNamespace(
		homedir,
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
