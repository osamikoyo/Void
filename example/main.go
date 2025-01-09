package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/osamikoyo/void"
)

var currentDir string

func init() {
    var err error
    currentDir, err = os.Getwd()
    if err != nil {
        fmt.Printf("Error getting current directory: %v\n", err)
        os.Exit(1)
    }
}

func main() {
    cli := void.NewCLI("fsexplorer", "1.0.0")

    // Register commands
    commands := map[string]struct {
        desc    string
        handler void.HandlerCommand
    }{
        "ls":   {"List directory contents", handleList},
        "cd":   {"Change current directory", handleChangeDir},
        "info": {"Show file information", handleFileInfo},
    }

    for cmd, details := range commands {
        if err := cli.RegisterCommand(cmd, details.desc, details.handler); err != nil {
            fmt.Printf("Error registering command %s: %v\n", cmd, err)
            os.Exit(1)
        }
    }

    if err := cli.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

func handleList(args *void.ArgRouter) error {
    // Get target directory (current dir if no args provided)
    targetDir := currentDir
    if len(args.Args()) > 0 {
        targetDir = filepath.Join(currentDir, args.Args()[0])
    }

    // Read directory contents
    entries, err := os.ReadDir(targetDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    // Display entries
    for _, entry := range entries {
        prefix := "f"
        if entry.IsDir() {
            prefix = "d"
        }
        fmt.Printf("[%s] %s\n", prefix, entry.Name())
    }

    return nil
}

func handleChangeDir(args *void.ArgRouter) error {
    if len(args.Args()) < 1 {
        return fmt.Errorf("directory path required")
    }

    // Handle absolute and relative paths
    newDir := args.Args()[0]
    if !filepath.IsAbs(newDir) {
        newDir = filepath.Join(currentDir, newDir)
    }

    // Verify directory exists and is accessible
    info, err := os.Stat(newDir)
    if err != nil {
        return fmt.Errorf("invalid directory: %w", err)
    }
    if !info.IsDir() {
        return fmt.Errorf("path is not a directory: %s", newDir)
    }

    currentDir = newDir
    return nil
}

func handleFileInfo(args *void.ArgRouter) error {
    if len(args.Args()) < 1 {
        return fmt.Errorf("file path required")
    }

    // Get file path
    path := args.Args()[0]
    if !filepath.IsAbs(path) {
        path = filepath.Join(currentDir, path)
    }

    // Get file info
    info, err := os.Stat(path)
    if err != nil {
        return fmt.Errorf("failed to get file info: %w", err)
    }

    // Display file information
    fmt.Printf("Name: %s\n", info.Name())
    fmt.Printf("Size: %d bytes\n", info.Size())
    fmt.Printf("Mode: %s\n", info.Mode())
    fmt.Printf("Modified: %s\n", info.ModTime())
    fmt.Printf("Is Directory: %t\n", info.IsDir())
	return nil
}