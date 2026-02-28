package yamlequal

import (
	"bytes"
	"fmt"
	"io"
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
// Multi-document YAML streams are fully supported; all documents in each stream
// are decoded and compared.
func CompareYAML(yamlContent1, yamlContent2 []byte) (bool, string, error) {
	// Decode all documents from each YAML stream
	docs1, err := decodeAllDocuments(yamlContent1)
	if err != nil {
		return false, "", fmt.Errorf("error parsing first YAML content: %v", err)
	}
	docs2, err := decodeAllDocuments(yamlContent2)
	if err != nil {
		return false, "", fmt.Errorf("error parsing second YAML content: %v", err)
	}

	// Compare the two document slices using reflect.DeepEqual.
	// Note: This comparison is independent of map key order.
	if reflect.DeepEqual(docs1, docs2) {
		return true, equalMessage, nil
	}

	return false, notEqualMessage, nil
}

// decodeAllDocuments decodes all YAML documents from a byte stream.
func decodeAllDocuments(content []byte) ([]any, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	var docs []any
	for {
		var doc any
		err := decoder.Decode(&doc)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}
