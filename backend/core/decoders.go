package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/NoAnguish/PearlerBackend/backend/objects/image"
	"github.com/rs/zerolog/log"
)

func jsonDecoder[Request any](r *http.Request, v *Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing request body")
		return err
	}

	err = json.Unmarshal(reqBody, v)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing json")
		return err
	}

	log.Info().Interface("Request", *v).Msg("Successfylly got JSON request")

	return nil
}

func imageDecoder[Request any](r *http.Request, v *Request) (*image.ImageRequest, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing request body")
		return nil, err
	}

	// parsing json from "data" field
	err = json.Unmarshal([]byte(r.FormValue("data")), v)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing json")
		return nil, err
	}
	log.Info().Interface("Request", *v).Msg("Successfylly got JSON request")

	image, err := decodeImage(r)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func decodeImage(r *http.Request) (*image.ImageRequest, error) {
	// please note that multipart should be already parsed
	// trying to get an image file
	file, _, err := r.FormFile("image_data")
	if err == http.ErrMissingFile {
		log.Info().Msg("No image file was provided")
		return nil, nil
	}
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing image file")
		return nil, err
	}
	defer file.Close()

	// getting an image from file
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Error().Err(err).Msg("Error occured while copying image data")
		return nil, err
	}
	data := buf.Bytes()

	if len(data) == 0 {
		log.Info().Msg("Image body is empty")
		return nil, nil
	}

	extension := r.FormValue("file_extension")
	if extension == "" {
		err = errors.New("File extension expected but was not found")
		log.Error().Err(err).Msg("Error occured while extracting file extension")
		return nil, err
	}

	log.Info().Msg("Successfylly got a picture")
	return &image.ImageRequest{Extension: extension, Data: data}, nil
}

func queryParameterDecoder[Request any](r *http.Request, v *Request) error {
	queryValues := r.URL.Query()
	simpleValues := make(map[string]string)

	for key, value := range queryValues {
		if len(value) > 1 {
			log.Error().Str("Key", key).Interface("Query", queryValues).Msg("More than one value for some key")
			return errors.New("more than one value for some key")
		}
		simpleValues[key] = value[0]
	}

	jsonString, err := json.Marshal(simpleValues)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing query parameters")
		return err
	}

	err = json.Unmarshal(jsonString, v)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing query parameters")
		return err
	}

	log.Info().Interface("Request", *v).Msg("Successfylly got QueryParameters request")
	return nil
}
