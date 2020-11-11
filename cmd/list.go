package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/opengalaxy-cli/cmd/helpers"
)

// RunList print a list of requested resources
func RunList(cmd *cobra.Command, args []string) {
	fmt.Printf("TODO: Implement command %s\n", cmd.Name())
}

// ListCmd command to list resources
var ListCmd = &cobra.Command{
	Use:                   "list [ARGUMENT]",
	Short:                 "list resources",
	Args:                  helpers.ValidateArgs,
	ValidArgs:             []string{"galaxies"},
	Run:                   RunList,
	DisableFlagsInUseLine: true,
}
