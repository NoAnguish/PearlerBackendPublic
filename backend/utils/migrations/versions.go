package migrations

import (
	"errors"
	"fmt"

	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog/log"
)

var versionTableName = "MigrDbVersion"

func createVersionTable(s *database.Session) error {
	err := database.Modify(fmt.Sprintf("CREATE TABLE \"%s\"(version VARCHAR(25));", versionTableName), s)
	return err
}

func getDbVersion(s *database.Session) string {
	query, _, _ := goqu.From(versionTableName).ToSQL()
	version, err := database.Get[string](query, s)
	if err != nil {
		log.Error().Err(err).Msg("error occured while getting db version")
		panic(err)
	}

	if len(version) == 0 {
		err = errors.New("no migration version in database")
		log.Error().Err(err).Msg("error occured while getting db version")
		panic(err)
	}
	if len(version) > 1 {
		err = errors.New("too much migration variables")
		log.Error().Err(err).Msg("error occured while getting db version")
		panic(err)
	}

	return version[0]
}

func setDbVersion(s *database.Session, version string) {
	err := database.Modify(fmt.Sprintf("TRUNCATE TABLE \"%s\";", versionTableName), s)
	if err != nil {
		log.Error().Err(err).Msg("error occured while applying version to db")
		panic(err)
	}

	query, _, _ := goqu.Insert(versionTableName).Cols("version").Vals(goqu.Vals{version}).ToSQL()
	err = database.Modify(query, s)
	if err != nil {
		log.Error().Err(err).Msg("error occured while applying version to db")
		panic(err)
	}
}
