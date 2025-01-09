package void

import (
	"fmt"
	"sort"
)

func (cli *VoidCLI) printHelp() {
	fmt.Printf("%s version %s\n\n", cli.appName, cli.appVersion)
	fmt.Println("Usage:")
	fmt.Printf("  %s <command> [arguments]\n\n", cli.appName)
	fmt.Println("Available commands:")

	// Get sorted command names for consistent output
	var names []string
	for name := range cli.commands {
		names = append(names, name)
	}
	sort.Strings(names)

	// Find the longest command name for padding
	maxLen := 0
	for _, name := range names {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	// Print commands with aligned descriptions
	for _, name := range names {
		cmd := cli.commands[name]
		fmt.Printf("  %-*s  %s\n", maxLen, name, cmd.Description)
	}

	fmt.Println("\nCommon commands:")
	fmt.Printf("  %-*s  %s\n", maxLen, "help", "Show this help message")
	fmt.Printf("  %-*s  %s\n", maxLen, "version", "Show version information")
}