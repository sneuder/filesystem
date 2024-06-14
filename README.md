# Filesystem Package

The `filesystem` package provides utilities for file system operations in Go. It includes functions for creating, reading, writing, and managing files and directories.

## Installation

To install the package, use the following command:

```bash
go get github.com/yourusername/filesystem@<version>
```
Replace <version> with:
- v1.0.0


## Usage

### Build a file

It is main feature of package where you can build a file with simple steps.

```go
package main

import (
    "github.com/yourusername/filesystem"
    "log"
)

func main() {
    fileInfo := filesystem.FileInfo{
        Name: "example.txt",
        Path: "/path/to/directory",
        Directives: []filesystem.FileDirective{
            {Content: "Hello, World!", Indent: 0, NewLine: false},
        },
    }

    file, err := filesystem.BuildFile(fileInfo, false)

    if err != nil {
        log.Fatalf(err)
    }

    log.Println("File built successfully.")
}
```

With <code>BuildFile</code> is not required to open the file and close, you just need pass the parameters based on your needs. However, if you want to have a different control, the package has the following methods:

### Open

```go
func main() {
    filePath := "/path/to/directory/example.txt"
    file, err := filesystem.Open("example.txt", "/path/to/directory")
    if err != nil {
        log.Fatalf(err)
    }

    defer filesystem.Close(file)
    //...
}
```

### Read file
With this methd, if the file does not exist, the <code>Open</code> func creates it.

```go
func main() {
    filePath := "/path/to/directory/example.txt"
    content, err := filesystem.Read(filePath)

    if err != nil {
        log.Fatalf(err)
        return
    }

    log.Println("File content:", content)
}
```

### Empty file

```go
func main() {
    filePath := "/path/to/directory/example.txt"
    file, err := filesystem.Open("example.txt", "/path/to/directory")
    if err != nil {
        log.Fatalf(err)
    }

    defer filesystem.Close(file)

    isEmpty, err := filesystem.IsEmpty(file)
    if err != nil {
        log.Fatalf(err)
    }

    if isEmpty {
        log.Println("File is empty.")
    } else {
        log.Println("File is not empty.")
    }
}
```

### File exists

```go
func main() {
    fileName := "example.txt"
    exists := filesystem.Exists(fileName)
    if exists {
        log.Printf("File %s exists.", fileName)
    } else {
        log.Printf("File %s does not exist.", fileName)
    }
}
```

### Add directive or directives

```go
func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
    
	defer file.Close()

	directive := filesystem.FileDirective{
		Content: "This is a directive",
		Indent:  4,
		NewLine: true,
	}

	if err := filesystem.AddDirective(file, directive, false); err != nil {
        log.Fatalf(err)
	}

	log.Println("Directive added successfully.")
}
```

```go
func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatalf(err)
	}

    defer file.Close()

	directives := []filesystem.FileDirective{
		{Content: "First directive", Indent: 0, NewLine: true},
		{Content: "Second directive", Indent: 2, NewLine: true},
		{Content: "Third directive", Indent: 4, NewLine: false},
	}

	if err := filesystem.AddDirectives(file, directives); err != nil {
		log.Fatalf(err)
	}

	log.Println("Directives added successfully.")
}
