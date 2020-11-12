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

// ExecuteCommand blabal
func ExecuteCommand(cmd *cobra.Command, args []string) {
	command := New(cmd, args)
	command.Run()
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
	switch cmd.Operation {
	case "list":
		listEntities(cmd.Entity)
		break

	case "describe":
		describeEntity(cmd.Entity, cmd.ExtraArgs)
		break

	default:
		fmt.Printf("Error: invalid operation %s\n", cmd.Operation)
	}
}

func listEntities(entity string) {
	fmt.Printf("listing all %s\n", entity)
}

func describeEntity(entity string, extraArgs []string) {
	fmt.Printf("describing entity %s with args %s\n", entity, extraArgs)
}
