package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/marcellmartini/aocclt/cmd/apply"
	"github.com/marcellmartini/aocclt/cmd/config"
	"github.com/marcellmartini/aocclt/cmd/get"
)

var rootCmd = &cobra.Command{
	Use:   "aocctl",
	Short: "CLI to help with the advent of code",
	Long: `aocctl is imagined to be a versatile tool that 
provides various functionalities to enhance the 
coding experience during the Advent of Code event`,
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func addSubcommands() {
	rootCmd.AddCommand(apply.CmdApply)
	rootCmd.AddCommand(config.CmdConfig)
	rootCmd.AddCommand(get.CmdGet)
}

func init() {
	addSubcommands()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
