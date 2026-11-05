# Chappie CLI

Chappie is a simple, customizable command-line utility designed to run specific actions on a scheduled basis. It is intended to be expanded over time for various functionalities.

## Features

*   **Immediate Action:** Executes a defined initial action ("hello").
*   **Scheduled Execution:** Runs a defined action once per hour until the end of the calendar day.
*   **Extensible Design:** The core logic is separated from the main execution loop, allowing easy replacement of the current "hello" action with complex future tasks.

## Setup

### Prerequisites

You must have Go (Golang) installed on your system.

### Installation

1.  **Clone the repository:**
    ```bash
    git clone git@github.com-personal:ravirajsubramanian/chappie.git
    cd chappie
    ```
2.  **Build the executable:**
    ```bash
    go build -o bin main.go
    ```

## Usage

### Running Chappie

To run the CLI:
```bash
./bin/main cli
```

**Note:** The program will execute its initial action immediately, and then continue printing "hello" every hour until the midnight of the current day.

## Future Development

The `runAction()` function within `main.go` is the primary hook for extending Chappie's functionality. Any major feature should be implemented here.

---
*This documentation file is auto-generated.*
