// datacleaner.go
// Package datacleaner provides functions to clean input data.
// This example focuses on trimming whitespace from strings,
// but additional sanitization/validation rules can be added.
package datacleaner
import (
	"encoding/json"
	"strings"
)
// Clean recursively cleans an input data structure.
// It trims whitespace from strings and recursively processes maps and slices.
func Clean(input interface{}) interface{} {
	switch v := input.(type) {
	case string:
		// Clean string: trim whitespace and add additional cleaning logic as needed.
		return strings.TrimSpace(v)
	case []interface{}:
		// Recursively clean each element in a slice.
		for i, elem := range v {
			v[i] = Clean(elem)
		}
		return v
	case map[string]interface{}:
		// Recursively clean all values in a map.
		for key, val := range v {
			v[key] = Clean(val)
		}
		return v
	default:
		// For all other types, return the value as is.
		return v
	}
}
// CleanJSON accepts a raw JSON byte slice, cleans it, and returns the cleaned JSON.
// It unmarshals the JSON to process it as a generic object, cleans it, then marshals it back.
func CleanJSON(rawJSON []byte) ([]byte, error) {
	// Unmarshal the JSON into a generic data structure.
	var data interface{}
	if err := json.Unmarshal(rawJSON, &data); err != nil {
		return nil, err
	}
	// Clean the data recursively.
	cleanedData := Clean(data)
	// Marshal the cleaned data back to JSON.
	return json.Marshal(cleanedData)
}
