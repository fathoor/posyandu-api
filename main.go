package main

import (
	"github.com/fathoor/posyandu-api/core/config"
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/provider"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg = config.ProvideConfig()
		app = config.ProvideApp()
		db  = config.ProvideDB(cfg)
	)

	provider.ProvideModule(app, db)

	err := app.Listen(cfg.Get("APP_ADDRESS"))
	exception.PanicIfNeeded(err)
}
