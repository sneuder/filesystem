# Filesystem Package

The `filesystem` package provides utilities for file system operations in Go. It includes functions for creating, reading, writing, and managing files and directories.

## Installation

To install the package, use the following command:

```bash
go get github.com/yourusername/filesystem@<version>
```
Replace <version> with your prefered version.


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
```

## JSON Files

### Create JSON File

```go
type TestData struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

func main() {
    data := TestData{Name: "example", Value: 42}
    path := "/path/to/directory"
    fileName := "test_file"

    err := filesystem.CreateJSONFile(data, path, fileName)
    if err != nil {
        log.Fatalf("Error creating JSON file: %v", err)
    }

    log.Println("JSON file created successfully.")
}
```

### Read JSON File

```go
type TestData struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

func main() {
    path := "/path/to/directory"
    fileName := "test_file"

    var data TestData
    err := filesystem.ReadJSONFile(path, fileName, &data)
    if err != nil {
        log.Fatalf("Error reading JSON file: %v", err)
    }

    log.Printf("Read data: Name = %s, Value = %d", data.Name, data.Value)
}
```

### Check JSON File

```go
func main() {
    path := "/path/to/directory"
    fileName := "test_file"

    exists := filesystem.CheckJSONFile(path, fileName)
    if exists {
        log.Printf("File %s exists.", fileName)
    } else {
        log.Printf("File %s does not exist.", fileName)
    }
}
```

### Remove JSON File

```go
func main() {
    path := "/path/to/directory"
    fileName := "test_file"

    removed := filesystem.RemoveJSONFile(fileName, path)
    if removed {
        log.Println("JSON file removed successfully.")
    } else {
        log.Println("Failed to remove JSON file.")
    }
}
```

### Add JSON File Extension

```go
func main() {
    fileName := "test_file"
    fullFileName := filesystem.AddJSONExtension(fileName)
    
    log.Printf("Full file name: %s", fullFileName)
}
```