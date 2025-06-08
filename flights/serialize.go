package flights

import "github.com/rodatboat/google-flights/internal"

func serializeFlights(flights []Flight) []*internal.Flight {
	serializedFlights := make([]*internal.Flight, len(flights))
	for idx, _ := range flights {
		currFlight := flights[idx]
		serializedFlights = append(serializedFlights,
			&internal.Flight{
				Date:     currFlight.Date,
				MaxStops: &currFlight.MaxStops,

				Airlines: currFlight.Airlines,
				FromAirport: &internal.Airport{
					Name: currFlight.FromAirport.name,
				},
				ToAirport: &internal.Airport{
					Name: currFlight.ToAirport.name,
				},
			})
	}
	return serializedFlights
}

func serializePassgengers(passengers Passengers) []internal.Passenger {
	totalLength := passengers.Adults + passengers.Children + passengers.InfantInSeat + passengers.InfantOnLap
	serializedPassengers := make([]internal.Passenger, totalLength)

	for range passengers.Adults {
		serializedPassengers = append(serializedPassengers, internal.Passenger_ADULT)
	}
	for range passengers.Children {
		serializedPassengers = append(serializedPassengers, internal.Passenger_CHILD)
	}
	for range passengers.InfantInSeat {
		serializedPassengers = append(serializedPassengers, internal.Passenger_INFANT_IN_SEAT)
	}
	for range passengers.InfantOnLap {
		serializedPassengers = append(serializedPassengers, internal.Passenger_INFANT_ON_LAP)
	}
	return serializedPassengers
}

func serializeTrip(trip Trip) internal.Trip {
	switch trip {
	case OneWay:
		return internal.Trip_ONE_WAY
	case RoundTrip:
		return internal.Trip_ROUND_TRIP
	default:
		return internal.Trip_UNKNOWN_TRIP
	}
}

func serializeClass(class Class) internal.Class {
	switch class {
	case Economy:
		return internal.Class_ECONOMY
	case PremiumEconomy:
		return internal.Class_PREMIUM_ECONOMY
	case Business:
		return internal.Class_BUSINESS
	case First:
		return internal.Class_FIRST
	default:
		return internal.Class_UNKNOWN_CLASS
	}
}
