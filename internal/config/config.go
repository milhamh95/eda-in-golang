package config

import (
	"eda-in-golang/internal/rpc"
	"eda-in-golang/internal/web"
	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
	"os"
	"time"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVL" default:"DEBUG"`
		PG              PGConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT")))
	if err != nil {
		return
	}

	err = envconfig.Process("", &cfg)
	return
}
