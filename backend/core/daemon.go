package core

import (
	"context"
	"net/http"
	"os"

	"github.com/NoAnguish/PearlerBackend/backend/utils/config"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Daemon struct {
	configPath string
	router     *mux.Router

	ctx context.Context
}

func PrepareSession() *Daemon {
	router := mux.NewRouter().StrictSlash(true)

	configPath := os.Getenv("config_path")

	return &Daemon{
		configPath: configPath,
		ctx:        context.Background(),
		router:     router,
	}
}

func (d *Daemon) RegisterPOST(url string, handler func(http.ResponseWriter, *http.Request)) {
	d.router.HandleFunc(url, handler).Methods("POST")
}

func (d *Daemon) RegisterGET(url string, handler func(http.ResponseWriter, *http.Request)) {
	d.router.HandleFunc(url, handler).Methods("GET")
}

func (d *Daemon) RegisterImage(url string, handler func(http.ResponseWriter, *http.Request)) {
	d.router.HandleFunc(url, handler)
}

func (d *Daemon) InitConfig() {
	err := config.LoadConfig(d.configPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to init config")
		panic(err)
	}
}

func (d *Daemon) StartSession() {
	serverConfig, err := config.ServerConfig()
	if err != nil {
		log.Error().Err(err).Msg("failed get config")
		panic(err)
	}

	url := serverConfig.Host + ":" + serverConfig.Port

	err = http.ListenAndServe(url, d.router)
	if err != nil {
		log.Error().Err(err).Msg("error occured while starting session")
		panic(err)
	}
}
