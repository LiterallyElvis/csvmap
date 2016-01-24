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

func (m *MapWriter) createOutputHeaderMap() {
	result := map[string]int{}
	for index, header := range m.InputHeaders {
		result[header] = index
	}
	m.OutputHeaderMap = result
}

func (m *MapWriter) createOutputSlice(row map[string]string) ([]string, error) {
	outputSlice := make([]string, len(m.OutputHeaderMap))
	for header, value := range row {
		if _, ok := m.OutputHeaderMap[header]; ok {
			outputSlice[m.OutputHeaderMap[header]] = value
		} else {
			return []string{}, fmt.Errorf("Provided row contains invalid header field: %v", header)
		}
	}
	return outputSlice, nil
}

// Write recreates the built-in CSV writer's Write method
func (m *MapWriter) Write(row map[string]string) error {
	if len(row) > len(m.OutputHeaderMap) {
		return fmt.Errorf("Provided row has %v fields, whereas there are only %v headers", len(row), len(m.OutputHeaderMap))
	}
	outputSlice, err := m.createOutputSlice(row)
	if err != nil {
		return err
	}
	err = m.Writer.Write(outputSlice)
	return err
}

// WriteAll rereates the built-in CSV writer's WriteAll method
func (m *MapWriter) WriteAll(rows []map[string]string) error {
	for _, row := range rows {
		err := m.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}

// Flush simply calls the built-in CSV writer's flush method
func (m *MapWriter) Flush() {
	m.Writer.Flush()
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
