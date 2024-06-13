package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	defer os.Remove(filepath.Join(pathDir, fileName)) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	if _, err := os.Stat(filepath.Join(pathDir, fileName)); os.IsNotExist(err) {
		t.Fatalf("Expected file to exist")
	}
}

func TestWrite(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	defer os.Remove(filepath.Join(pathDir, fileName)) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	text := "Hello, World!"
	err = Write(file, text)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	file.Close()
	readText, err := Read(filepath.Join(pathDir, fileName))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if readText != text {
		t.Fatalf("Expected %q, got %q", text, readText)
	}
}

func TestRead(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	defer os.Remove(filepath.Join(pathDir, fileName)) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	text := "Hello, World!"
	err = Write(file, text)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	file.Close()
	readText, err := Read(filepath.Join(pathDir, fileName))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if readText != text {
		t.Fatalf("Expected %q, got %q", text, readText)
	}
}

func TestIsEmpty(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	defer os.Remove(filepath.Join(pathDir, fileName)) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	isEmpty, err := IsEmpty(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !isEmpty {
		t.Fatalf("Expected file to be empty")
	}

	Write(file, "Hello")
	isEmpty, err = IsEmpty(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if isEmpty {
		t.Fatalf("Expected file not to be empty")
	}
}

func TestClose(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	defer os.Remove(filepath.Join(pathDir, fileName)) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = Close(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestExists(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	fullPath := filepath.Join(pathDir, fileName)
	defer os.Remove(fullPath) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	file.Close()

	if !Exists(fullPath) {
		t.Fatalf("Expected file to exist")
	}

	os.Remove(fullPath)
	if Exists(fullPath) {
		t.Fatalf("Expected file not to exist")
	}
}

func TestRemove(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	fullPath := filepath.Join(pathDir, fileName)

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	file.Close()

	err = Remove(fullPath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if Exists(fullPath) {
		t.Fatalf("Expected file to be removed")
	}

	err = Remove(fullPath)
	if err == nil {
		t.Fatalf("Expected an error for removing non-existent file")
	}
}

func TestAddDirective(t *testing.T) {
	fileName := "testfile.txt"
	pathDir := os.TempDir()
	fullPath := filepath.Join(pathDir, fileName)
	defer os.Remove(fullPath) // Clean up

	file, err := Open(fileName, pathDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer file.Close()

	directive := FileDirectives{Content: "RUN echo Hello", Indent: 4}
	AddDirective(file, directive)
	file.Close()

	readText, err := Read(fullPath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expectedText := "    RUN echo Hello\n" // Indent with 4 spaces and new line
	if readText != expectedText {
		t.Fatalf("Expected %q, got %q", expectedText, readText)
	}
}

func TestGetIndent(t *testing.T) {
	firstIndent := getIndent(2)
	secondIndent := getIndent(4)
	thirdIndent := getIndent(6)
	fourthIndent := getIndent(8)

	firstExpectedIndent := "  "
	if firstIndent != firstExpectedIndent {
		t.Fatalf("Expected %q, got %q", firstExpectedIndent, firstIndent)
	}

	secondExpectedIndent := "    "
	if secondIndent != secondExpectedIndent {
		t.Fatalf("Expected %q, got %q", secondExpectedIndent, secondIndent)
	}

	thirdExpectedIndent := "      "
	if thirdIndent != thirdExpectedIndent {
		t.Fatalf("Expected %q, got %q", thirdExpectedIndent, thirdIndent)
	}

	fourthExpectedIndent := "        "
	if fourthIndent != fourthExpectedIndent {
		t.Fatalf("Expected %q, got %q", fourthExpectedIndent, fourthIndent)
	}
}
