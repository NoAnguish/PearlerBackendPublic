package account

import (
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/doug-martin/goqu/v9"
)

var TableName string = "UserAccounts"

func GetById(s *database.Session, id string) (*Account, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"id": id}).ToSQL()
	data, err := database.Get[Account](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("account does not found")
	}

	return &data[0], nil
}

func Insert(s *database.Session, account Account) *api_errors.Error {
	query, _, _ := goqu.Insert(TableName).Rows(account).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func Update(s *database.Session, account Account) *api_errors.Error {
	query, _, _ := goqu.Update(TableName).Set(account).Where(goqu.Ex{"id": account.Id}).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func GetByFirebaseUId(s *database.Session, firebaseUId string) (*Account, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"firebase_uid": firebaseUId}).ToSQL()
	data, err := database.Get[Account](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("account does not found")
	}

	return &data[0], nil
}

func ExistsByFirebaseUId(s *database.Session, firebaseUId string) (bool, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"firebase_uid": firebaseUId}).ToSQL()
	data, err := database.Get[Account](query, s)

	if err != nil {
		return false, api_errors.NewInternalDatabaseError(err)
	}
	return (len(data) > 0), nil
}
