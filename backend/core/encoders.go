package core

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func jsonEncoder[Response any](w http.ResponseWriter, r *Response) error {
	err := json.NewEncoder(w).Encode(*r)
	if err != nil {
		return err
	}

	log.Info().Interface("Response", *r).Msg("Successfully returned JSON response")
	return nil
}

func errorEncoder(w http.ResponseWriter, apiErr error) error {
	resp := make(map[string]string)
	resp["message"] = apiErr.Error()

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}
	return nil
}
