package csvmap

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVMap struct {
	Headers        []string
	HeaderIndexMap map[string]int
	FileContents   []map[string]string
}

func closeIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CSVMap) CreateHeaderIndexMap(headers []string) map[string]int {
	headerIndexMap := make(map[string]int, len(headers))

	for index, header := range headers {
		headerIndexMap[header] = index
	}

	return headerIndexMap
}

func (m *CSVMap) CreateRowMap(csvRow []string, headerMap map[string]int) map[string]string {
	result := map[string]string{}
	for header, index := range headerMap {
		result[header] = csvRow[index]
	}

	return result
}

func (m *CSVMap) CreateAllMaps(fileContents [][]string, headerMap map[string]int) []map[string]string {
	result := []map[string]string{}

	for _, row := range fileContents {
		newRow := map[string]string{}
		for header, index := range headerMap {
			newRow[header] = row[index]
		}
		result = append(result, newRow)
	}

	return result
}

func New(filePath string) *CSVMap {
	// Open our file.
	inputCSV, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputCSV.Close()
	reader := csv.NewReader(inputCSV)

	// Create our resulting struct
	output := &CSVMap{}

	inputHeaders, err := reader.Read()
	closeIfError(err)

	// Use our methods (defined above) to populate our struct fields
	output.Headers = inputHeaders
	output.HeaderIndexMap = output.CreateHeaderIndexMap(inputHeaders)

	remainderOfFile, err := reader.ReadAll()
	closeIfError(err)

	output.CreateAllMaps(remainderOfFile, output.HeaderIndexMap)

	// We're done!
	return output
}
