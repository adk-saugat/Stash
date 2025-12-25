package commands

import (
	"fmt"

	"github.com/adk-saugat/stash/core"
)

type HelpCommand struct {
	Registry *core.Registry
}

func NewHelpCommand(registry *core.Registry) *HelpCommand {
	return &HelpCommand{Registry: registry}
}

func (c *HelpCommand) Name() string        { return "help" }
func (c *HelpCommand) Description() string { return "Show available commands" }

func (c *HelpCommand) Run(args []string) error {
	fmt.Println("Stash - A simple version control system")
	fmt.Println()
	fmt.Println("Usage: stash <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")

	for name, cmd := range c.Registry.All() {
		fmt.Printf("  %-12s %s\n", name, cmd.Description())
	}

	return nil
}

