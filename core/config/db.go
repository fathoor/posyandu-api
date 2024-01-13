package config

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ProvideDB(cfg Config) *gorm.DB {
	var (
		dsn     = cfg.Get("DB_DSN")
		maxOpen = cfg.GetInt("DB_MAX_OPEN")
		maxIdle = cfg.GetInt("DB_MAX_IDLE")
		maxLife = cfg.GetInt("DB_MAX_LIFE")
	)

	dialect := mysql.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	exception.PanicIfNeeded(err)

	mysql, err := db.DB()
	exception.PanicIfNeeded(err)

	mysql.SetMaxOpenConns(maxOpen)
	mysql.SetMaxIdleConns(maxIdle)
	mysql.SetConnMaxLifetime(time.Minute * time.Duration(maxLife))

	return db
}
