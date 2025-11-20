# Aben WC

`aben-wc` is a lightweight command-line utility written in Go, designed to replicate the core functionality of the standard Unix `wc` (word count) tool. It provides a simple and efficient way to analyze files directly from your terminal.

## Features

-   **Byte Count**: Calculate the size of a file in bytes using the `-c` flag.
-   **Line Count**: Count the number of lines in a file using the `-l` flag.
-   **Word Count**: Count the number of words in a file using the `-w` flag.
-   **Character Count**: Count the number of characters in a file using the `-m` flag.
-   **Default Mode**: Display line, word, and byte counts when no flag is provided.
-   **Standard Input**: Support reading from standard input (stdin) via pipes.

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

Run the built executable with the desired flag and the target file path, or pipe content to it.

### Read from Standard Input
To analyze content piped from another command:
```bash
cat test.txt | ./abenwc -l
```
**Example:**
```bash
cat test.txt | ./abenwc -l
# Output: 7145
```

### Default (All Counts)
To display line, word, and byte counts:
```bash
./abenwc <filename>
```
**Example:**
```bash
./abenwc test.txt
# Output: 7145  58164 342190 test.txt
```

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

### Count Words
To display the number of words in a file:
```bash
./abenwc -w <filename>
```
**Example:**
```bash
./abenwc -w test.txt
# Output: 58164 test.txt
```

### Count Characters
To display the number of characters in a file:
```bash
./abenwc -m <filename>
```
**Example:**
```bash
./abenwc -m test.txt
# Output: 339292 test.txt
```

## Running Tests

To run the unit tests for the project:
```bash
go test ./tests/...
```
