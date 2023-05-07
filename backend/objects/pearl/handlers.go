package pearl

import (
	"math"

	"github.com/NoAnguish/PearlerBackend/backend/objects/image"
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/s3"
)

func CreatePearlHandler(
	request CreatePearlRequest,
	image *image.ImageRequest,
) (*PearlIdResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	imageURL := ""
	if image != nil {
		imageURL, err_ = s3.UploadImage(image.Data, image.Extension)
		if err_ != nil {
			return nil, api_errors.NewS3UploadError(err_)
		}
	}

	pearl := Pearl{
		Id:         formatters.GenerateId(),
		AccountId:  request.AccountId,
		CocktailId: request.CocktailId,
		Grade:      request.Grade,
		Review:     request.Review,
		CreatedAt:  formatters.GetTimestap(),
		ImageURL:   imageURL,
	}

	err := Insert(s, pearl)
	if err != nil {
		return nil, err
	}
	return &PearlIdResponse{Id: pearl.Id}, nil
}

func GetStatsByCocktailIdHandler(request ObjectIdRequest) (*PearlsStatsResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	stats, err := GetStatsByCocktailId(s, request.Id)
	if err != nil {
		return nil, err
	}

	var rating float64

	if stats.PearlsAmount == 0 {
		rating = 0
	} else {
		rating = math.Round(float64(stats.PearlsGradesSum)/float64(stats.PearlsAmount)*100) / 100
	}

	response := PearlsStatsResponse{
		PearlsAmount:  stats.PearlsAmount,
		AverageRating: rating,
	}

	return &response, nil
}

func GetStatsByAccountIdHandler(request ObjectIdRequest) (*PearlsStatsResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultScanSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	stats, err := GetStatsByAccountId(s, request.Id)
	if err != nil {
		return nil, err
	}

	var rating float64

	if stats.PearlsAmount == 0 {
		rating = 0
	} else {
		rating = math.Round(float64(stats.PearlsGradesSum)/float64(stats.PearlsAmount)*100) / 100
	}

	response := PearlsStatsResponse{
		PearlsAmount:  stats.PearlsAmount,
		AverageRating: rating,
	}

	return &response, nil
}
