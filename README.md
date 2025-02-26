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
	equal, diff, err := yamlequal.CompareFiles("file11.yaml", "file2.yaml")
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
	// Compare two YAML files
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
