package filesystem

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type FileInfo struct {
	Name       string
	Path       string
	Directives []FileDirectives
}

type FileDirectives struct {
	Content string
	Indent  int
}

func Open(fileName string, pathDir string) (*os.File, error) {
	fullPathFile := path.Join(pathDir, fileName)
	file, err := os.Create(fullPathFile)

	if err != nil {
		err = fmt.Errorf("failed to open file %s: %w", fileName, err)
	}

	return file, err
}

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

func Read(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("failed to read file with path %s: %w", filePath, err)
	}
	return string(data), err
}

func IsEmpty(file *os.File) (bool, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		err = fmt.Errorf("failed to check file info %s: %w", file.Name(), err)
	}
	return fileInfo.Size() == 0, err
}

func Close(file *os.File) error {
	err := file.Close()
	if err != nil {
		err = fmt.Errorf("failed to close file %s: %w", file.Name(), err)
	}
	return err
}

func Exists(filename string) bool {
	exists := true
	_, err := os.Stat(filename)
	if err != nil {
		exists = false
	}

	return exists
}

func Remove(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		err = fmt.Errorf("failed to remove file %s: %w", fileName, err)
	}
	return err
}

func AddDirective(file *os.File, directive FileDirectives) {
	newLine := true

	indent := getIndent(directive.Indent)
	content := indent + directive.Content

	if newLine {
		content = content + "\n"
	}

	Write(file, content)
}

//

func getIndent(indentCount int) string {
	indent := strings.Repeat(" ", indentCount)
	return indent
}
