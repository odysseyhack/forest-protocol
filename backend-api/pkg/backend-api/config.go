package backend_api

import "github.com/unchainio/pkg/xlogger"

type Config struct {
	Logger *xlogger.Config
	Port string
}
