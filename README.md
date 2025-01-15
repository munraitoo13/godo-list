


<div align="center">
  <h1>Godo List</h1>
  <p>
    Welcome to <b>Godo List</b>, a lightweight and efficient CLI-based task manager built entirely in Go. 
    Manage your tasks with ease, right from your terminal.
  </p>
  <p>
    <img width="50%" src="https://img.icons8.com/?size=512&id=44442&format=png" alt="Gopher" />
  </p>
</div>

<div align="center">
  <h2>Built With</h2>
  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
</div>

<h2 align="center">Built With</h2>

- Lightweight and fast
- Easy-to-use CLI commands
- JSON-based storage for portability
- Minimalistic and intuitive design

<h2 align="center">Getting Started</h2>

### How Do I Download the CLI?

Simply head to the [releases](https://github.com/munraitoo13/godo-list/releases) page and download the latest version!

### How Do I Use It?

Here are the available commands:

- **View Task**: To view details about a task, type:


  ```bash
  view, then <task_index>
  ```
- **Add Task**: Add a new task with:


  ```bash
  add, then <task_description>
  ```
- **Edit Task**: Modify an existing task using:


  ```bash
  edit, then <task_index>
  ```
- **Change Task's Status**: Update the status of a task with:


  ```bash
  status, then <task_index> and <done || doing || pending>
  ```
- **Delete Task**: Remove a task by typing:


  ```bash
  delete, then <task_index>
  ```
- **Delete All Tasks**: Start fresh with:


  ```bash
  delete all
  ```
- **Clear Terminal**: Clear the screen using:


  ```bash
  clear
  ```
- **Exit**: Exit the application with:


  ```bash
  exit
  ```

### Building the Project Yourself

#### Prerequisites

- [Git](https://git-scm.com/downloads)
- [Go](https://go.dev/doc/install)

#### Steps to Build

1. Clone the repository:


   ```bash
   git clone https://github.com/munraitoo13/godo-list
   ```
2. Build the project and run the binary:


   ```bash
   go build ./cmd
   ./cmd
   ```
   
   Or, run directly without building:

   
   ```bash
   go run ./cmd
   ```

<h2 align="center">Project Structure</h2>

```plaintext
godo-list/
├── cmd/           # Main entry point for the CLI application
├── internal/      # Core logic and helper functions
├── go.mod         # Dependency and module definitions
├── go.sum         # Project dependencies
└── tasks.json     # Storage file for tasks
```

<h2 align="center">Contributing</h2>

<p>
  Contributions are what make the open-source community an amazing place to learn and create. Any contributions will be greatly appreciated.
</p>
<p>
  If you have a suggestion that would make this project better, please fork the repo and create a pull request. You can also simply open an issue with the <b>suggestion</b> tag. You can find more info on how to contribute to a project <a href="https://git-scm.com/book/en/v2/GitHub-Contributing-to-a-Project">here</a>.
</p>

<div align="center">
  <h2>Contact</h2>
  <b>Augusto Mendes</b><br>
  <a href="https://linktr.ee/munraitoo13">@munraitoo13</a> on all socials or click on my @.<br>
  Consider giving this project a star and thanks for your attention. Cheers!
</div>
