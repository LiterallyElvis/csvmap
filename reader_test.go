package cartogopher

import (
	"os"
	"reflect"
	"testing"
)

// Test functions using file path approach.
func TestHeaderSliceCreation(t *testing.T) {
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned while opening file: %v", err)
	}
	reader, err := NewReader(inputCSV)
	if err != nil {
		t.Errorf("Error returned while creating new reader: %v", err)
	}

	result := reader.Headers

	t.Logf("Generated Result: \n%v", result)

	expectedResult := []string{"first", "second", "third"}

	// Test that the value was actually created.
	if result == nil {
		t.Error("Test CSV headers returned nil\n", expectedResult)
	} else {
		t.Log("Test CSV headers generated and are not nil")
	}

	// Test that there are as many generated headers as we expect.
	if len(expectedResult) > len(result) {
		t.Errorf("Resulting header slice length is %v, which is %v less than expected\n", len(result), len(expectedResult)-len(result))
	} else if len(expectedResult) < len(result) {
		t.Errorf("Resulting header slice length is %v, which is %v more than expected\n", len(result), len(result)-len(expectedResult))
	}

	// Check equality of the generated header slice and our expected result.
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Resulting header slice does not equal the expected result:\n\n%v\n\t!=\n%v\n", result, expectedResult)
	}
}

func TestHeaderMapCreation(t *testing.T) {
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned while opening file: %v", err)
	}
	reader, err := NewReader(inputCSV)
	if err != nil {
		t.Errorf("Error returned while creating new reader: %v", err)
	}

	result := reader.HeaderIndexMap
	expectedResult := map[string]int{
		"first":  0,
		"second": 1,
		"third":  2,
	}

	// Test that the value was actually created.
	if result == nil {
		t.Error("Test CSV header map returned nil\n", expectedResult)
	}

	// Test that every value matches up with the expected result.
	for key, value := range expectedResult {
		if _, ok := result[key]; !ok {
			t.Errorf("The following key is not located in the resulting header map: %v\n", key)
		} else if result[key] != value {
			t.Errorf("The generated header map has incorrect values for this key: %v\n", key)
		}
	}
}

func TestFileReading(t *testing.T) {
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned while opening file: %v", err)
	}
	reader, err := NewReader(inputCSV)
	if err != nil {
		t.Errorf("Error returned while creating new reader: %v", err)
	}

	expectedCSVMapRows := []map[string]string{
		{
			"first":  "one",
			"second": "two",
			"third":  "three",
		},
		{
			"first":  "a",
			"second": "b",
			"third":  "c",
		},
		{
			"first":  "Athos",
			"second": "Aramis",
			"third":  "Porthos",
		},
	}

	producedCSVMapRows, err := reader.ReadAll()
	if err != nil {
		t.Errorf("Error returned while reading all rows: %v", err)
	}

	for index, expectedRow := range expectedCSVMapRows {
		for header, value := range expectedRow {
			if producedCSVMapRows[index][header] != value {
				t.Errorf("Mismatch in expected value:\n\t'%v': %v\nvs. produced value: %v", header, value, producedCSVMapRows[index][header])
			}
		}
	}
}

// Benchmarks for the file path approach.
func BenchmarkSmallFileHandlingViaFilePath(b *testing.B) {
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		b.Errorf("Error returned while opening file: %v", err)
	}
	NewReader(inputCSV)
}
