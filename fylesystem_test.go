package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestEnv() string {
	testDir := "test_dir"
	os.Mkdir(testDir, 0755)
	return testDir
}

func teardownTestEnv(testDir string) {
	os.RemoveAll(testDir)
}

func TestOpen(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	if _, err := os.Stat(filepath.Join(testDir, "test.txt")); os.IsNotExist(err) {
		t.Fatalf("file was not created")
	}
}

func TestWrite(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	content := "Hello, World!"
	if err := Write(file, content); err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	readContent, err := Read(filepath.Join(testDir, "test.txt"))
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	if readContent != content {
		t.Fatalf("expected %s, got %s", content, readContent)
	}
}

func TestRead(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	filePath := filepath.Join(testDir, "test.txt")
	err := os.WriteFile(filePath, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("failed to write initial content: %v", err)
	}

	content, err := Read(filePath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	expectedContent := "Hello, World!"
	if content != expectedContent {
		t.Fatalf("expected %s, got %s", expectedContent, content)
	}
}

func TestIsEmpty(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	isEmpty, err := IsEmpty(file)
	if err != nil {
		t.Fatalf("failed to check if file is empty: %v", err)
	}

	if !isEmpty {
		t.Fatalf("expected file to be empty")
	}

	if err := Write(file, "not empty"); err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	isEmpty, err = IsEmpty(file)
	if err != nil {
		t.Fatalf("failed to check if file is empty: %v", err)
	}

	if isEmpty {
		t.Fatalf("expected file to not be empty")
	}
}

func TestClose(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	if err := Close(file); err != nil {
		t.Fatalf("failed to close file: %v", err)
	}

	// Check if the file is closed by attempting to read its information
	_, err = file.Stat()
	if err == nil {
		t.Fatalf("expected error when accessing closed file, got nil")
	}
}

func TestExists(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	filePath := filepath.Join(testDir, "test.txt")

	if Exists(filePath) {
		t.Fatalf("expected file to not exist")
	}

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	if !Exists(filePath) {
		t.Fatalf("expected file to exist")
	}
}

func TestRemove(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	filePath := filepath.Join(testDir, "test.txt")

	if err := Remove(filePath); err != nil {
		t.Fatalf("failed to remove file: %v", err)
	}

	if Exists(filePath) {
		t.Fatalf("expected file to be removed")
	}
}

func TestAddDirective(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	directive := FileDirective{
		Content: "Hello",
		Indent:  4,
		NewLine: true,
	}

	if err := AddDirective(file, directive, true); err != nil {
		t.Fatalf("failed to add directive: %v", err)
	}

	readContent, err := Read(filepath.Join(testDir, "test.txt"))
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	expectedContent := "\n    Hello"
	if readContent != expectedContent {
		t.Fatalf("expected %s, got %s", expectedContent, readContent)
	}
}

func TestAddDirectives(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	file, err := Open("test.txt", testDir)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	directives := []FileDirective{
		{Content: "Line1", Indent: 2, NewLine: true},
		{Content: "Line2", Indent: 4, NewLine: false},
	}

	if err := AddDirectives(file, directives); err != nil {
		t.Fatalf("failed to add directives: %v", err)
	}

	readContent, err := Read(filepath.Join(testDir, "test.txt"))
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	expectedContent := "\n  Line1\n    Line2"
	if readContent != expectedContent {
		t.Fatalf("expected %s, got %s", expectedContent, readContent)
	}
}

func TestBuildFile(t *testing.T) {
	testDir := setupTestEnv()
	defer teardownTestEnv(testDir)

	fileInfo := FileInfo{
		Name: "test.txt",
		Path: testDir,
		Directives: []FileDirective{
			{Content: "Line1", Indent: 2, NewLine: true},
			{Content: "Line2", Indent: 4, NewLine: false},
		},
	}

	file, err := BuildFile(fileInfo, true)
	if err != nil {
		t.Fatalf("failed to build file: %v", err)
	}

	if file != nil {
		t.Fatalf("expected file to be closed and nil")
	}

	readContent, err := Read(filepath.Join(testDir, "test.txt"))
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	expectedContent := "\n  Line1\n    Line2"
	if readContent != expectedContent {
		t.Fatalf("expected %s, got %s", expectedContent, readContent)
	}
}
