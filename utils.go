package void

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ArgRouter handles command-line argument parsing and routing
type ArgRouter struct {
	args      []string
	flags     map[string]string
	boolFlags map[string]bool
}

// NewArgRouter creates a new argument router
func NewArgRouter(args []string) *ArgRouter {
	router := &ArgRouter{
		args:      make([]string, 0),
		flags:     make(map[string]string),
		boolFlags: make(map[string]bool),
	}
	
	// Parse arguments and flags
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			// Handle long flags
			name := strings.TrimPrefix(arg, "--")
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				router.flags[name] = args[i+1]
				i++
			} else {
				router.boolFlags[name] = true
			}
		} else if strings.HasPrefix(arg, "-") {
			// Handle short flags
			name := strings.TrimPrefix(arg, "-")
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				router.flags[name] = args[i+1]
				i++
			} else {
				router.boolFlags[name] = true
			}
		} else {
			// Regular argument
			router.args = append(router.args, arg)
		}
	}
	
	return router
}

// Args returns the non-flag arguments
func (r *ArgRouter) Args() []string {
	return r.args
}

// Flag returns the value of a flag and whether it exists
func (r *ArgRouter) Flag(name string) (string, bool) {
	value, exists := r.flags[name]
	return value, exists
}

// HasFlag checks if a boolean flag is set
func (r *ArgRouter) HasFlag(name string) bool {
	return r.boolFlags[name]
}

// GetFlag returns the value of a flag or a default value if not found
func (r *ArgRouter) GetFlag(name string, defaultValue string) string {
	if value, exists := r.flags[name]; exists {
		return value
	}
	return defaultValue
}

// HandlerCommand represents a command handler function
type HandlerCommand func(args *ArgRouter) error

// VoidCommand represents a CLI command with its metadata
type VoidCommand struct {
	Name        string
	Description string
	Handler     HandlerCommand
}

// VoidCLI is the main CLI application
type VoidCLI struct {
	commands    map[string]VoidCommand
	appName     string
	appVersion  string
}

// NewCLI creates a new CLI application
func NewCLI(appName string, version string) *VoidCLI {
	return &VoidCLI{
		commands:    make(map[string]VoidCommand),
		appName:     appName,
		appVersion:  version,
	}
}

// RegisterCommand adds a new command to the CLI
func (cli *VoidCLI) RegisterCommand(name string, description string, handler HandlerCommand) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("command name cannot be empty")
	}

	if handler == nil {
		return fmt.Errorf("command handler cannot be nil")
	}

	if _, exists := cli.commands[name]; exists {
		return fmt.Errorf("command '%s' already registered", name)
	}

	cli.commands[name] = VoidCommand{
		Name:        name,
		Description: description,
		Handler:     handler,
	}
	return nil
}

// Run executes the CLI application
func (cli *VoidCLI) Run() error {
	if len(os.Args) < 2 {
		cli.printHelp()
		return nil
	}

	cmdName := os.Args[1]
	if cmdName == "help" || cmdName == "-h" || cmdName == "--help" {
		cli.printHelp()
		return nil
	}

	if cmdName == "version" || cmdName == "-v" || cmdName == "--version" {
		fmt.Printf("%s version %s\n", cli.appName, cli.appVersion)
		return nil
	}

	cmd, exists := cli.commands[cmdName]
	if !exists {
		cli.printHelp()
		return fmt.Errorf("unknown command: %s", cmdName)
	}

	if err := cmd.Handler(NewArgRouter(os.Args[2:])); err != nil {
		return fmt.Errorf("error executing '%s': %w", cmdName, err)
	}
	return nil
}

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