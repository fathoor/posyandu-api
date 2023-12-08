package main

import (
	"github.com/itsLeonB/posyandu-api/core/config"
	"github.com/itsLeonB/posyandu-api/core/exception"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp(cfg)
		_   = config.ProvideDB(cfg)
	)

	err := app.Listen(cfg.Get("APP_ADDRESS"))
	exception.PanicIfNeeded(err)
}
