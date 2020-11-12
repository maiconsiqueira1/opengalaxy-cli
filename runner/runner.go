package runner

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

// New create a new command to run
func New(cmd *cobra.Command, args []string) Command {
	return Command{
		Operation: cmd.Name(),
		Entity:    args[0],
		ExtraArgs: args[1:],
	}
}

// Run runs the called command
func (cmd Command) Run() {
	fmt.Printf("%+v", cmd)
}
