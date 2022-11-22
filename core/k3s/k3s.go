package k3s

import (
	xos "crucible/core/os"
	"os"
)

func Install() error {
	// curl -sfL https://get.k3s.io -o k3s-install.sh
	err := xos.ExecuteCmd(
		"curl",
		"https://get.k3s.io",
		"-o",
		"k3s-install.sh",
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

	// sh ./k3s-install.sh -s - --write-kubeconfig-mode 644
	err = os.Chmod("k3s-install.sh", 0700)
	if err != nil {
		return err
	}

	_, err = xos.ExecuteScript(
		"./k3s-install.sh",
		"./k3s-install.sh",
		"-s",
		"-",
		"--write-kubeconfig-mode=644",
	)
	if err != nil {
		return err
	}

	return nil
}
