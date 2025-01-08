package main

import (
    "fmt"
    "os"
    "github.com/osamikoyo/void"
)

// Task represents a simple task
type Task struct {
    Title  string
    Status string
}

var tasks []Task

func main() {
    // Initialize the CLI application
    cli := void.NewCLI("taskman", "1.0.0")

    // Register the "add" command
    err := cli.RegisterCommand("add", "Add a new task", handleAddTask)
    if err != nil {
        fmt.Printf("Error registering add command: %v\n", err)
        os.Exit(1)
    }

    // Register the "list" command
    err = cli.RegisterCommand("list", "List all tasks", handleListTasks)
    if err != nil {
        fmt.Printf("Error registering list command: %v\n", err)
        os.Exit(1)
    }

    // Register the "complete" command
    err = cli.RegisterCommand("complete", "Mark a task as complete", handleCompleteTask)
    if err != nil {
        fmt.Printf("Error registering complete command: %v\n", err)
        os.Exit(1)
    }

    // Run the CLI application
    if err := cli.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

func handleAddTask(args *void.ArgRouter) error {
    if len(args.Args()) < 1 {
        return fmt.Errorf("task title is required")
    }

    title := args.Args()[0]
    tasks = append(tasks, Task{
        Title:  title,
        Status: "pending",
    })

    fmt.Printf("Added task: %s\n", title)
    return nil
}

func handleListTasks(args *void.ArgRouter) error {
    if len(tasks) == 0 {
        fmt.Println("No tasks found")
        return nil
    }

    fmt.Println("Tasks:")
    for i, task := range tasks {
        fmt.Printf("%d. [%s] %s\n", i+1, task.Status, task.Title)
    }
    return nil
}

func handleCompleteTask(args *void.ArgRouter) error {
    if len(args.Args()) < 1 {
        return fmt.Errorf("task number is required")
    }

    taskNum := 0
    _, err := fmt.Sscanf(args.Args()[0], "%d", &taskNum)
    if err != nil {
        return fmt.Errorf("invalid task number: %s", args.Args()[0])
    }

    if taskNum < 1 || taskNum > len(tasks) {
        return fmt.Errorf("task number out of range")
    }

    tasks[taskNum-1].Status = "completed"
    fmt.Printf("Marked task %d as completed\n", taskNum)
    return nil
}