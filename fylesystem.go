package fileService

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
	f, err := os.Create(fullPathFile)

	if err != nil {
		return f, err
	}

	return f, err
}

func Write(text string, file *os.File) {
	data := []byte(text)
	var content []byte

	if IsEmpty(file) {
		content = data
	} else {
		content = append([]byte("\n"), data...)
	}

	_, err := file.Write(content)

	if err != nil {
		return
	}
}

func Read(filePath string) string {
	data, _ := os.ReadFile(filePath)
	return string(data)
}

func IsEmpty(file *os.File) bool {
	fileInfo, err := file.Stat()

	if err != nil {
		return false
	}

	return fileInfo.Size() == 0
}

func Close(file *os.File) {
	file.Close()
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func Remove(fileName string) {
	err := os.Remove(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func AddDirective(dockerFile *os.File, directive FileDirectives) {
	indent := getIndent(directive.Indent)
	content := indent + directive.Content
	Write(content, dockerFile)
}

//

func getIndent(indentCount int) string {
	indent := strings.Repeat(" ", indentCount)
	return indent
}
