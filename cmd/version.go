package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	buildVsn   string
	buildHash  string
	buildState string
	buildTime  string
)

// versionExec print the cli version
func versionExec(cmd *cobra.Command, args []string) {
	fmt.Printf("open galaxy cli\n\n")

	if buildVsn != "" {
		fmt.Printf("build version:\t%s-%s\n", buildVsn, buildHash)
	} else {
		fmt.Printf("build version:\t%s\n", buildHash)
	}

	fmt.Printf("build state:\t%s\n", buildState)
	fmt.Printf("build time:\t%s\n", buildTime)
}

// VersionCmd command to show the version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the cli version",
	Run:   versionExec,

	DisableFlagsInUseLine: true,
}
