package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/opengalaxy-cli/cmd/helpers"
	"github.com/thiagohdeplima/opengalaxy-cli/runner"
)

// DescribeCmd command to list resources
var DescribeCmd = &cobra.Command{
	Use:       "describe [ARGUMENT]",
	Short:     "describe a resource",
	Args:      helpers.ValidateArgs,
	ValidArgs: []string{"galaxies"},
	Run:       runner.ExecuteCommand,

	DisableFlagsInUseLine: true,
}
