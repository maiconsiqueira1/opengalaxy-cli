package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/opengalaxy-cli/cmd/helpers"
)

// describeExec print a list of requested resourcces
func describeExec(cmd *cobra.Command, args []string) {
	fmt.Printf("TODO: Implement command %s\n", cmd.Name())
}

// DescribeCmd command to list resources
var DescribeCmd = &cobra.Command{
	Use:                   "describe [ARGUMENT]",
	Short:                 "describe a resource",
	Args:                  helpers.ValidateArgs,
	ValidArgs:             []string{"galaxies"},
	Run:                   describeExec,
	DisableFlagsInUseLine: true,
}
