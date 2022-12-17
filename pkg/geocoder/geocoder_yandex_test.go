package geocoder

import (
	"fmt"
	"testing"
)

func TestYandexGeocodere_Request(t *testing.T) {
	g := NewYandexGeocoder("220d4c84-d54d-4a96-af30-e00235c569e3")
	fmt.Println(g.Request(GeocoderInput{
		Lon: 81.460766,
		Lat: 50.999759,
	}))
}
