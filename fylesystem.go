// Package filesystem provides utilities for file system operations.
package filesystem

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// FileInfo holds metadata about a file.
type FileInfo struct {
	Name       string          // Name is the name of the file.
	Path       string          // Path is the directory path where the file is located.
	Directives []FileDirective // Directives are instructions for file content.
}

// FileDirective represents a directive with content, indentation, and newline information.
type FileDirective struct {
	Content string // Content is the text of the directive.
	Indent  int    // Indent is the number of spaces to indent the content.
	NewLine bool   // NewLine indicates whether to add a newline before the content.
}

// Open creates and opens a file with the given name in the specified directory.
// It returns a pointer to the file and an error if one occurred.
func Open(fileName string, pathDir string) (*os.File, error) {
	fullPathFile := path.Join(pathDir, fileName)
	file, err := os.Create(fullPathFile)

	if err != nil {
		err = fmt.Errorf("failed to open file %s: %w", fileName, err)
	}

	return file, err
}

// Write writes the given text to the file. If the file is not empty, the content is appended.
// It returns an error if one occurred.
func Write(file *os.File, text string) error {
	data := []byte(text)
	var content []byte

	isEmpty, _ := IsEmpty(file)

	if isEmpty {
		content = data
	} else {
		content = append([]byte(""), data...)
	}

	_, err := file.Write(content)

	if err != nil {
		err = fmt.Errorf("failed to write file %s: %w", file.Name(), err)
		return err
	}

	return err
}

// Read reads the content of the file at the given path and returns it as a string.
// It returns an error if one occurred.
func Read(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("failed to read file with path %s: %w", filePath, err)
	}
	return string(data), err
}

// IsEmpty checks if the file is empty. It returns a boolean indicating whether the file is empty
// and an error if one occurred.
func IsEmpty(file *os.File) (bool, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		err = fmt.Errorf("failed to check file info %s: %w", file.Name(), err)
	}
	return fileInfo.Size() == 0, err
}

// Close closes the given file. It returns an error if one occurred.
func Close(file *os.File) error {
	err := file.Close()
	if err != nil {
		err = fmt.Errorf("failed to close file %s: %w", file.Name(), err)
	}
	return err
}

// Exists checks if the file with the given name exists. It returns true if the file exists, false otherwise.
func Exists(filename string) bool {
	exists := true
	_, err := os.Stat(filename)
	if err != nil {
		exists = false
	}

	return exists
}

// Remove deletes the file with the given name. It returns an error if one occurred.
func Remove(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		err = fmt.Errorf("failed to remove file %s: %w", fileName, err)
	}
	return err
}

// AddDirective adds a directive to the file with the specified indentation and newline settings.
// If isLast one is true it does not add new line at the end.
// It returns an error if one occurred.
func AddDirective(file *os.File, directive FileDirective, isLast bool) error {
	indent := getIndent(directive.Indent)
	content := indent + directive.Content

	if !isLast {
		content += "\n"
	}

	if directive.NewLine {
		content = "\n" + content
	}

	if err := Write(file, content); err != nil {
		return fmt.Errorf("failed to add directive %s: %w", directive.Content, err)
	}

	return nil
}

// AddDirectives adds multiple directives to the file. It returns an error if one occurred.
func AddDirectives(file *os.File, directives []FileDirective) error {
	for i, directive := range directives {
		isLast := len(directives) == (i + 1)
		if err := AddDirective(file, directive, isLast); err != nil {
			return err
		}
	}

	return nil
}

// BuildFile creates and opens a file with the specified FileInfo and adds the directives.
// If closeFile is true, the file is closed after the directives are added.
// It returns a pointer to the file and an error if one occurred.
func BuildFile(fileInfo FileInfo, closeFile bool) (*os.File, error) {
	file, err := Open(fileInfo.Name, fileInfo.Path)
	if err != nil {
		return nil, err
	}

	if err = AddDirectives(file, fileInfo.Directives); err != nil {
		return nil, err
	}

	if closeFile {
		if err = Close(file); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return file, nil
}

// getIndent returns a string of spaces for the given indentation count.
func getIndent(indentCount int) string {
	indent := strings.Repeat(" ", indentCount)
	return indent
}
