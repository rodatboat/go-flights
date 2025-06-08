package flights

import (
	"fmt"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
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
	URL          string
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

func GetFlights(search FlightSearch) FlightSearchResult {
	// TODO: Add logic if trip is RoundTrip, need to also search for 1) Flip destinations and add return date & 2) Append flights array.
	URL := SerializeFlightsURL(search)
	flightList, _ := scrapeFlights(URL)
	fmt.Println(flightList)
	return FlightSearchResult{
		URL: URL,
	}
}

func SerializeFlightsURL(search FlightSearch) string {
	tfsData := Build(search.Flights, search.Passengers, search.Options.Class, search.Options.Trip)
	return "https://www.google.com/travel/flights/search" +
		"?tfs=" + tfsData.ToBase64() +
		"&curr=" + search.Options.Currency.String() +
		"&hl=" + search.Options.Language.String()
}

func scrapeFlights(URL string) ([]FlightResult, error) {
	Ja3 := "123"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 OPR/117.0.0.0"
	options := cycletls.Options{
		Ja3: Ja3,
		Headers: map[string]string{
			"User-Agent": userAgent,
		},
	}

	client := cycletls.Init()
	resp, err := client.Do(URL, options, "GET")
	if resp.Status == 400 || err != nil {
		return nil, err
	}
	parseResponse(resp.Body)
	return nil, nil
}

func parseResponse(response string) []FlightResult {
	return nil
}
