package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	var tasks TaskFile
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <filename>")
		return
	}
	filename := os.Args[1]

	data, err := os.ReadFile(filename)

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
