package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func AddJSONExtension(fileName string) string {
	return fileName + ".json"
}

func CreateJSONFile(data interface{}, path string, fileName string) error {
	// Marshal the struct to JSON with indentation for readability
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %w", err)
	}

	// Create or open the specified JSON file
	fullPathFile := filepath.Join(path, AddJSONExtension(fileName))
	file, err := os.Create(fullPathFile)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	if _, err = file.Write(jsonData); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func ReadJSONFile(path string, fileName string, out interface{}) error {
	fullPathFile := filepath.Join(path, AddJSONExtension(fileName))
	data, err := os.ReadFile(fullPathFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	if err := json.Unmarshal(data, &out); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return nil
}

func RemoveJSONFile(fileName string, path string) bool {
	removed := true

	fullPathFile := filepath.Join(path, AddJSONExtension(fileName))
	exists := CheckJSONFile(fileName, path)

	if exists {
		err := os.Remove(fullPathFile)
		removed = err != nil
	}

	return removed
}

func CheckJSONFile(path string, fileName string) bool {
	fullPathFile := filepath.Join(path, AddJSONExtension(fileName))
	exists := true

	if _, err := os.Stat(fullPathFile); err != nil {
		exists = false
	}

	return exists
}
