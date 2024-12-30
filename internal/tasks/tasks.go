package tasks

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
	"github.com/munraitoo13/godo-list/internal/utils"
)

// status enum
type Status int

// status constants
const (
	Pending Status = iota
	Doing
	Done
)

// task struct
type Task struct {
	Title       string
	Description string
	CreatedAt   time.Time
	CompletedAt *time.Time
	Status      Status
}

// tasks struct (slice of tasks)
type Tasks []Task

// validate task index
func (tasks *Tasks) ValidateTaskIndex(index int) error {
	if index < 0 || index >= len(*tasks) {
		return errors.New("invalid task index")
	}

	return nil
}

// convert status to string
func (status Status) ToString() string {
	switch status {
	case Pending:
		return "Pending"
	case Doing:
		return "Doing"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

// read input
func readInput(prompt string) (string, error) {
	// starts the reader
	reader := bufio.NewReader(os.Stdin)

	// command prompt and input reader
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	// returns the input already trimmed
	return strings.TrimSpace(input), err
}

// read index input
func (tasks *Tasks) readIndex(prompt string) (int, error) {
	// initialize index varibale
	var index int

	// command prompt and input reader
	fmt.Print(prompt)
	_, err := fmt.Scanln(&index)
	if err != nil || index < 0 || index >= len(*tasks) {
		return index, errors.New("no valid index provided")
	}

	// returns the index and no error
	return index, nil
}

// show all tasks
func (tasks *Tasks) List() {
	fmt.Println("")

	// create new table
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Status", "Created At", "Completed At")

	// loops through each task and creates a row for it
	for index, t := range *tasks {
		// completed tasks values
		completedAt := ""
		if t.CompletedAt != nil {
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}

		// adds new row with provided data
		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			t.Status.ToString(),
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	// render the table
	table.Render()
}

// view task
func (tasks *Tasks) View() error {
	// dereference tasks slice to access its values
	task := *tasks

	// task index input
	index, _ := tasks.readIndex("Select the task number to be viewed: ")

	// clears screen
	utils.ClearScreen()

	// show selected task's title
	fmt.Println("Title:")
	fmt.Println(task[index].Title)

	// show selected task's description
	fmt.Println("\nDescription:")
	fmt.Println(task[index].Description)

	// return no error
	return nil
}

// add new task
func (tasks *Tasks) CreateTask() error {
	// task's title and decription
	title, _ := readInput("Task's title: ")
	description, _ := readInput("Task's description: ")

	// set the local task values
	task := Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
		Status:      Pending,
	}

	// appends the local task value to the slice of tasks
	*tasks = append(*tasks, task)

	// clears screen, success message, returns no errors
	utils.ClearScreen()
	fmt.Println("Task created successfully!")
	return nil
}

// edit task
func (tasks *Tasks) EditTask() error {
	// dereference tasks slice to access its values
	task := *tasks

	// reads index input
	index, _ := tasks.readIndex("Select the task to be edited: ")

	// title and description inputs
	title, _ := readInput("New title: ")
	description, _ := readInput("New description: ")

	// set the task's title and description
	task[index].Title = title
	task[index].Description = description

	// clears screen, success message, returns no errors
	utils.ClearScreen()
	fmt.Println("Task edited successfully!")
	return nil
}

// set task status
func (tasks *Tasks) SetStatus() error {
	// dereference tasks slice to access its values
	task := *tasks

	// task index input
	index, _ := tasks.readIndex("Task index to be edited: ")

	// new status
	statusInput, _ := readInput("New task status: ")

	// switch to change the task's status based on input
	switch strings.ToLower(statusInput) {
	case "pending":
		task[index].Status = Pending

		task[index].CompletedAt = nil
	case "doing":
		task[index].Status = Doing

		task[index].CompletedAt = nil
	case "done":
		task[index].Status = Done

		completionTime := time.Now()
		task[index].CompletedAt = &completionTime
	default:
		return errors.New("invalid status")
	}

	// clears screen, success message, returns no errors
	utils.ClearScreen()
	fmt.Println("Task status updated successfully!")
	return nil
}

// delete task
func (tasks *Tasks) DeleteTask() error {
	// dereference tasks slice to access its values
	task := *tasks

	// task index input
	index, _ := tasks.readIndex("Task index to be deleted: ")

	// appends all tasks except the one to be deleted
	*tasks = append(task[:index], task[index+1:]...)

	// clears screen, success message, returns no errors
	utils.ClearScreen()
	fmt.Println("Task deleted successfully!")
	return nil
}

// delete all tasks
func (tasks *Tasks) DeleteAllTasks() error {
	// confirmation input
	confirmation, _ := readInput("Are you sure you want to delete all tasks? (yes/no): ")

	// switch to check if the user wants to delete all tasks
	switch strings.ToLower(confirmation) {
	case "yes":
		utils.ClearScreen()

		fmt.Println("All tasks deleted.")
		*tasks = []Task{}
	case "no":
		utils.ClearScreen()

		fmt.Println("Nothing was deleted.")
	default:
		utils.ClearScreen()

		fmt.Println("Invalid option.")
	}

	// returns no errors
	return nil
}
