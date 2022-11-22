package helm

import (
	"crucible/core/helm"
	"github.com/spf13/cobra"
	"log"
)

// installHelm represents the list command
var installHelm = &cobra.Command{
	Use:   "install-helm",
	Short: "install helm CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := ExecuteInstallHelmCmd()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	Cmd.AddCommand(installHelm)
}

func ExecuteInstallHelmCmd() error {
	// Add helm repo
	err := helm.InstallHelmCLI()
	if err != nil {
		return err
	}
	return nil
}
