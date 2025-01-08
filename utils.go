package void

import (
	"fmt"
	"os"
)

// HandlerCommand represents a command handler function
type HandlerCommand func([]string) error

// VoidCommand represents a CLI command with its metadata
type VoidCommand struct {
	Name        string
	Description string
	Handler     HandlerCommand
}

// VoidCLI is the main CLI application
type VoidCLI struct {
	commands map[string]VoidCommand
}

// NewCLI creates a new CLI application
func NewCLI() *VoidCLI {
	return &VoidCLI{
		commands: make(map[string]VoidCommand),
	}
}

// RegisterCommand adds a new command to the CLI
func (cli *VoidCLI) RegisterCommand(name string, description string, handler HandlerCommand) {
	cli.commands[name] = VoidCommand{
		Name:        name,
		Description: description,
		Handler:     handler,
	}
}

// Run executes the CLI application
func (cli *VoidCLI) Run() error {
	if len(os.Args) < 2 {
		cli.printHelp()
		return nil
	}

	cmdName := os.Args[1]
	if cmdName == "help" {
		cli.printHelp()
		return nil
	}

	cmd, exists := cli.commands[cmdName]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmdName)
	}

	return cmd.Handler(os.Args[2:])
}

// printHelp displays the help message with available commands
func (cli *VoidCLI) printHelp() {
	fmt.Println("Available commands:")
	for name, cmd := range cli.commands {
		fmt.Printf("  %s: %s\n", name, cmd.Description)
	}
}