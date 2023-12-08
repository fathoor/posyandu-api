package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
)

func ProvideFiber(cfg Config) *fiber.Config {
	return &fiber.Config{
		CaseSensitive: cfg.GetBool("FIBER_CASE_SENSITIVE"),
		StrictRouting: cfg.GetBool("FIBER_STRICT_ROUTING"),
		ErrorHandler:  exception.Handler,
	}
}
