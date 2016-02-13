package cartogopher

import (
	"os"
	"testing"
)

// Testing creation of the header to index map
func TestHeaderMapCreation(t *testing.T) {
	// open our test file
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned while opening file: %v", err)
	}

	// create our new Cartogopher reader
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

// Testing the entire contents fo a file
func TestFileReading(t *testing.T) {
	// open our test CSV
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned while opening file: %v", err)
	}

	// create our new Cartogopher reader
	reader, err := NewReader(inputCSV)
	if err != nil {
		t.Errorf("Error returned while creating new reader: %v", err)
	}

	// mapped contents of our demo CSV
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

	// read all the lines in the file, as you would do with encoding/csv
	producedCSVMapRows, err := reader.ReadAll()
	if err != nil {
		t.Errorf("Error returned while reading all rows: %v", err)
	}

	// assert euqality of results with expectations
	for index, expectedRow := range expectedCSVMapRows {
		for header, value := range expectedRow {
			if producedCSVMapRows[index][header] != value {
				t.Errorf("Mismatch in expected value:\n\t'%v': %v\nvs. produced value: %v", header, value, producedCSVMapRows[index][header])
			}
		}
	}
}

// Benchmarks for the file path approach.
func BenchmarkSmallFileHandling(b *testing.B) {
	inputCSV, err := os.Open("test_csvs/test.csv")
	if err != nil {
		b.Errorf("Error returned while opening file: %v", err)
	}
	NewReader(inputCSV)
}
