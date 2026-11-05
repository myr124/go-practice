package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// captilization of field names is important for exporting
type Task struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TaskFile struct {
	Tasks []Task `json:"tasks"`
}

func main() {

	// ts := Task{}

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error finding dir:", err)
	}

	var tasks TaskFile
	filename := os.Args[1]

	filepath := filepath.Join(wd, filename)
	data, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &tasks)

	fmt.Printf("Tasks (%d total)\n", len(tasks.Tasks))

	for _, task := range tasks.Tasks {
		status := "✗"

		if task.Completed {
			status = "✓"
		}
		fmt.Printf("[%s] %d. %s\n", status, task.ID, task.Title)
	}

}
