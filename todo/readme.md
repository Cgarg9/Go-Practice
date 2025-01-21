# Todo CLI Application

## Overview
The **Todo CLI Application** is a command-line tool for managing a list of tasks (todos). It provides functionality to add, edit, delete, toggle, and list todos efficiently using Go. The application is implemented across multiple files: `main.go`, `todo.go`, `storage.go`, and `command.go`, each serving a specific purpose.

---

## Features
1. **Add a Todo**: Create a new task with a title.
2. **Edit a Todo**: Update the title of an existing task.
3. **Delete a Todo**: Remove a task by its index.
4. **Toggle a Todo**: Mark a task as completed or uncompleted.
5. **List All Todos**: Display all tasks with their details.

---

## File Breakdown

### 1. `main.go`
This is the entry point of the application. It initializes the todo list and processes user commands.

#### Key Functionality:
- Initializes the `Todos` data structure.
- Calls the `NewCmdFlags()` function to parse user commands.
- Executes the respective command using `cmdFlags.Execute()`.

#### Code Snippet:
```go
package main

func main() {
    todos := &Todos{}
    cmdFlags := NewCmdFlags() // Parse command-line flags
    cmdFlags.Execute(todos)   // Execute commands based on flags
}

### 2. `todo.go`
Defines the core Todo structure and implements methods for managing todos.

#### Structures:
- Todo: Represents a task with fields for title, completion status, creation time, and completion time.
- Todos: A slice of Todo objects.

#### Methods:
- add(title string): Adds a new todo to the list.
- validateIndex(index int): Validates if the given index is within bounds.
- delete(index int): Removes a todo at the specified index.
- edit(index int, newTitle string): Updates the title of a todo.
- toggle(index int): Toggles the completion status of a todo.
- print(): Displays all todos with their details.