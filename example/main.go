package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/osamikoyo/void"
)

func main() {
	cli := void.NewCLI()

	// Register a greeting command
	cli.RegisterCommand("greet", "Greet someone by name", func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("please provide a name")
		}
		name := strings.Join(args, " ")
		fmt.Printf("Hello, %s!\n", name)
		return nil
	})

	// Register an echo command
	cli.RegisterCommand("echo", "Echo the provided arguments", func(args []string) error {
		fmt.Println(strings.Join(args, " "))
		return nil
	})

	// Run the CLI application
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
