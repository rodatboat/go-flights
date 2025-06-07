package flights

import "github.com/rodatboat/google-flights/internal"

type TFS struct {
	Flights    []internal.Flight
	Passengers []internal.Passenger
	Class      internal.Class
	Trip       internal.Trip
}

func Init(
	flights []Flight,
	passengers Passengers,
	class Class,
	trip Trip,
) *TFS {
	return &TFS{
		Flights:    serializeFlights(flights),
		Passengers: serializePassgengers(passengers),
		Class:      serializeClass(class),
		Trip:       serializeTrip(trip),
	}
}

func (tfs *TFS) ToBase64() {
	// Serialize TFS
}

func (tfs *TFS) ToSerializedString() {
	// Serialize TFS
}
