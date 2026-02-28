package yamlequal

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestYAMLFiles(t *testing.T) {
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
			notEqualFlagPath := filepath.Join(dirPath, "notequal.flag")

			// Check if files should be not equal
			expectEqual := true
			if _, err := os.Stat(notEqualFlagPath); err == nil {
				expectEqual = false
			}

			// If files don't exist test fails
			if _, err := os.Stat(file1Path); os.IsNotExist(err) {
				t.Fatalf("File %s does not exist", file1Path)
			}
			if _, err := os.Stat(file2Path); os.IsNotExist(err) {
				t.Fatalf("File %s does not exist", file2Path)
			}

			equal, result, err := CompareFiles(file1Path, file2Path)
			if err != nil {
				t.Fatalf("Error comparing YAML files in %s: %v", dirPath, err)
			}

			// Print comparison result for clarity
			t.Logf("Test case: %s - Files %s - %s",
				dir.Name(),
				map[bool]string{true: "are equal", false: "are different"}[equal],
				result)

			if expectEqual && !equal {
				t.Fatalf("Expected files to be equal, but they were not. Result: %s", result)
			}

			if !expectEqual && equal {
				t.Fatalf("Expected files to be different, but they were equal. Result: %s", result)
			}
		})
	}
}

func TestCompareFilesFileNotFound(t *testing.T) {
	t.Run("first file missing", func(t *testing.T) {
		_, _, err := CompareFiles("nonexistent1.yaml", "nonexistent2.yaml")
		if err == nil {
			t.Fatal("Expected error for missing first file, got nil")
		}
		if !strings.Contains(err.Error(), "nonexistent1.yaml") {
			t.Fatalf("Expected error to reference first file, got: %v", err)
		}
	})

	t.Run("second file missing", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "yamlequal-*.yaml")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.WriteString("foo: bar\n")
		tmpFile.Close()

		_, _, err = CompareFiles(tmpFile.Name(), "nonexistent2.yaml")
		if err == nil {
			t.Fatal("Expected error for missing second file, got nil")
		}
		if !strings.Contains(err.Error(), "nonexistent2.yaml") {
			t.Fatalf("Expected error to reference second file, got: %v", err)
		}
	})
}

func TestCompareYAMLInvalid(t *testing.T) {
	t.Run("invalid first YAML", func(t *testing.T) {
		invalid := []byte(":\n  :\n    - :\n      invalid: [")
		valid := []byte("foo: bar\n")

		_, _, err := CompareYAML(invalid, valid)
		if err == nil {
			t.Fatal("Expected error for invalid first YAML, got nil")
		}
		if !strings.Contains(err.Error(), "first YAML") {
			t.Fatalf("Expected error to reference first YAML, got: %v", err)
		}
	})

	t.Run("invalid second YAML", func(t *testing.T) {
		valid := []byte("foo: bar\n")
		invalid := []byte(":\n  :\n    - :\n      invalid: [")

		_, _, err := CompareYAML(valid, invalid)
		if err == nil {
			t.Fatal("Expected error for invalid second YAML, got nil")
		}
		if !strings.Contains(err.Error(), "second YAML") {
			t.Fatalf("Expected error to reference second YAML, got: %v", err)
		}
	})
}

func TestCompareYAMLDirect(t *testing.T) {
	t.Run("empty content equal", func(t *testing.T) {
		equal, msg, err := CompareYAML([]byte(""), []byte(""))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !equal {
			t.Fatalf("Expected empty content to be equal, got: %s", msg)
		}
	})

	t.Run("multi-doc equal", func(t *testing.T) {
		yaml1 := []byte("foo: bar\n---\nbaz: qux\n")
		yaml2 := []byte("foo: bar\n---\nbaz: qux\n")

		equal, msg, err := CompareYAML(yaml1, yaml2)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !equal {
			t.Fatalf("Expected multi-doc YAML to be equal, got: %s", msg)
		}
	})

	t.Run("multi-doc different second doc", func(t *testing.T) {
		yaml1 := []byte("foo: bar\n---\nbaz: qux\n")
		yaml2 := []byte("foo: bar\n---\ndifferent: value\n")

		equal, _, err := CompareYAML(yaml1, yaml2)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if equal {
			t.Fatal("Expected multi-doc YAML with different second document to be NOT equal")
		}
	})

	t.Run("multi-doc different count", func(t *testing.T) {
		yaml1 := []byte("foo: bar\n---\nbaz: qux\n")
		yaml2 := []byte("foo: bar\n")

		equal, _, err := CompareYAML(yaml1, yaml2)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if equal {
			t.Fatal("Expected multi-doc YAML with different document count to be NOT equal")
		}
	})
}
