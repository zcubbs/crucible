package os

import (
	"fmt"
	"os/exec"
)

func Install(packages ...string) error {
	for _, p := range packages {
		stdout, err := exec.Command("/bin/sh", "-c",
			fmt.Sprintf("sudo apt install -y %s", p)).Output()
		if err != nil {
			return err
		}
		fmt.Println(string(stdout))
	}
	return nil
}

func Update() error {
	stdout, err := exec.Command("/bin/sh", "-c", "sudo apt update -y").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))
	return nil
}

func Upgrade() error {
	stdout, err := exec.Command("/bin/sh", "-c", "sudo apt upgrade -y").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))
	return nil
}
