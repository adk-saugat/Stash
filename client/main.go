package main

import (
	"fmt"
	"os"

	"github.com/adk-saugat/stash/client/commands"
	"github.com/adk-saugat/stash/client/core"
)

func main() {
	registry := core.NewRegistry()

	registry.Register(&commands.LoginCommand{})
	registry.Register(&commands.CreateCommand{})
	registry.Register(&commands.WatchCommand{})
	registry.Register(&commands.StoreCommand{})
	registry.Register(commands.NewHelpCommand(registry))
	registry.Register(&commands.ConfigCommand{})
	registry.Register(&commands.LogCommand{})

	if len(os.Args) < 2 {
		fmt.Println("Usage: stash <command> [args]")
		fmt.Println("Run 'stash help' for available commands.")
		os.Exit(1)
	}

	cmd, found := registry.Get(os.Args[1])
	if !found {
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Run 'stash help' for available commands.")
		os.Exit(1)
	}

	if err := cmd.Run(os.Args[2:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
