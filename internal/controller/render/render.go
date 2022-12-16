package render

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

func NewReponse(status int, w http.ResponseWriter, data interface{}) {
	var (
		response []byte = make([]byte, 0)
		err      error
	)

	if serr, ok := data.(error); ok {
		response, err = json.Marshal(map[string]string{
			"error": serr.Error(),
		})
		if err != nil {
			logrus.Error(err)
		}
	} else if data == nil {
		response, err = json.Marshal(map[string]string{
			"error": "unknown error",
		})
		if err != nil {
			logrus.Error(err)
		}
	} else {
		response, err = json.Marshal(data)
		if err != nil {
			logrus.Error(err)
		}
	}

	w.WriteHeader(status)
	w.Write(response)
}
