package cartogopher

import (
	"os"
	"strconv"
	"testing"
)

func TestWriterCreationAndOutput(t *testing.T) {
	testCSVFileName := "test_csvs/writer_test.csv"
	testHeaders := []string{"first", "second", "third"}
	testFile, err := os.Create(testCSVFileName)
	if err != nil {
		t.Errorf("Error returned while creating file: %v", err)
	}
	defer testFile.Close()
	writer := NewWriter(testFile, testHeaders)

	// first test that writing a row at all works.
	completeMap := map[string]string{
		"first":  "one",
		"second": "two",
		"third":  "three",
	}
	err = writer.Write(completeMap)
	if err != nil {
		t.Errorf("Error returned while writing row to file: %v", err)
	}

	// let's test making an empty map and writing it.
	emptyMap := make(map[string]string, len(completeMap))
	err = writer.Write(emptyMap)
	if err != nil {
		t.Errorf("Error returned while writing row to file: %v", err)
	}

	// then test that the row can be an 'incomplete' map
	incompleteMap := map[string]string{
		"first": "The Fellowship of the Ring",
		"third": "The Return of the King",
	}
	err = writer.Write(incompleteMap)
	if err != nil {
		t.Errorf("Error returned while writing row to file: %v", err)
	}

	// providing a map that is too large, however, should throw an error
	tooBigAMap := map[string]string{
		"first":  "one",
		"second": "two",
		"third":  "three",
	}
	for i := 0; i < 42; i++ {
		tooBigAMap[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	err = writer.Write(tooBigAMap)
	if err == nil {
		t.Errorf("No error returned while writing a row too big for the file")
	}

	// providing a map with an invalid field should throw an error
	badMap := map[string]string{
		"none":   "Athos",
		"should": "Aramis",
		"match":  "Porthos",
	}
	err = writer.Write(badMap)
	if err == nil {
		t.Errorf("No error returned while writing a bad row to the file")
	}
	// write the rows already
	writer.Flush()

	// test the content of the created CSV.
	// TODO

	// clean up created files.
	os.Remove(testCSVFileName)
}

func BenchmarkWriterShort(b *testing.B) {
	testCSVFileName := "test_csvs/writer_short_benchmark.csv"
	testHeaders := []string{"first", "second", "third"}
	testFile, err := os.Create(testCSVFileName)
	if err != nil {
		b.Errorf("Error returned while creating file: %v", err)
	}
	defer testFile.Close()
	writer := NewWriter(testFile, testHeaders)

	simpleRow := map[string]string{
		"first":  "one",
		"second": "two",
		"third":  "three",
	}
	err = writer.Write(simpleRow)
	if err != nil {
		b.Errorf("Error returned while writing row to file: %v", err)
	}
	writer.Flush()
	os.Remove(testCSVFileName)
}

func BenchmarkWriterLonger(b *testing.B) {
	testCSVFileName := "test_csvs/writer_longer_benchmark.csv"
	arbitrarilyLargeNumber := 100
	testHeaders := make([]string, arbitrarilyLargeNumber)
	for i := 0; i < arbitrarilyLargeNumber; i++ {
		testHeaders[i] = strconv.Itoa(i)
	}
	testFile, err := os.Create(testCSVFileName)
	if err != nil {
		b.Errorf("Error returned while creating file: %v", err)
	}
	defer testFile.Close()
	writer := NewWriter(testFile, testHeaders)

	for rowNumber := 0; rowNumber < arbitrarilyLargeNumber*10; rowNumber++ {
		newRow := map[string]string{}
		for i := 0; i < arbitrarilyLargeNumber; i++ {
			newRow[strconv.Itoa(i)] = strconv.Itoa(i)
		}
		err = writer.Write(newRow)
		if err != nil {
			b.Errorf("Error returned while writing row to file: %v", err)
		}
	}

	writer.Flush()
	os.Remove(testCSVFileName)
}
