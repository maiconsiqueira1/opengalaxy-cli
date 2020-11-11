package main

import (
	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/opengalaxy-cli/cmd"
)

var rootCmd = &cobra.Command{
	Use:                   "og-cli command [RESOURCE]",
	Short:                 "og-cli is a tool to manage open galaxy by the command line",
	DisableFlagsInUseLine: true,
}

func main() {
	rootCmd.AddCommand(
		cmd.VersionCmd,
		cmd.ListCmd,
		cmd.DescribeCmd,
	)

	rootCmd.Execute()
}
