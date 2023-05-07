package core

import (
	"net/http"

	"github.com/NoAnguish/PearlerBackend/backend/objects/image"
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/rs/zerolog/log"
)

func handlerWrapper[Request any, Response any](
	decoder func(*http.Request, *Request) error,
	handler func(Request) (*Response, *api_errors.Error),
	encoder func(http.ResponseWriter, *Response) error,
	d *Daemon) func(http.ResponseWriter, *http.Request) {

	wrap := func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("Url", r.URL.Path).Msg("HTTP request processing")

		var request Request
		err := decoder(r, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorEncoder(w, err)
			return
		}

		response, apiErr := handler(request)

		if apiErr != nil {
			log.Error().Err(apiErr).Msg("Error occured while processing request")
			w.WriteHeader(apiErr.StatusCode())
			errorEncoder(w, apiErr)
			return
		}
		log.Info().Msg("Request was processed")

		err = encoder(w, response)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while encoding response")
			w.WriteHeader(http.StatusInternalServerError)
			errorEncoder(w, err)
			return
		}
	}

	return wrap
}

func imageHandlerWrapper[Request any, Response any](
	decoder func(*http.Request, *Request) (*image.ImageRequest, error),
	handler func(Request, *image.ImageRequest) (*Response, *api_errors.Error),
	encoder func(http.ResponseWriter, *Response) error,
	d *Daemon) func(http.ResponseWriter, *http.Request) {

	wrap := func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("Url", r.URL.Path).Msg("HTTP request processing")

		var request Request
		image, err := decoder(r, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorEncoder(w, err)
			return
		}

		response, apiErr := handler(request, image)

		if apiErr != nil {
			log.Error().Err(apiErr).Msg("Error occured while processing request")
			w.WriteHeader(apiErr.StatusCode())
			errorEncoder(w, apiErr)
			return
		}
		log.Info().Msg("Request was processed")

		err = encoder(w, response)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while encoding response")
			w.WriteHeader(http.StatusInternalServerError)
			errorEncoder(w, err)
			return
		}
	}

	return wrap
}

func emptyRequestHandlerWrapper[Response any](
	handler func() (*Response, *api_errors.Error),
	encoder func(http.ResponseWriter, *Response) error,
	d *Daemon) func(http.ResponseWriter, *http.Request) {

	wrap := func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("Url", r.URL.Path).Msg("HTTP request processing")

		response, apiErr := handler()

		if apiErr != nil {
			log.Error().Err(apiErr).Msg("Error occured while processing request")
			w.WriteHeader(http.StatusBadRequest)
			errorEncoder(w, apiErr)
			return
		}
		log.Info().Msg("Request was processed")

		err := encoder(w, response)
		if err != nil {
			log.Error().Err(err).Msg("Error occured while encoding response")
			w.WriteHeader(http.StatusInternalServerError)
			errorEncoder(w, err)
			return
		}
	}

	return wrap
}

func JsonHandlerWrapper[Request any, Response any](
	handler func(Request) (*Response, *api_errors.Error), d *Daemon) func(http.ResponseWriter, *http.Request) {

	decoder := jsonDecoder[Request]
	encoder := jsonEncoder[Response]
	return handlerWrapper(decoder, handler, encoder, d)
}

func QueryParametersHandlerWrapper[Request any, Response any](
	handler func(Request) (*Response, *api_errors.Error), d *Daemon) func(http.ResponseWriter, *http.Request) {

	decoder := queryParameterDecoder[Request]
	encoder := jsonEncoder[Response]
	return handlerWrapper(decoder, handler, encoder, d)
}

func NoRequestHandlerWrapper[Response any](
	handler func() (*Response, *api_errors.Error), d *Daemon) func(http.ResponseWriter, *http.Request) {

	encoder := jsonEncoder[Response]
	return emptyRequestHandlerWrapper(handler, encoder, d)
}

func InsertImageHandlerWrapper[Request any, Response any](
	handler func(Request, *image.ImageRequest) (*Response, *api_errors.Error),
	d *Daemon,
) func(http.ResponseWriter, *http.Request) {

	decoder := imageDecoder[Request]
	encoder := jsonEncoder[Response]
	return imageHandlerWrapper(decoder, handler, encoder, d)
}
