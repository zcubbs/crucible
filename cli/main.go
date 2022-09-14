package main

import (
	"crucible/cli/cmd"
	"crucible/cli/configs"
)

func main() {
	configs.Bootstrap()
	cmd.Execute()
}
