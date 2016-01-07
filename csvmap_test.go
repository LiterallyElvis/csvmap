package cartogopher

import (
	"reflect"
	"testing"
)

// Test functions using file path approach.
func TestHeaderSliceCreation(t *testing.T) {
	csvMap, err := New("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}
	result := csvMap.Headers

	t.Logf("Generated Result: \n%v", result)

	expectedResult := []string{"first", "second", "third"}

	// Test that the value was actually created.
	if result == nil {
		t.Errorf("Test CSV headers returned nil\n", expectedResult)
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
	csvMap, err := New("test_csvs/test.csv")
	if err != nil {
		t.Errorf("Error returned: %v", err)
	}

	result := csvMap.HeaderIndexMap
	expectedResult := map[string]int{
		"first":  0,
		"second": 1,
		"third":  2,
	}

	// Test that the value was actually created.
	if result == nil {
		t.Errorf("Test CSV header map returned nil\n", expectedResult)
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

// Benchmarks for the file path approach.
func BenchmarkSmallFileHandlingViaFilePath(b *testing.B) {
	New("test_csvs/test.csv")
}

func BenchmarkLargeFileHandlingViaFilePath(b *testing.B) {
	New("test_csvs/FakeNameGeneratorFile.csv")
}
