# CLI ToDo App

A simple command-line ToDo application written in Go. Manage your tasks easily from your terminal.

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/CLI-ToDo-App.git
   cd CLI-ToDo-App
   ```

2. **Build the application:**
   ```sh
   go build -o todo
   ```

   This will create an executable named `todo` in the current directory.

3. **(Optional) Make `todo` accessible from anywhere:**

   Move the `todo` executable to a directory that's in your `PATH`, such as `/usr/local/bin`:

   ```sh
   sudo mv todo /usr/local/bin/
   ```

   Now you can run `todo` from any location in your terminal.

## Usage

### Add a new task

```sh
todo add "Your task description"
```

### List all tasks

```sh
todo list
```

### Mark a task as complete

```sh
todo complete <task_id>
```

Replace `<task_id>` with the ID of the task you want to mark as complete (as shown in the list).

## Example

```sh
todo add "Buy groceries"
todo add "Read a book"
todo list
todo complete 1
todo list
```

## Requirements

- Go 1.18 or higher

## License

MIT License
