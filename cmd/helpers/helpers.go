package helpers

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// ValidateArgs validate
func ValidateArgs(cmd *cobra.Command, args []string) error {
	// when command doesn't expects arguments and user don't inform one
	if len(cmd.ValidArgs) == 0 && len(args) == 0 {
		return nil
	}

	// when user inform an argument but command doesn't expects it
	if len(cmd.ValidArgs) == 0 && len(args) > 0 {
		return showCorrectUsage(cmd)
	}

	// when user doesn't inform an argument but command expects it
	if len(cmd.ValidArgs) > 0 && len(args) == 0 {
		return showCorrectUsage(cmd)
	}

	for _, validArg := range cmd.ValidArgs {
		if validArg == args[0] {
			return nil
		}
	}

	return fmt.Errorf("Invalid option %s", args[0])
}

func showCorrectUsage(cmd *cobra.Command) error {
	if len(cmd.ValidArgs) == 0 {
		return fmt.Errorf("The command %s doesn't expect arguments", cmd.Name())
	}

	if len(cmd.ValidArgs) == 1 {
		return fmt.Errorf("You need to specify `%s` as argument", cmd.ValidArgs[0])
	}

	return fmt.Errorf("Invalid argument. The valid options are `[%s]`", strings.Join(cmd.ValidArgs, "\n"))
}
