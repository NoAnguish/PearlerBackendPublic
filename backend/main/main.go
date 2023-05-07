package main

import (
	"github.com/NoAnguish/PearlerBackend/backend/core"
	"github.com/NoAnguish/PearlerBackend/backend/tables"
	"github.com/NoAnguish/PearlerBackend/backend/utils/migrations"
	"github.com/rs/zerolog/log"
)

func main() {
	core.SetupLogs()

	log.Info().Msg("perparing session...")
	d := core.PrepareSession()
	log.Info().Msg("session prepared")

	log.Info().Msg("initialising config...")
	d.InitConfig()
	log.Info().Msg("config initialised")

	log.Info().Msg("making core migrations...")
	migrations.MakeCoreMigration(tables.GetTableData())
	log.Info().Msg("migration was made")

	log.Info().Msg("registering handlers...")
	d.RegisterHandlers()
	log.Info().Msg("Handlers were registered")

	log.Info().Msg("starting daemon session...")
	d.StartSession()
}
