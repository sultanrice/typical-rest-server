// Autogenerated by Typical-Go. DO NOT EDIT!!

package typical

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-rest-server/config"
)

type Config struct {
	App *config.AppConfig
	Pg  *config.PostgresConfig
}

func init() {
	Context.AddConstructor(func() (*Config, error) {
		var cfg Config
		err := envconfig.Process("", &cfg)
		return &cfg, err
	})
	Context.AddConstructor(func(cfg *Config) *config.AppConfig {
		return cfg.App
	})
	Context.AddConstructor(func(cfg *Config) *config.PostgresConfig {
		return cfg.Pg
	})
}
