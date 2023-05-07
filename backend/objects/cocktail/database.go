package cocktail

import (
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/doug-martin/goqu/v9"
)

var TableName string = "Cocktails"

func GetAll(s *database.Session) (*[]CocktailTruncated, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Select(&CocktailTruncated{}).ToSQL()
	data, err := database.Get[CocktailTruncated](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func Insert(s *database.Session, cocktail Cocktail) *api_errors.Error {
	query, _, _ := goqu.Insert(TableName).Rows(cocktail).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func Update(s *database.Session, cocktail Cocktail) *api_errors.Error {
	query, _, _ := goqu.Update(TableName).Set(cocktail).Where(goqu.Ex{"id": cocktail.Id}).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func GetById(s *database.Session, id string) (*Cocktail, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"id": id}).ToSQL()
	data, err := database.Get[Cocktail](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("cocktail by such id not found")
	}
	return &data[0], nil
}
