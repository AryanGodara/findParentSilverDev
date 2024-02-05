# File System Hierarchy Project

## Introduction

This project is implemented in Go, a statically typed, compiled programming language designed at Google. I chose Go for its simplicity, efficiency, and excellent support for concurrency. These attributes make Go a fitting choice for developing systems-level programs, such as a file system hierarchy manager. Go's powerful standard library, particularly its data structure and error handling capabilities, allows for robust and efficient manipulation of complex hierarchical data structures like file systems. Moreover, Go's strong type system and compilation checks ensure the reliability and correctness of the program, which is critical for managing file system operations.

```
And now, the real reason is that I'm a Go enthusiast and I love to code in Go 😄
Plus, since it's easier to read the code, even for a non-Go developer, I thought it would be a good choice for this project.

But I'd love to create a CLI tool for this in Rust, using the actual filesystem and not a simulated one.
```

## Project Overview

The project provides a simplistic simulation of a file system hierarchy. It includes functionalities to create files, establish hierarchical relationships between them (parent-child), create aliases for files, and find files by name. It also implements a method to determine the closest common parent directory of two given files, which is a common operation in file systems for path resolution and permission checks.

### Key Components

- **File Structure**: Represents a file or a directory in the file system. Each file can have children (sub-files or sub-directories) and aliases (alternative names or links to other files).

- **Error Handling**: Defines custom errors to handle specific situations, such as when a file is not found or a nil reference is encountered.

- **File Operations**:
  - `addChild`: Adds a child file to the current file, simulating the creation of a new file or directory within a directory.
  - `addAlias`: Creates an alias for a file, allowing it to be referenced by multiple names or paths.
  - `findFileByName`: Searches for a file by its name, considering both direct names and aliases. It recursively searches through children and aliases to find the specified file.
  - `findParent`: Determines the closest common parent directory of two files. It uses a helper function to find the paths from the root to each file and then compares these paths to find the common parent.

### Testing

The project includes a test suite written with the `testing` package and the `testify` assertion library for Go. The tests cover various scenarios, including:

- Finding the common parent of two files in different branches of the file system.
- Handling aliases in the search for common parents.
- Error handling for nil file references.

The test cases are designed to be exhaustive, covering not only basic parent-child relationships but also more complex interactions involving aliases and error conditions. This ensures the robustness of the file system operations against a wide range of inputs and scenarios.

## Conclusion

This Go project demonstrates a fundamental approach to managing a file system hierarchy, including creating files and directories, establishing relationships among them, and resolving common ancestors. The choice of Go as the implementation language underscores the importance of efficiency, reliability, and simplicity in systems programming. The comprehensive test suite further ensures the correctness and stability of the core functionalities, making this project a solid foundation for more advanced file system operations and management tasks.