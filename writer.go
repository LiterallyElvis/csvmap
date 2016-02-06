package cartogopher

import (
	"encoding/csv"
	"fmt"
	"io"
)

// MapWriter mimics the writer struct.
type MapWriter struct {
	InputHeaders    []string
	OutputHeaderMap map[string]int
	Writer          *csv.Writer
}

func (w *MapWriter) createOutputHeaderMap() {
	result := map[string]int{}
	for index, header := range w.InputHeaders {
		result[header] = index
	}
	w.OutputHeaderMap = result
}

func (w MapWriter) createOutputSlice(row map[string]string) ([]string, error) {
	outputSlice := make([]string, len(w.OutputHeaderMap))
	for header, value := range row {
		if _, ok := w.OutputHeaderMap[header]; ok {
			outputSlice[w.OutputHeaderMap[header]] = value
		} else {
			return []string{}, fmt.Errorf("Provided row contains invalid header field: %v", header)
		}
	}
	return outputSlice, nil
}

// Write recreates the built-in CSV writer's Write method
func (w MapWriter) Write(row map[string]string) error {
	if len(row) > len(w.OutputHeaderMap) {
		return fmt.Errorf("Provided row has %v fields, whereas there are only %v headers", len(row), len(w.OutputHeaderMap))
	}
	outputSlice, err := w.createOutputSlice(row)
	if err != nil {
		return err
	}
	err = w.Writer.Write(outputSlice)
	return err
}

// WriteAll rereates the built-in CSV writer's WriteAll method
func (w MapWriter) WriteAll(rows []map[string]string) error {
	for _, row := range rows {
		err := w.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}

// Flush simply calls the built-in CSV writer's flush method
func (w *MapWriter) Flush() {
	w.Writer.Flush()
}

// NewWriter creates a new writer
func NewWriter(w io.Writer, headers []string) *MapWriter {
	writer := csv.NewWriter(w)
	result := &MapWriter{
		InputHeaders: headers,
		Writer:       writer,
	}
	result.Writer.Write(result.InputHeaders)
	result.createOutputHeaderMap()
	return result
}
