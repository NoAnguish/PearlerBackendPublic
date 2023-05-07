package migrations

import (
	"fmt"
	"sort"

	"github.com/NoAnguish/PearlerBackend/backend/utils/database"
	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog/log"
)

func getAllTableNames(s *database.Session) ([]string, error) {
	query, _, _ := goqu.Select("table_name").From("information_schema.tables").Where(goqu.Ex{"table_schema": "public"}).ToSQL()
	data, err := database.Get[string](query, s)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func dropDatabase() error {
	s, err := database.PrepareDefaultWriteSession()
	if err != nil {
		return err
	}

	tables, err := getAllTableNames(s)
	if err != nil {
		return err
	}

	for _, table := range tables {
		// TODO (noanguish) forgot to drop indexes
		err = database.Modify(fmt.Sprintf("DROP table \"%s\"", table), s)
		if err != nil {
			return err
		}
	}

	err = s.Close()
	if err != nil {
		return err
	}
	return nil
}

func MakeCoreMigration(migrData map[string]func(*database.Session)) {
	s, err := database.PrepareDefaultWriteSession()

	if err != nil {
		log.Error().Err(err).Msg("error while preparing default write session")
		panic(err)
	}

	defer s.Close()

	keys := make([]string, 0)
	for k := range migrData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	version := getDbVersion(s)
	for _, v := range keys {
		if v > version {
			log.Info().Str("Version", v).Msg("applying version to database")

			migrData[v](s)
			version = v
		}
	}

	log.Info().Msg("setting db version...")
	setDbVersion(s, version)
}

func MakeTestMigration(migrData map[string]func(*database.Session)) error {
	err := dropDatabase()
	if err != nil {
		return err
	}

	s, err := database.PrepareDefaultWriteSession()
	if err != nil {
		return err
	}

	keys := make([]string, 0)
	for k := range migrData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		migrData[v](s)
	}

	err = s.Close()
	if err != nil {
		return err
	}
	return nil
}
