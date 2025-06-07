package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

type Airport struct {
	Code     string
	TZ       string
	Name     string
	CityCode string
}

func readCSV(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s.\n%w", filePath, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	return reader.ReadAll()
}

func main() {

	out := make(chan string)
	var wg sync.WaitGroup

	iataEntries, err := readCSV("./common/iata/generate/airports.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		panic("")
	}

	fmt.Println("Successfully read", len(iataEntries), "records.")

	// Read CSV file into map.
	for _, entry := range iataEntries[1:] {
		airport := &Airport{
			Code:     entry[0],
			TZ:       entry[1],
			Name:     entry[2],
			CityCode: entry[3],
		}

		wg.Add(1)
		go func(a *Airport) {
			defer wg.Done()

			row := fmt.Sprintf(`	%s IATA = "%s"`, a.Code, a.Code)
			res := fmt.Sprintln(row)
			out <- res
		}(airport)
	}

	iataFileContent := fmt.Sprintf(
		`// This package contains IATA airport codes, which Google Flights API supports.
//
// Command: go run ./common/iata/generate/generate.go
//
// Generation date: %s
//
package iata

type IATA string

// IATA codes enum
const (
	UNKNOWN IATA = "UNKNOWN"
`, time.Now().Format(time.DateOnly))

	go func() {
		wg.Wait()
		close(out)
	}()

	lines := []string{}
	for line := range out {
		lines = append(lines, line)
	}

	sort.Strings(lines)

	for _, line := range lines {
		iataFileContent += line
	}

	iataFileContent += `)`
	iataFile, err := os.Create("./common/iata/iata.go")
	if err != nil {
		log.Fatal(err)
	}
	defer iataFile.Close()

	_, err = iataFile.WriteString(iataFileContent)
	if err != nil {
		log.Fatal(err)
	}
}
