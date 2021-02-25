package responsehandler

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/pkg/encode"
	"github.com/csothen/onlinestand/pkg/models"
)

// Respond : Handles error from Service response and writes to HTTP Response Writer
func Respond(logger *log.Logger, rw http.ResponseWriter, res models.ServiceResponse, err error) {
	if err != nil {
		logger.Println(err)
		http.Error(rw, res.Err, res.Status)
	}

	rw.WriteHeader(res.Status)

	if err := encode.ToJSON(res.Value, rw); err != nil {
		logger.Println(err)
		http.Error(rw, "Unable to respond", http.StatusInternalServerError)
	}
}
