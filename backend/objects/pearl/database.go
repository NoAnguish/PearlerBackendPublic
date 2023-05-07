package pearl

import (
	"fmt"

	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/doug-martin/goqu/v9"
)

var TableName string = "Pearls"

func GetById(s *database.Session, id string) (*Pearl, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"id": id}).ToSQL()
	data, err := database.Get[Pearl](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("pearl does not exist")
	}
	return &data[0], nil
}

func Insert(s *database.Session, pearl Pearl) *api_errors.Error {
	query, _, _ := goqu.Insert(TableName).Rows(pearl).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func Update(s *database.Session, pearl Pearl) *api_errors.Error {
	query, _, _ := goqu.Update(TableName).Set(pearl).Where(goqu.Ex{"id": pearl.Id}).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func GetByAccountId(s *database.Session, accountId string) (*[]Pearl, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"account_id": accountId}).Order(goqu.I("created_at").Desc()).ToSQL()
	data, err := database.Get[Pearl](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetByCocktailId(s *database.Session, cocktailId string) (*[]Pearl, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"cocktail_id": cocktailId}).Order(goqu.I("created_at").Desc()).ToSQL()
	data, err := database.Get[Pearl](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetStatsByCocktailId(s *database.Session, cocktailId string) (*PearlStats, *api_errors.Error) {
	query := `
		SELECT 
			COUNT(*) AS "pearls_amount",
			COALESCE(SUM("grade"), 0) AS "pearls_grades_sum"
		FROM "%s"
		WHERE "cocktail_id" = '%s';
	`
	query = fmt.Sprintf(query, TableName, cocktailId)
	data, err := database.Get[PearlStats](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("cocktail does not exist")
	}
	return &data[0], nil
}

func GetStatsByAccountId(s *database.Session, cocktailId string) (*PearlStats, *api_errors.Error) {
	query := `
		SELECT 
			COUNT(*) AS "pearls_amount",
			COALESCE(SUM("grade"), 0) AS "pearls_grades_sum"
		FROM "%s"
		WHERE "account_id" = '%s';
	`
	query = fmt.Sprintf(query, TableName, cocktailId)
	data, err := database.Get[PearlStats](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("account does not exist")
	}
	return &data[0], nil
}
