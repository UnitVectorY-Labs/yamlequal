package yamlequal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEqual(t *testing.T) {
	testDirs, err := os.ReadDir("tests")
	if err != nil {
		t.Fatalf("Error reading tests directory: %v", err)
	}

	for _, dir := range testDirs {
		if !dir.IsDir() {
			continue
		}

		dirPath := filepath.Join("tests", dir.Name())

		t.Run(dir.Name(), func(t *testing.T) {
			file1Path := filepath.Join(dirPath, "file1.yaml")
			file2Path := filepath.Join(dirPath, "file2.yaml")

			// If files don't exist test fails
			if _, err := os.Stat(file1Path); os.IsNotExist(err) {
				t.Fatalf("File %s does not exist", file1Path)
			}
			if _, err := os.Stat(file2Path); os.IsNotExist(err) {
				t.Fatalf("File %s does not exist", file2Path)
			}

			equal, _, err := CompareYAMLFiles(file1Path, file2Path)
			if err != nil {
				t.Fatalf("Error comparing YAML files in %s: %v", dirPath, err)
			}
			if !equal {
				t.Fatalf("Files %s and %s are not equal", file1Path, file2Path)
			}
		})
	}
}
