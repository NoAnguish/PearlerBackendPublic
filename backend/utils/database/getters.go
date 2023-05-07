package database

import (
	"github.com/georgysavva/scany/pgxscan"
	postgre "github.com/jackc/pgx/v4"
)

func parseModel[Model any](rows postgre.Rows) ([]Model, error) {
	if rows == nil {
		return make([]Model, 0), nil
	}

	var result []Model

	defer rows.Close()
	for rows.Next() {
		var val Model
		err := pgxscan.ScanRow(&val, rows)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}

	return result, nil
}

func Get[Model any](query string, s *Session) ([]Model, error) {
	needToCloseFlag := false
	var err error

	if s == nil {
		s, err = PrepareDefaultReadSession()
		if err != nil {
			return make([]Model, 0), err
		}
		needToCloseFlag = true
	}

	rows, err := s.get(query)

	if needToCloseFlag {
		s.Close()
	}

	if err != nil {
		return make([]Model, 0), err
	}
	result, err := parseModel[Model](rows)
	return result, err
}
