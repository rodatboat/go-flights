package flights

import (
	"encoding/base64"

	"github.com/rodatboat/go-flights/internal"
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
	return &TFS{
		Flights:    serializeFlights(flights),
		Passengers: serializePassengers(passengers),
		Class:      serializeClass(class),
		Trip:       serializeTrip(trip),
	}
}

func (tfs *TFS) ToBase64() string {
	tfsData, err := tfs.ToSerializedTFS()
	if err != nil {
		panic("Failed to serialize TFS")
	}
	return base64.RawURLEncoding.EncodeToString(tfsData)
}

func (tfs *TFS) ToSerializedTFS() ([]byte, error) {
	return proto.Marshal(&internal.FlightPayload{
		Data:       tfs.Flights,
		Passengers: tfs.Passengers,
		Class:      tfs.Class,
		Trip:       tfs.Trip,
	})
}
