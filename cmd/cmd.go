package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command represents an requested
type Command struct {
	Operation string
	Entity    string
	ExtraArgs []string
}

// ExecuteCommand runs the called command
func ExecuteCommand(cmd *cobra.Command, args []string) {
	command := Command{
		Operation: cmd.Name(),
		Entity:    args[0],
		ExtraArgs: args[1:],
	}

	command.run()
}

func (command *Command) run() {
	fmt.Printf("%+v", command)
}
