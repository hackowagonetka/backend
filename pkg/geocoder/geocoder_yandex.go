package geocoder

import (
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

var _ Geocoder = (*YandexGeocoder)(nil)

type YandexGeocoder struct {
	token string
}

func NewYandexGeocoder(token string) *YandexGeocoder {
	return &YandexGeocoder{token: token}
}

type GeocoderYandexResponse struct {
	Response struct {
		GeoObjectCollection struct {
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Text string `json:"text"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

func (g *YandexGeocoder) Request(i GeocoderInput) (*GeocoderOutput, error) {
	url := "https://geocode-maps.yandex.ru/1.x/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", g.token)
	q.Add("geocode", fmt.Sprintf("%f,%f", i.Lon, i.Lat))
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geocoder: YandexGeocoder: %s", string(b))
	}

	var body GeocoderYandexResponse
	if err := json.Unmarshal(b, &body); err != nil {
		return nil, fmt.Errorf("geocoder: YandexGeocoder: %w", err)
	}

	if len(body.Response.GeoObjectCollection.FeatureMember) == 0 {
		return nil, fmt.Errorf("geocoder: YandexGeocoder: %s", string(b))
	}

	output := GeocoderOutput{
		Name: body.Response.GeoObjectCollection.FeatureMember[0].GeoObject.MetaDataProperty.GeocoderMetaData.Text,
	}

	return &output, nil
}
