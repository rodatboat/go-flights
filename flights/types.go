package flights

import "github.com/rodatboat/go-flights/common/iata"

type Flight struct {
	Date     string
	MaxStops int32
	Airlines []string

	FromAirport iata.IATA
	ToAirport   iata.IATA
}

type Passengers struct {
	Adults       int
	Children     int
	InfantInSeat int
	InfantOnLap  int
}

type Class int64

const (
	Economy Class = iota + 1
	PremiumEconomy
	Business
	First
)

type Trip int64

const (
	RoundTrip Trip = iota + 1
	OneWay
)
