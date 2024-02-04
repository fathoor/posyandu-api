package config

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/gofiber/fiber/v2"
)

func ProvideFiber(cfg Config) *fiber.Config {
	return &fiber.Config{
		CaseSensitive: cfg.GetBool("FIBER_CASE_SENSITIVE"),
		StrictRouting: cfg.GetBool("FIBER_STRICT_ROUTING"),
		ErrorHandler:  exception.Handler,
	}
}
