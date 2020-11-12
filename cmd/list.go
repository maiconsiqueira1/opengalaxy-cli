package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/opengalaxy-cli/cmd/helpers"
	"github.com/thiagohdeplima/opengalaxy-cli/runner"
)

// ListCmd command to list resources
var ListCmd = &cobra.Command{
	Use:       "list [ARGUMENT]",
	Short:     "list [ARGUMENT]",
	Args:      helpers.ValidateArgs,
	ValidArgs: []string{"galaxies"},
	Run:       runner.ExecuteCommand,

	DisableFlagsInUseLine: true,
}
