package main

import (
	"fmt"

	"github.com/rodatboat/google-flights/common/iata"
	"github.com/rodatboat/google-flights/flights"
)

func main() {
	searchOptions := flights.DefaultSearchOptions()
	search := flights.FlightSearch{
		Flights: []flights.Flight{
			{
				Date:        "2025-08-01",
				MaxStops:    0,
				Airlines:    []string{"AA"},
				FromAirport: iata.DFW,
				ToAirport:   iata.MIA,
			},
		},
		Passengers: flights.Passengers{
			Adults: 2,
		},
		Options: searchOptions,
	}

	URL := flights.GetFlights(search)
	fmt.Println(URL)
}
