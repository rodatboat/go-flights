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

func GetFlights(search FlightSearch) string {
	// Add logic if trip is RoundTrip, need to also search for 1) Flip destinations and add return date & 2) Append flights array.
	// tfs := Build(search.Flights, search.Passengers, search.Options.Class, search.Options.Trip)
	urlString := SerializeURL(context.Background(), search)
	return urlString
}

func SerializeURL(ctx context.Context, search FlightSearch) string {
	tfsData := Build(search.Flights, search.Passengers, search.Options.Class, search.Options.Trip)

	return "https://www.google.com/travel/flights/search" +
		"?tfs=" + tfsData.ToBase64() +
		"&curr=" + search.Options.Currency.String() +
		"&hl=" + search.Options.Language.String()
}
