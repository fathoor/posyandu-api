package config

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"os"
	"strconv"
)

type Config interface {
	Get(key string) string
	GetInt(key string) int
}

type configImpl struct {
}

func (c *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (c *configImpl) GetInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	exception.PanicIfNeeded(err)

	return value
}

func ProvideConfig() Config {
	return &configImpl{}
}
