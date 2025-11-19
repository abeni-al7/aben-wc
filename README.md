# Aben WC

`aben-wc` is a lightweight command-line utility written in Go, designed to replicate the core functionality of the standard Unix `wc` (word count) tool. It provides a simple and efficient way to analyze files directly from your terminal.

## Features

-   **Byte Count**: Calculate the size of a file in bytes using the `-c` flag.
-   **Line Count**: Count the number of lines in a file using the `-l` flag.

## Architecture

For a detailed overview of the project's design, components, and data flow, please refer to the [Architecture Documentation](docs/architecture.md).

## Getting Started

### Prerequisites

-   [Go](https://go.dev/dl/) (version 1.18 or higher recommended)

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/abeni-al7/aben-wc.git
    cd aben-wc
    ```

2.  Build the application:
    ```bash
    go build -o abenwc main.go
    ```

## Usage

Run the built executable with the desired flag and the target file path.

### Count Bytes
To display the number of bytes in a file:
```bash
./abenwc -c <filename>
```
**Example:**
```bash
./abenwc -c test.txt
# Output: 342190 test.txt
```

### Count Lines
To display the number of lines in a file:
```bash
./abenwc -l <filename>
```
**Example:**
```bash
./abenwc -l test.txt
# Output: 7145 test.txt
```

## Running Tests

To run the unit tests for the project:
```bash
go test ./tests/...
```
