[![GitHub release](https://img.shields.io/github/release/UnitVectorY-Labs/yamlequal.svg)](https://github.com/UnitVectorY-Labs/yamlequal/releases/latest) [![License](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/MIT) [![Active](https://img.shields.io/badge/Status-Active-green)](https://guide.unitvectorylabs.com/bestpractices/status/#active) [![codecov](https://codecov.io/gh/UnitVectorY-Labs/yamlequal/graph/badge.svg?token=JH6fDDrnah)](https://codecov.io/gh/UnitVectorY-Labs/yamlequal) [![Go Report Card](https://goreportcard.com/badge/github.com/UnitVectorY-Labs/yamlequal)](https://goreportcard.com/report/github.com/UnitVectorY-Labs/yamlequal)

# yamlequal

A lightweight Go library that verifies the semantic equality of YAML files.

## Features

- Compare two YAML files for semantic equality, regardless of formatting or key order
- Compare two YAML content strings directly

## Usage

```go
import "github.com/UnitVectorY-Labs/yamlequal"
```

### Comparing YAML Files

```go
func main() {
	// Compare two YAML files
	equal, diff, err := yamlequal.CompareFiles("file1.yaml", "file2.yaml")
	if err != nil {
		// Handle error
		fmt.Println("Error comparing files:", err)
		return
	}

	if equal {
		// Files are semantically equivalent
		fmt.Println("Same:", diff)
	} else {
		// Files differ semantically
		fmt.Println("Different: ", diff)
	}
}
```

### Comparing YAML Content Directly

```go
func main() {

	yamlContent1 := []byte(`
foo: bar
value: 42
`)

	yamlContent2 := []byte(`
foo: bar
`)
	// Compare two YAML content strings directly
	equal, diff, err := yamlequal.CompareYAML(yamlContent1, yamlContent2)
	if err != nil {
		// Handle error
		fmt.Println("Error comparing files:", err)
		return
	}

	if equal {
		// Files are semantically equivalent
		fmt.Println("Same:", diff)
	} else {
		// Files differ semantically
		fmt.Println("Different: ", diff)
	}
}
```
