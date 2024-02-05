# File System Hierarchy Project

## Introduction

This project is implemented in Go, a statically typed, compiled programming language designed at Google. I chose Go for its simplicity, efficiency, and excellent support for concurrency. These attributes make Go a fitting choice for developing systems-level programs, such as a file system hierarchy manager. Go's powerful standard library, particularly its data structure and error handling capabilities, allows for robust and efficient manipulation of complex hierarchical data structures like file systems. Moreover, Go's strong type system and compilation checks ensure the reliability and correctness of the program, which is critical for managing file system operations.

```
And now, the real reason is that I'm a Go enthusiast and I love to code in Go 😄
Plus, since it's easier to read the code, even for a non-Go developer, I thought it would be a good choice for this project.

But I'd love to create a CLI tool for this in Rust, using the actual filesystem and not a simulated one.
```

## Approach Overview
#### First Iteration
I startd out with a basic file struct, recursively holding a list of children and aliases. I then implemented the basic file operations, such as adding children and aliases, and searching for files by name. 

#### Second Iteration
Then defined custom errors to handle specific situations, such as when a file is not found or a nil reference is encountered. I then implemented a helper function to find the closest common parent directory of two files. This function uses a recursive approach to find the paths from the root to each file and then compares these paths to find the common parent. 

#### Third Iteration
After this, Iadded a test suite to cover various scenarios, including finding the common parent of two files in different branches of the file system, handling aliases and soft links in the search for common parents, and error handling for nil file references.

#### Fourth Iteration
Added support for Aliases and Sys links. This was a bit tricky, but I managed to get it working. I also added a few more tests to cover these new functionalities.

#### Fifth Iteration (Improving Time Compexity of findPath function to O(log n))
I add an extra field 'parent' of type *file, this way, given a file, we can find its parent in logarithmic time. This way, the time complexity of the findPath function is reduced to O(log n) from O(n).

I decided `against` using precomputation or memoization to store paths, since it didn't make sense to use the extra storage for this task without knowledge of the type of system it'll be running on.

#### Final Iteration
I added a few more tests to cover the new functionalities and also refactored the code a bit to make it more readable and maintainable.

### Key Components

- **File Structure**: Represents a file or a directory in the file system. Each file can have children (sub-files or sub-directories) and aliases (alternative names or links to other files).

- **Error Handling**: Defines custom errors to handle specific situations, such as when a file is not found or a nil reference is encountered.

- **File Operations**:
  - `addChild`: Adds a child file to the current file, simulating the creation of a new file or directory within a directory.
  - `addAlias`: Creates an alias for a file, allowing it to be referenced by multiple names or paths.
  - `findFileByName`: Searches for a file by its name, considering both direct names and aliases. It recursively searches through children and aliases to find the specified file.
  - `findParent`: Determines the closest common parent directory of two files. It uses a helper function to find the paths from the root to each file and then compares these paths to find the common parent.

### Softlinks vs Aliases (according to me)

The concepts of "alias" and "soft link" (or symbolic link) come from different contexts but share the idea of referring to another entity. However, they operate at different levels and in somewhat different manners:

#### Alias
**Context**: Aliases are often used in various software applications, command shells, and programming environments to refer to another command, function, or sometimes file paths with a simpler or alternative name. 

**Functionality**: An alias is essentially a shortcut or a nickname to another command or entity. In command shells, for example, an alias might allow a user to execute a long command using a short nickname. 

**Scope**: Typically, the use of an alias is limited to the context or environment in which it is defined, such as a shell session or a specific application. 

#### Soft Link (Symbolic Link)
**Context**: Soft links are used within file systems to create a reference to another file or directory. 

**Functionality**: A soft link is a special type of file that points to another file or directory path. When you access a soft link, the file system redirects you to the file or directory it points to. Unlike hard links, soft links can cross file systems and can link to directories. 

**Durability**: Soft links remain valid even if the target file or directory is moved within the same file system (though they will break if the target is deleted or moved to a different file system). They are more flexible than hard links, which cannot link directories or cross file systems.
File System Level: Soft links operate at the file system level, meaning they are recognized and resolved by the operating system's file system. 

In essence, while both aliases and soft links provide a way to refer to another entity using a different name or path, aliases are more about convenience within a specific context (like a command shell or an application), and soft links are about creating a navigable reference within a file system that points to another file or directory. Soft links are more versatile and integrated into the file system's structure, allowing for more complex file and directory arrangements and access patterns.

### Testing

The project includes a test suite. The tests cover various scenarios, including:

- Finding the common parent of two files in different branches of the file system.
- Handling aliases and soft link in the search for common parents.
- Error handling for nil file references.

The test cases are designed to be exhaustive, covering not only basic parent-child relationships but also more complex interactions involving aliases and error conditions. This ensures the robustness of the file system operations against a wide range of inputs and scenarios.

## Conclusion

This demonstrates a fundamental approach to managing a file system hierarchy, including creating files and directories, establishing relationships among them, and resolving common ancestors. The choice of Go as the implementation language underscores the importance of efficiency, reliability, and simplicity in systems programming. The comprehensive test suite further ensures the correctness and stability of the core functionalities, making this project a solid foundation for more advanced file system operations and management tasks.