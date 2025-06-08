package main

import (
	"fmt"

	"github.com/rodatboat/go-flights/common/iata"
	"github.com/rodatboat/go-flights/flights"
)

func main() {
	searchOptions := flights.DefaultSearchOptions()
	search := flights.FlightSearch{
		Flights: []flights.Flight{
			{
				Date:        "2025-08-01",
				MaxStops:    3,
				FromAirport: iata.DFW,
				ToAirport:   iata.MIA,
			},
		},
		Passengers: flights.Passengers{
			Adults: 2,
		},
		Options: searchOptions,
	}

	result := flights.GetFlights(search)
	fmt.Println(result.URL)
}
