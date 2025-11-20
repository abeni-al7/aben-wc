# Project Architecture

This document provides a comprehensive overview of the architecture for the `aben-wc` project. The application is a command-line tool written in Go, designed to mimic the functionality of the standard `wc` (word count) utility.

## 1. High-Level Overview

The project follows a layered architecture pattern, separating concerns into:
-   **Entry Point**: Initialization and configuration.
-   **Controller Layer**: Handling user input, command-line arguments, and output formatting.
-   **Service Layer**: Core business logic and file system interactions.

### System Context Diagram

```mermaid
graph LR
    User[User] -- CLI Arguments --> App[aben-wc Application]
    App -- Read Operations --> FS[File System]
    App -- Formatted Output --> Stdout[Standard Output]
    App -- Errors --> Stderr[Standard Error]
```

## 2. Component Design

### 2.1 Entry Point (`main.go`)
The `main` package serves as the composition root. It is responsible for:
-   Instantiating the `FileService`.
-   Injecting the service into the `FileIO` controller.
-   Invoking the controller to start processing user input.

### 2.2 Controller Layer (`controllers/file_io.go`)
The `FileIO` struct acts as the controller. It handles the interaction with the user via the command line.
-   **Responsibilities**:
    -   Parsing command-line flags (e.g., `-c` for byte count, `-l` for line count, `-w` for word count, `-m` for character count).
    -   Handling the default case (no flags) by calculating line, word, and byte counts.
    -   Validating arguments (ensuring a file path is provided).
    -   Routing requests to the appropriate service method.
    -   Formatting and printing the results to `stdout`.
    -   Handling errors and printing them to `stderr`.

### 2.3 Service Layer (`services/file_services.go`)
The `FileService` struct encapsulates the core logic for text analysis. It is stateless and operates on in-memory byte slices, decoupling the logic from file system operations.
-   **`GetFileSize(data []byte) int`**:
    -   Returns the length of the byte slice.
-   **`GetLineCount(data []byte) int`**:
    -   Counts the number of newline characters (`\n`) in the data.
-   **`GetWordCount(data []byte) int`**:
    -   Parses the data as a string and counts the number of fields (words).
-   **`GetCharCount(data []byte) int`**:
    -   Counts the number of UTF-8 runes in the data.

## 3. Data Flow

The following sequence diagram illustrates the control flow when a user executes the application. The controller handles I/O, while the service performs pure calculation.

```mermaid
sequenceDiagram
    participant User
    participant Main as main.go
    participant Controller as FileIO (Controller)
    participant Service as FileService (Service)
    participant Source as File System / Stdin

    User->>Main: Run `abenwc -l file.txt`
    Main->>Controller: AcceptInput()
    Controller->>Controller: parseFlags()
    Controller->>Source: Open file or check Stdin
    Source-->>Controller: Input Source
    Controller->>Source: Read All Data
    Source-->>Controller: data ([]byte)
    
    alt Flag is -c (Byte Count)
        Controller->>Service: GetFileSize(data)
        Service-->>Controller: size (int)
    else Flag is -l (Line Count)
        Controller->>Service: GetLineCount(data)
        Service-->>Controller: count (int)
    else Flag is -w (Word Count)
        Controller->>Service: GetWordCount(data)
        Service-->>Controller: count (int)
    else Flag is -m (Character Count)
        Controller->>Service: GetCharCount(data)
        Service-->>Controller: count (int)
    else No Flag (Default)
        Controller->>Service: GetLineCount(data)
        Service-->>Controller: lines (int)
        Controller->>Service: GetWordCount(data)
        Service-->>Controller: words (int)
        Controller->>Service: GetFileSize(data)
        Service-->>Controller: size (int)
    end

    Controller-->>User: Print Result
```

## 4. Design Decisions

-   **Dependency Injection**: The `FileService` is injected into `FileIO`. This promotes loose coupling and makes the controller easier to test (e.g., by mocking the service).
-   **Standard Library**: The project relies heavily on Go's standard library (`flag`, `os`, `bufio`), avoiding unnecessary external dependencies for core functionality.
-   **Buffered I/O**: `bufio.Scanner` is used for line counting to efficiently handle large files without loading the entire content into memory.
