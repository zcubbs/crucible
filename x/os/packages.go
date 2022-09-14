package os

import (
	"fmt"
	"os/exec"
)

func Install(packages ...string) error {
	for _, p := range packages {
		stdout, err := exec.Command("apt", "install", "-y", p).Output()
		if err != nil {
			return err
		}
		fmt.Println(string(stdout))
	}
	return nil
}

func Update() error {
	stdout, err := exec.Command("apt", "update").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))
	return nil
}

func Upgrade() error {
	stdout, err := exec.Command("apt", "upgrade").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))
	return nil
}
