package system

import (
	"hyphen-backend-SISS/system/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfig(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicLogging(err)
	return &configImpl{}
}
