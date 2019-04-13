package main

import (
	"github.com/BurntSushi/toml"
	"github.com/unchainio/pkg/iferr"
	"github.com/unchainio/pkg/xlogger"
	"gitlab.com/henkvanramshorst/forest/backend-api/pgk/backend-api"
	"gitlab.com/henkvanramshorst/forest/backend-api/pgk/router"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	logger, err := xlogger.New(nil)
	iferr.Exit(err)

	logger.Printf("start FOREST backend-api")

	cfgFile, err := ioutil.ReadFile("./config/dev/config.toml")
	iferr.Exit(err)

	cfg := new(backend_api.Config)
	err = toml.Unmarshal(cfgFile, cfg)
	iferr.Exit(err)

	logger.Printf("config loaded: %+v", cfg)

	h, err := router.NewRouter(logger)
	iferr.Exit(err)

	log.Printf("listening on: http://localhost%s", cfg.ApiPort)
	log.Fatal(http.ListenAndServe(cfg.ApiPort, h))
}
