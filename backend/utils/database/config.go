package database

import (
	"os"

	"github.com/NoAnguish/PearlerBackend/backend/utils/config"
	postgre "github.com/jackc/pgx/v4"
)

type PostgreConfig struct {
	connConfig postgre.ConnConfig
}

func InitConfig() (*PostgreConfig, error) {
	connConfig, err := getConnConfig()

	if err != nil {
		return nil, err
	}

	return &PostgreConfig{connConfig: *connConfig}, nil
}

func getConnConfig() (*postgre.ConnConfig, error) {
	var url string
	dbConfig, err := config.DatabaseConfig()

	if err == nil {
		connConfig, err := postgre.ParseConfig("")

		if err != nil {
			return nil, err
		}

		connConfig.Host = dbConfig.Host
		connConfig.Port = uint16(dbConfig.Port)
		connConfig.Database = dbConfig.Database
		connConfig.User = dbConfig.Username
		connConfig.Password = dbConfig.Password

		return connConfig, nil
	}

	// local_test postgres://test:lolkek@localhost:5432/test_db
	url = os.Getenv("database_url")

	if url == "" {
		return nil, err
	}

	return postgre.ParseConfig(url)
}
