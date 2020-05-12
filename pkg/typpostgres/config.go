package typpostgres

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Config is postgres configuration
type Config struct {
	DBName   string `required:"true"`
	User     string `required:"true" default:"postgres"`
	Password string `required:"true" default:"pgpass"`
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
}

func (c *Config) ConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

func (c *Config) ConnStrForAdmin() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, "template1")
}

// Configuration of postgres
func Configuration(s *Setting) *typcfg.Configuration {
	if s == nil {
		s = &Setting{}
	}
	return &typcfg.Configuration{
		Name: GetConfigName(s),
		Spec: &Config{
			DBName:   GetDBName(s),
			User:     GetUser(s),
			Password: GetPassword(s),
			Host:     GetHost(s),
			Port:     GetPort(s),
		},
	}
}
