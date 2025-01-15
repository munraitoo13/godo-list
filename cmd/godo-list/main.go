package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/munraitoo13/godo-list/internal/storage"
	"github.com/munraitoo13/godo-list/internal/tasks"
	"github.com/munraitoo13/godo-list/internal/utils"
)

func main() {
	// clears the screen when stating cli and show message
	utils.ClearScreen()
	fmt.Println("Welcome to godo-list! Go to github.com/munraitoo13/godo-list to view all commands.")

	// initialize tasks slice
	taskList := tasks.Tasks{}
	// initialize storage with tasks type and tasks json file
	storage := storage.NewStorage[tasks.Tasks]("tasks.json")
	// load tasks from file and store in taskList
	storage.LoadTasks(&taskList)

	// initialize a reader to read user input
	reader := bufio.NewReader(os.Stdin)

	// repl loop to read user input and execute commands
	for {
		// every iteration, list all tasks and print cursor
		taskList.List()
		fmt.Print("\n> ")

		// read user command and remove trailing newline
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command: ", err)
			continue
		}
		command = strings.TrimSpace(command)

		// switch on command
		switch command {
		case "view":
			utils.ClearScreen()

			fmt.Println("Tasks: ")
			taskList.List()

			if err := taskList.View(); err != nil {
				utils.ClearScreen()

				fmt.Println("Error:", err)
			}
		case "add":
			utils.ClearScreen()

			taskList.CreateTask()
			storage.SaveTasks(taskList)
		case "edit":
			utils.ClearScreen()

			fmt.Println("Tasks:")
			taskList.List()

			if err := taskList.EditTask(); err != nil {
				utils.ClearScreen()

				fmt.Println("Error:", err)
			} else {
				storage.SaveTasks(taskList)
			}
		case "status":
			utils.ClearScreen()

			fmt.Println("Tasks:")
			taskList.List()

			if err := taskList.SetStatus(); err != nil {
				utils.ClearScreen()

				fmt.Println("Error:", err)
			} else {
				storage.SaveTasks(taskList)
			}
		case "delete":
			utils.ClearScreen()

			fmt.Println("Tasks:")
			taskList.List()

			if err := taskList.DeleteTask(); err != nil {
				utils.ClearScreen()

				fmt.Println("Error:", err)
			} else {
				storage.SaveTasks(taskList)
			}
		case "delete all":
			utils.ClearScreen()

			taskList.DeleteAllTasks()
			storage.SaveTasks(taskList)
		case "clear":
			utils.ClearScreen()

			fmt.Println("Welcome to godo-list! Go to github.com/munraitoo13/godo-list to view all commands.")
		case "exit":
			utils.ClearScreen()

			fmt.Println("Goodbye!")
			fmt.Println("")

			return
		default:
			utils.ClearScreen()

			fmt.Println("Unknown command")
		}
	}
}
