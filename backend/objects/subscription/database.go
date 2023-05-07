package subscription

import (
	"github.com/NoAnguish/PearlerBackend/backend/utils/api_errors"
	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/doug-martin/goqu/v9"
)

var TableName string = "Subscriptions"

func GetById(s *database.Session, id string) (*Subscription, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"id": id}).ToSQL()
	data, err := database.Get[Subscription](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("subscription does not exist")
	}
	return &data[0], nil
}

func Exists(s *database.Session, source string, target string) (bool, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"source": source, "target": target}).ToSQL()
	data, err := database.Get[Subscription](query, s)

	if err != nil {
		return false, api_errors.NewInternalDatabaseError(err)
	}
	return (len(data) > 0), nil
}

func GetBySourceAndTarget(s *database.Session, source string, target string) (*Subscription, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"source": source, "target": target}).ToSQL()
	data, err := database.Get[Subscription](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	if len(data) == 0 {
		return nil, api_errors.NewNotFoundError("subscription does not exist")
	}
	return &data[0], nil
}

func GetBySource(s *database.Session, source string) (*[]Subscription, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"source": source}).ToSQL()
	data, err := database.Get[Subscription](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func GetByTarget(s *database.Session, target string) (*[]Subscription, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Where(goqu.Ex{"target": target}).ToSQL()
	data, err := database.Get[Subscription](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data, nil
}

func Insert(s *database.Session, subscription Subscription) *api_errors.Error {
	query, _, _ := goqu.Insert(TableName).Rows(subscription).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func Update(s *database.Session, subscription Subscription) *api_errors.Error {
	query, _, _ := goqu.Update(TableName).Set(subscription).Where(goqu.Ex{"id": subscription.Id}).ToSQL()
	err := database.Modify(query, s)

	if err != nil {
		return api_errors.NewInternalDatabaseError(err)
	}
	return nil
}

func GetSubsAmountBySourceId(s *database.Session, source string) (*int32, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Select(goqu.COUNT("*").As("cnt")).Where(
		goqu.Ex{"source": source},
		goqu.Ex{"deleted": false},
	).ToSQL()
	data, err := database.Get[int32](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data[0], nil
}

func GetSubsAmountByTargetId(s *database.Session, target string) (*int32, *api_errors.Error) {
	query, _, _ := goqu.From(TableName).Select(goqu.COUNT("*").As("cnt")).Where(
		goqu.Ex{"target": target},
		goqu.Ex{"deleted": false},
	).ToSQL()
	data, err := database.Get[int32](query, s)

	if err != nil {
		return nil, api_errors.NewInternalDatabaseError(err)
	}
	return &data[0], nil
}
