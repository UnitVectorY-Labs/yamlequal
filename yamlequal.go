package yamlequal

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

const (
	equalMessage    = "YAML content is equal"
	notEqualMessage = "YAML content is NOT equal"
)

// CompareFiles reads two YAML files, parses them, and compares their
// content for structural and data equality. It returns a boolean indicating
// whether the files are equal, a diff message (if they differ), and an error if any.
func CompareFiles(filePath1, filePath2 string) (bool, string, error) {
	// Read both files
	content1, err := os.ReadFile(filePath1)
	if err != nil {
		return false, "", fmt.Errorf("error reading file %q: %v", filePath1, err)
	}
	content2, err := os.ReadFile(filePath2)
	if err != nil {
		return false, "", fmt.Errorf("error reading file %q: %v", filePath2, err)
	}

	// Handle the case of two empty files (considered equal)
	if len(content1) == 0 && len(content2) == 0 {
		return true, equalMessage, nil
	}

	return CompareYAML(content1, content2)
}

// CompareYAML compares two YAML content strings for structural and data equality.
// It returns a boolean indicating whether the content is equal, a diff message
// (if they differ), and an error if any parsing errors occur.
func CompareYAML(yamlContent1, yamlContent2 []byte) (bool, string, error) {
	// Unmarshal the YAML content into generic data structures
	var data1, data2 any
	if err := yaml.Unmarshal(yamlContent1, &data1); err != nil {
		return false, "", fmt.Errorf("error parsing first YAML content: %v", err)
	}
	if err := yaml.Unmarshal(yamlContent2, &data2); err != nil {
		return false, "", fmt.Errorf("error parsing second YAML content: %v", err)
	}

	// Compare the two data structures using reflect.DeepEqual.
	// Note: This comparison is independent of map key order.
	if reflect.DeepEqual(data1, data2) {
		return true, equalMessage, nil
	}

	return false, notEqualMessage, nil
}
