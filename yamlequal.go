package yamlequal

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// CompareYAMLFiles reads two YAML files, parses them, and compares their
// content for structural and data equality. It returns a boolean indicating
// whether the files are equal, a diff message (if they differ), and an error if any.
func CompareYAMLFiles(filePath1, filePath2 string) (bool, string, error) {
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
		return true, "", nil
	}

	// Unmarshal the YAML content into generic data structures
	var data1, data2 interface{}
	if err := yaml.Unmarshal(content1, &data1); err != nil {
		return false, "", fmt.Errorf("error parsing YAML from %q: %v", filePath1, err)
	}
	if err := yaml.Unmarshal(content2, &data2); err != nil {
		return false, "", fmt.Errorf("error parsing YAML from %q: %v", filePath2, err)
	}

	// Compare the two data structures using reflect.DeepEqual.
	// Note: This comparison is independent of map key order.
	if reflect.DeepEqual(data1, data2) {
		return true, "", nil
	}

	// TODO: Implement a more detailed diff message
	diff := "YAML files are NOT equal"
	return false, diff, nil
}
