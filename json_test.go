package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddJSONExtension(t *testing.T) {
	fileName := "test"
	expected := "test.json"
	result := AddJSONExtension(fileName)

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestCreateAndReadJSONFile(t *testing.T) {
	type TestData struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	data := TestData{Name: "example", Value: 42}

	tempDir := t.TempDir()
	fileName := "test_file"

	// Test CreateJSONFile
	err := CreateJSONFile(data, tempDir, fileName)
	if err != nil {
		t.Fatalf("CreateJSONFile failed: %v", err)
	}

	// Verify file existence
	fullPathFile := filepath.Join(tempDir, AddJSONExtension(fileName))
	if _, err := os.Stat(fullPathFile); err != nil {
		t.Fatalf("expected file %s to exist, but it does not", fullPathFile)
	}

	// Test ReadJSONFile
	var readData TestData
	err = ReadJSONFile(tempDir, fileName, &readData)
	if err != nil {
		t.Fatalf("ReadJSONFile failed: %v", err)
	}

	if readData != data {
		t.Errorf("expected %v, got %v", data, readData)
	}
}

func TestRemoveJSONFile(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "test_file"
	fullPathFile := filepath.Join(tempDir, AddJSONExtension(fileName))

	// Create a dummy file
	err := os.WriteFile(fullPathFile, []byte("{}"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Test RemoveJSONFile
	removed := RemoveJSONFile(fileName, tempDir)
	if !removed {
		t.Errorf("expected file to be removed, but it was not")
	}

	// Verify file does not exist
	if _, err := os.Stat(fullPathFile); err != nil || os.IsNotExist(err) {
		t.Errorf("expected file %s to be removed, but it still exists", fullPathFile)
	}
}

func TestCheckJSONFile(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "test_file"
	fullPathFile := filepath.Join(tempDir, AddJSONExtension(fileName))

	// Verify non-existent file
	exists := CheckJSONFile(tempDir, fileName)
	if exists {
		t.Errorf("expected file %s to not exist, but CheckJSONFile returned true", fullPathFile)
	}

	// Create a dummy file
	err := os.WriteFile(fullPathFile, []byte("{}"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Verify file existence
	exists = CheckJSONFile(tempDir, fileName)
	if !exists {
		t.Errorf("expected file %s to exist, but CheckJSONFile returned false", fullPathFile)
	}
}
