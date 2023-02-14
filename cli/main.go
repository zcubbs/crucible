package main

import (
	"github.com/zcubbs/crucible/cli/cmd"
	"github.com/zcubbs/crucible/cli/configs"
)

func main() {
	configs.Bootstrap()
	cmd.Execute()
}
