package s3

import (
	"bytes"

	"github.com/NoAnguish/PearlerBackend/backend/utils/config"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

func UploadImage(data []byte, extension string) (string, error) {
	s3Config, err := config.S3Config()
	if err != nil {
		return "", err
	}

	session, err := session.NewSession(&aws.Config{
		Region:   &s3Config.Region,
		Endpoint: &s3Config.Endpoint,
		Credentials: credentials.NewStaticCredentials(
			s3Config.Credentials.AccessKeyId,
			s3Config.Credentials.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		return "", err
	}

	log.Info().Msg("uploading image to s3")
	imageKey, err := upload(session, data, extension)
	if err != nil {
		return "", err
	}

	imageURL := s3Config.Endpoint + "/" + s3Config.Bucket + "/" + imageKey
	log.Info().Str("ImageURL", imageURL).Msg("image has been successfully uploaded to s3")
	return imageURL, nil
}

func upload(session *session.Session, image []byte, extension string) (string, error) {
	s3Config, err := config.S3Config()
	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(session)
	imageKey := "images/" + formatters.GenerateImageId() + "." + extension

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &s3Config.Bucket,
		Key:    aws.String(imageKey),
		Body:   bytes.NewReader(image),
	})
	if err != nil {
		return "", err
	}

	return imageKey, nil
}
