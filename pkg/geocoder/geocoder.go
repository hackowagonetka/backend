package geocoder

type GeocoderInput struct {
	Lon float64
	Lat float64
}

type GeocoderOutput struct {
	Name string
}

type Geocoder interface {
	Request(i GeocoderInput) (*GeocoderOutput, error)
}
