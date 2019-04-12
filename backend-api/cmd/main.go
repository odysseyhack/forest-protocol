package main

import (
	"github.com/unchainio/pkg/xconfig"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/config"
	"github.com/unchainio/pkg/iferr"
	"github.com/unchainio/pkg/xlogger"
)

func main() {
	cfg := new(config.ServerConfig)
	info := new(xconfig.Info)

	errs := xconfig.Load(
		cfg,
		xconfig.FromPathFlag("cfg", "config/dev/config.toml"),
		xconfig.FromEnv(),
		xconfig.GetInfo(info),
	)

	log, err := xlogger.New(cfg.Logger)
	iferr.Exit(err)

	log.Printf("Attempted to load configs from %+v", info.Paths)
	iferr.Warn(errs)



}
