package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/opengalaxy-cli/cmd/helpers"
)

// ListCmd command to list resources
var ListCmd = &cobra.Command{
	Use:       "list [ARGUMENT]",
	Short:     "list [ARGUMENT]",
	Args:      helpers.ValidateArgs,
	ValidArgs: []string{"galaxies"},
	Run:       ExecuteCommand,

	DisableFlagsInUseLine: true,
}
