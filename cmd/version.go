package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionExec print the cli version
func versionExec(cmd *cobra.Command, args []string) {
	fmt.Println("Open Galaxy Command Line Interface -- DEVELOPMENT")
}

// VersionCmd command to show the version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the cli version",
	Run:   versionExec,
}
