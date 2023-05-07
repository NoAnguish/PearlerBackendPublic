package account

import (
	"github.com/NoAnguish/PearlerBackend/backend/objects/image"
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/NoAnguish/PearlerBackend/backend/utils/formatters"
	"github.com/NoAnguish/PearlerBackend/backend/utils/s3"
)

func CreateAccountHandler(request CreateAccountRequest) (*AccountIdResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	account := Account{
		Id:          formatters.GenerateId(),
		Email:       request.Email,
		FirebaseUId: request.FirebaseUId,
		Name:        request.Name,
	}

	exists, err := ExistsByFirebaseUId(s, request.FirebaseUId)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, api_errors.NewBadRequestError("user with such firebase user id exists")
	}

	err = Insert(s, account)
	if err != nil {
		return nil, err
	}

	return &AccountIdResponse{Id: account.Id}, nil
}

func GetIdByFirebaseUIdHandler(request FirebaseUIdRequest) (*GetByFirebaseUIdResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	exists, err := ExistsByFirebaseUId(s, request.FirebaseUId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return &GetByFirebaseUIdResponse{
			Metadata: GetByFirebaseUIdResponseMetadata{Exists: false},
		}, nil
	}

	account, err := GetByFirebaseUId(s, request.FirebaseUId)
	if err != nil {
		return nil, err
	}

	return &GetByFirebaseUIdResponse{
		Id:       account.Id,
		Metadata: GetByFirebaseUIdResponseMetadata{Exists: true},
	}, nil
}

func GetByIdHandler(request AccountIdRequest) (*AccountResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultReadSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	account, err := GetById(s, request.Id)
	if err != nil {
		return nil, err
	}
	response := AccountResponse(*account)

	return &response, nil
}

func UpdateAccountHandler(
	request UpdateAccountRequest,
	image *image.ImageRequest,
) (*AccountIdResponse, *api_errors.Error) {
	s, err_ := database.PrepareDefaultWriteSession()
	if err_ != nil {
		return nil, api_errors.NewDatabaseConnectionError(err_)
	}
	defer s.Close()

	account, err := GetById(s, request.Id)
	if err != nil {
		return nil, err
	}

	if image != nil {
		imageURL, err_ := s3.UploadImage(image.Data, image.Extension)
		if err_ != nil {
			return nil, api_errors.NewS3UploadError(err_)
		}
		account.ImageURL = imageURL
	}

	if request.Name != "" {
		account.Name = request.Name
	}
	account.Description = request.Description // description could be empty

	err = Update(s, *account)
	if err != nil {
		return nil, err
	}

	return &AccountIdResponse{account.Id}, nil
}
