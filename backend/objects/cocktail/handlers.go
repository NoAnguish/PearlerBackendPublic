package cocktail

import (
	"github.com/NoAnguish/PearlerBackend/backend/objects/image"
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/s3"
)

func GetAllHandler() (*CocktailsTruncatedResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultReadSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	cocktails, err := GetAll(s)
	if err != nil {
		return nil, err
	}
	response := CocktailsTruncatedResponse{*cocktails}

	return &response, nil
}

func GetByIdHandler(request CocktailIdRequest) (*CocktailResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultReadSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	cocktail, err := GetById(s, request.Id)
	if err != nil {
		return nil, err
	}

	response := CocktailResponse(*cocktail)
	return &response, nil
}

func CreateHandler(
	request CreateCocktailRequest,
	image *image.ImageRequest,
) (*CocktailIdResponse, *api_errors.Error) {
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

	cocktail := Cocktail{
		Id:          formatters.GenerateId(),
		Name:        request.Name,
		Recipe:      request.Recipe,
		ImageURL:    imageURL,
		Description: request.Description,
	}

	err := Insert(s, cocktail)
	if err != nil {
		return nil, err
	}
	return &CocktailIdResponse{Id: cocktail.Id}, nil
}

func UpdateHandler(request UpdateCocktailRequest) (*CocktailIdResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	cocktail, err := GetById(s, request.Id)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		cocktail.Name = request.Name
	}
	if request.Recipe != "" {
		cocktail.Recipe = request.Recipe
	}
	if request.Description != "" {
		cocktail.Description = request.Description
	}

	err = Update(s, *cocktail)
	if err != nil {
		return nil, err
	}
	return &CocktailIdResponse{Id: cocktail.Id}, nil
}
