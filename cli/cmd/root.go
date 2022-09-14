package cmd

import (
	"crucible/cli/cmd/config"
	"crucible/cli/cmd/helm"
	"crucible/cli/cmd/info"
	"crucible/cli/cmd/k3s"
	"crucible/cli/cmd/os"
	"crucible/cli/configs"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	xos "os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "crucible",
		Short: "This is a CLI for the Crucible Bot project",
		Long:  "",
	}

	viperCommand = &cobra.Command{
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(viper.GetString("Flag"))
		},
	}

	aboutCmd = &cobra.Command{
		Use:   "about",
		Short: "Print the info about crucible-cli",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(configs.Splash + "\n\n")
			fmt.Println("steering-cli")
			fmt.Printf("Version: %s\n", configs.Version)
			fmt.Println("This is a CLI for the Crucible Bot project")
			fmt.Println("Copyright (c) 2022 zcubbs")
			fmt.Println("Repository: https://github.com/zcubbs/crucible/cli")
		},
	}

	persistRootFlag bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persist", "p", false, "Persist the CLI")
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(viperCommand)
	addSubCommandPalettes()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		xos.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(config.Cmd)
	rootCmd.AddCommand(os.Cmd)
	rootCmd.AddCommand(k3s.Cmd)
	rootCmd.AddCommand(info.Cmd)
	rootCmd.AddCommand(helm.Cmd)
}
