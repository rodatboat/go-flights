package flights

import (
	"encoding/base64"

	"github.com/rodatboat/google-flights/internal"
	"google.golang.org/protobuf/proto"
)

type TFS struct {
	Flights    []*internal.Flight
	Passengers []internal.Passenger
	Class      internal.Class
	Trip       internal.Trip
}

func Build(
	flights []Flight,
	passengers Passengers,
	class Class,
	trip Trip,
) *TFS {
	// Add logic if trip is RoundTrip, need to also search for 1) Flip destinations and add return date & 2) Append flights array.
	return &TFS{
		Flights:    serializeFlights(flights),
		Passengers: serializePassgengers(passengers),
		Class:      serializeClass(class),
		Trip:       serializeTrip(trip),
	}
}

func (tfs *TFS) ToBase64() (string, error) {
	tfsData, err := tfs.ToSerializedTFS()
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(tfsData), nil
}

func (tfs *TFS) ToSerializedTFS() ([]byte, error) {
	return proto.Marshal(&internal.FlightPayload{
		Data:       tfs.Flights,
		Passengers: tfs.Passengers,
		Class:      tfs.Class,
		Trip:       tfs.Trip,
	})
}
