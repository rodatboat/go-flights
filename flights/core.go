package flights

import (
	"context"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

type SearchOptions struct {
	Class    Class
	Trip     Trip
	Currency currency.Unit
	Language language.Tag
}

type FlightSearch struct {
	Flights    []Flight
	Passengers Passengers
	Options    SearchOptions
}

type FlightResult struct {
}

type FlightSearchResult struct {
	FlightsFound []FlightResult
	Error        error
}

func DefaultSearchOptions() SearchOptions {
	return SearchOptions{
		Class:    Economy,
		Trip:     OneWay,
		Currency: currency.USD,
		Language: language.English,
	}
}

func Init() {
	defaultOptions := DefaultSearchOptions()

	search := &FlightSearch{
		Options: defaultOptions,
	}
}

func GetFlights(search FlightSearch) *TFS {

	return nil
}

func SerializeURL(ctx context.Context) {
}
