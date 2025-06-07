package iata

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

func readCSV(filePath string) ([]string, error) {
	csvFile := fs.Open(filePath)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.ReadAll()
	return reader.ReadAll() // returns

}

func main() {

	out := make(chan string)
	var wg sync.WaitGroup

	// Read CSV file into map.
	for _, k := range keys {
		a = airports[k]
		if a.Iata == "" || a.Iata == "0" {
			continue
		}
		if _, ok := checked[a.Iata]; ok {
			continue
		}
		checked[a.Iata] = struct{}{}

		wg.Add(1)
		go func(iata, tz, city string) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
			defer cancel()

			ok, err := session.IsIATASupported(ctx, iata)
			if err != nil {
				out <- result{err: err}
			}

			if ok {
				res := result{line: fmt.Sprintf(caseTmpl, iata, city, tz), err: nil}
				out <- res
			}
		}(a.Iata, a.Tz, a.City)
	}

	iataFileHeaderContent := fmt.Sprintf(
		`// This package contains IATA airport codes, which Google Flights API supports.
	//
	// Command: go run ./common/iata/_generate.go
	//
	// Generation date: %s
	//
	package iata

	type IATA string

	// IATA codes enum
	const (
		Unknown IATA = "UNKNOWN"
	`, time.Now().Format(time.DateOnly))

	go func() {
		wg.Wait()
		close(out)
	}()

	lines := []string{}

	for res := range out {
		if res.err != nil {
			log.Fatal(res.err)
		}
		lines = append(lines, res.line)
	}

	sort.Strings(lines)

	for _, line := range lines {
		iataFileContent += line
	}

	iataFileContent += `	}
	return Location{"Not supported IATA Code", "Not supported IATA Code"}
}
`

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

// Generate iata.go using CSV.
// Create both ENUM for IATA codes, and City/TimeZone mapping.
