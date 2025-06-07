package flights

type Airport struct {
	name string
}

type Flight struct {
	Date     string
	MaxStops int64
	Airlines []string

	FromAirport Airport
	ToAirport   Airport
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
