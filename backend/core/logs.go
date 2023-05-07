package core

import "github.com/rs/zerolog"

func SetupLogs() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
