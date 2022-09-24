package helm

import (
	xos "crucible/x/os"
	"os"
)

func InstallHelmCLI() error {
	// curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
	err := xos.ExecuteCmd(
		"curl",
		"https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3",
		"-fsSL",
		"-o",
		"get_helm.sh",
	)
	if err != nil {
		return err
	}

	// ls -l
	err = xos.ExecuteCmd(
		"ls",
		"-l",
	)
	if err != nil {
		return err
	}

	// chmod 700 get_helm.sh
	err = os.Chmod("get_helm.sh", 0700)
	if err != nil {
		return err
	}

	// sh ./get_helm.sh
	_, err = xos.ExecuteScript(
		"./get_helm.sh",
		"./get_helm.sh",
	)
	if err != nil {
		return err
	}

	return nil
}
