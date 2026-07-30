# TODO

A simple terminal-based task manager written in Go. Tasks are saved to a JSON file, so the program picks up where it left off after being closed and reopened.

## Features

- Add, list, edit, and delete tasks
- Mark tasks as completed
- Confirmation prompt before deletion
- Automatic JSON persistence — every change is written immediately
- Color-coded status (completed in green, pending in yellow)
- Total / completed / pending counts in the list view
- Input validation (empty titles, invalid IDs, non-numeric menu choices)

## Installation

Requires Go 1.26 or later.

```bash
git clone https://github.com/ucanalgan/TODO.git
cd TODO
go run .
```

To build an executable:

```bash
go build .
```

> `go run main.go` will not work — the project is split across multiple files, so use `go run .` to compile the whole package.

## Usage

The menu appears when the program starts:

```
1. Add Mission
2. List Missions
3. Complete Mission
4. Delete Mission
5. Edit Mission
6. Exit
Enter your choice:
```

Sample list output:

```
Missions:
Mission ID: 1, Title: Read the Go documentation, Status: Completed
Mission ID: 2, Title: Write the README, Status: Incomplete
Total Missions: 2, Completed: 1, Incomplete: 1
```

Tasks are stored in `missions.json` in the working directory. If the file does not exist, the program starts with an empty list and creates it when the first task is added.

## Project Structure

```
.
├── main.go       # Menu loop and command dispatch
├── mission.go    # Mission data model
├── input.go      # Terminal input reading and validation
├── storage.go    # Saving to and loading from JSON
└── missions.json # Task data (created automatically, not tracked in version control)
```

All four files belong to `package main`; splitting into separate packages was considered unnecessary for a project of this size.

### Data model

```go
type Mission struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```

IDs continue from one above the highest existing ID at startup, which prevents ID collisions after a deletion.

## Known Limitations

- Task titles are single-line only
- Deleted tasks cannot be recovered (a confirmation is requested before deleting)
- Concurrent instances of the program may conflict when writing to the file

## Roadmap

- [ ] Priority field and sorting by priority
- [ ] Hiding / filtering completed tasks
- [ ] Search by title
- [ ] Creation timestamp
- [ ] CLI commands instead of a menu (`todo add`, `todo list`)
- [ ] Unit tests
